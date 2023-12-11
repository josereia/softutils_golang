package pdf

type Pdf struct {
	TemplatePath string
	TemplateData interface{}
	Path         string
}

func (pdf *Pdf) Generate() (*[]byte, error) {
	generator := newPdfGenerator(pdf.TemplatePath, pdf.Path)
	buffer, err := generator.Create(pdf.TemplateData)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
