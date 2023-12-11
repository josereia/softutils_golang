package pdf

import (
	"os"

	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type pdfGenerator interface {
	Create(data interface{}) (*[]byte, error)
}

type pdfGeneratorImpl struct {
	TemplatePath string
	Path         string
}

func newPdfGenerator(templatePath string, path string) pdfGenerator {
	return &pdfGeneratorImpl{TemplatePath: templatePath, Path: path}
}

func (generator *pdfGeneratorImpl) Create(data interface{}) (
	*[]byte,
	error,
) {
	parser := newHtmlParser(generator.TemplatePath)
	template, err := parser.Create(data)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(template)
	if err != nil {
		return nil, err
	}

	pdf, err := wk.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdf.Cover.EnableLocalFileAccess.Set(true)
	pdf.AddPage(wk.NewPageReader(file))

	err = file.Close()
	if err != nil {
		return nil, err
	}

	err = pdf.Create()
	if err != nil {
		return nil, err
	}

	err = pdf.WriteFile(generator.Path)
	if err != nil {
		return nil, err
	}

	bytes := pdf.Bytes()
	return &bytes, nil
}
