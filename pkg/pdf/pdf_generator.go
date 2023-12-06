package pdf

import (
	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"os"
)

type pdfGenerator interface {
	Create(data interface{}) (string, error)
}

type pdfGeneratorImpl struct {
	TemplatePath string
	Path         string
}

func newPdfGenerator(templatePath string, path string) pdfGenerator {
	return &pdfGeneratorImpl{TemplatePath: templatePath, Path: path}
}

func (generator *pdfGeneratorImpl) Create(data interface{}) (string, error) {
	parser := newHtmlParser(generator.TemplatePath)
	template, err := parser.Create(data)
	if err != nil {
		return "", err
	}

	file, err := os.Open(template)
	if err != nil {
		return "", err
	}

	pdf, err := wk.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	pdf.AddPage(wk.NewPageReader(file))

	err = pdf.Create()
	if err != nil {
		return "", err
	}

	err = pdf.WriteFile(generator.Path)
	if err != nil {
		return "", err
	}

	err = file.Close()
	if err != nil {
		return "", err
	}

	// err = os.Remove(template)
	// if err != nil {
	// 	return "", err
	// }

	return generator.Path, nil
}
