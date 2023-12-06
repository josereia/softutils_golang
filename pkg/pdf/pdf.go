package pdf

type Pdf struct {
	TemplatePath string
	TemplateData interface{}
	Path         string
}

func (pdf *Pdf) Generate() (string, error) {
	generator := newPdfGenerator(pdf.TemplatePath, pdf.Path)
	path, err := generator.Create(pdf.TemplateData)
	if err != nil {
		return "", err
	}

	return path, nil
}
