package main

import (
  "bytes"
  cli "github.com/codegangsta/cli"
  "io"
  "io/ioutil"
  "log"
  "mime/multipart"
  "net/http"
  "os"
  "path/filepath"
)

func main() {
  const VERSION           = "0.0.1"
  const DOC2PDF_ENDPOINT  = "http://www.doc2pdf.net/convert/document.pdf" 

  app         := cli.NewApp()
  app.Version = VERSION
  app.Name    = "doc2pdf"
  app.Usage   = "Convert word documents to PDFs. Wrapper for doc2pdf.net."
  app.Flags   = []cli.Flag {
    cli.StringFlag{"input", "", "/path/to/word.doc"},
    cli.StringFlag{"output", "", "/desired/path/to/converted.pdf"},
  }

  app.Action = func(c *cli.Context) {
    input_path  := c.String("input")
    output_path := c.String("output")

    if len(input_path) == 0 {
      log.Fatal("FLAG 'input' IS REQUIRED")
    }
    if len(output_path) == 0 {
      log.Fatal("FLAG 'output' IS REQUIRED")
    }

    // Get the file data from the word doc
    file, err := os.Open(input_path)
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()

    // Build the {inputDocument: 'file/path'} into the multipart request
    body      := &bytes.Buffer{}
    writer    := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("inputDocument", filepath.Base(input_path))
    if err != nil {
      log.Fatal(err)
    }

    // Copy the file into the part
    io.Copy(part, file)

    // Close the writer
    err = writer.Close()
    if err != nil {
      log.Fatal(err)
    }

    // Prepare the request and headers
    req, err := http.NewRequest("POST", DOC2PDF_ENDPOINT, body)
    if err != nil {
      log.Fatal(err)
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    // Make the request
    res, err := http.DefaultClient.Do(req)
    if err != nil {
      log.Fatal(err)
    }

    // Read and save the repsonse
    responseBody, err := ioutil.ReadAll(res.Body);
    if err != nil {
      log.Fatal(err)
    }
    ioutil.WriteFile(output_path, responseBody, 0777)

    // Success message
    log.Println("SUCCESS "+output_path)
  }

  app.Run(os.Args)
}
