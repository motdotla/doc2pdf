package doc2pdf

import (
  "bytes"
  "errors"
  "io"
  "io/ioutil"
  "mime/multipart"
  "net/http"
  "os"
  "path/filepath"
)

const VERSION           = "0.0.1"
const DOC2PDF_ENDPOINT  = "http://www.doc2pdf.net/convert/document.pdf"

func Convert(input_path string, output_path string) (string, error) {
  if len(input_path) == 0 {
    err := errors.New("Missing input path")
    return "", err
  }
  if len(output_path) == 0 {
    err := errors.New("Missing output path")
    return "", err
  }

  // Get the file data from the word doc
  file, err := os.Open(input_path)
  if err != nil {
    return "", err
  }
  defer file.Close()

  // Build the {inputDocument: 'file/path'} into the multipart request
  body      := &bytes.Buffer{}
  writer    := multipart.NewWriter(body)
  part, err := writer.CreateFormFile("inputDocument", filepath.Base(input_path))
  if err != nil {
    return "", err
  }

  // Copy the file into the part
  io.Copy(part, file)

  // Close the writer
  err = writer.Close()
  if err != nil {
    return "", err
  }

  // Prepare the request and headers
  req, err := http.NewRequest("POST", DOC2PDF_ENDPOINT, body)
  if err != nil {
    return "", err
  }
  req.Header.Set("Content-Type", writer.FormDataContentType())

  // Make the request
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return "", err
  }

  // Read and save the repsonse
  responseBody, err := ioutil.ReadAll(res.Body);
  if err != nil {
    return "", err
  }
  ioutil.WriteFile(output_path, responseBody, 0777)

  return output_path, nil
}
