package pdf

import "bytes"

type Pdf struct {
	TemplatePath string
	TemplateData interface{}
	Path         string
}

func (pdf *Pdf) Generate() (*bytes.Buffer, error) {
	generator := newPdfGenerator(pdf.TemplatePath, pdf.Path)
	buffer, err := generator.Create(pdf.TemplateData)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
