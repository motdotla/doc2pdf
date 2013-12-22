package doc2pdf_test

import (
  doc2pdf "github.com/scottmotte/doc2pdf"
  "testing"
)

func TestConvert(t *testing.T) {
  output, err := doc2pdf.Convert("test/go.doc", "test/go.pdf")

  t.Log(output)
  t.Log(err)


  if len(output) == 0 {
    t.Errorf("Output was blank")
  }
}
