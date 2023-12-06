package pdf

import (
	"html/template"
	"os"
	p "path"
)

type htmlParser interface {
	Create(data interface{}) (string, error)
}

type htmlParserImpl struct {
	Path string
}

func newHtmlParser(path string) htmlParser {
	return &htmlParserImpl{Path: path}
}

func (parser *htmlParserImpl) Create(data interface{}) (string, error) {
	writer, err := template.ParseFiles(parser.Path)
	if err != nil {
		return "", err
	}

	newPath := p.Join(p.Dir(parser.Path), "final_template.html")
	file, err := os.Create(newPath)
	if err != nil {
		return "", err
	}

	err = writer.Execute(file, data)
	if err != nil {
		return "", err
	}

	return newPath, nil
}
