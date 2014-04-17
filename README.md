# doc2pdf

Convert Word documents to PDFs. Go wrapper for [doc2pdf.net](http://doc2pdf.net). 

[![BuildStatus](https://travis-ci.org/scottmotte/doc2pdf.png?branch=master)](https://travis-ci.org/scottmotte/doc2pdf)

## Usage

```bash
package main

import (
  "fmt"
  doc2pdf "github.com/scottmotte/doc2pdf"
)

func main() {
  path, err := doc2pdf.Convert("/path/to/word.doc", "/desired/path/to/converted.pdf")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(path)
}
```

## Installation

```bash
$ go get github.com/scottmotte/doc2pdf
```

## Running Tests

```bash
go test -v
```
