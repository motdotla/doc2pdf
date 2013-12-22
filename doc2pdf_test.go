package doc2pdf_test

import (
  doc2pdf "github.com/scottmotte/doc2pdf"
  "testing"
)

func TestConvert(t *testing.T) {
  output, _ := doc2pdf.Convert("test/go.doc", "test/go.pdf")

  if len(output) == 0 {
    t.Errorf("Output was blank")
  }
}

func TestMissingInput(t *testing.T) {
  _, err := doc2pdf.Convert("", "test/go.pdf")

  if err == nil {
    t.Errorf("There should have been an error and there was not.")
  }
}

func TestMissingOutput(t *testing.T) {
  _, err := doc2pdf.Convert("test/go.doc", "")

  if err == nil {
    t.Errorf("There should have been an error and there was not.")
  }
}
