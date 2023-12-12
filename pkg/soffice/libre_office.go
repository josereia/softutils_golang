package soffice

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type libreOffice interface {
	ToPdf(path string) (*[]byte, error)
}

type libreOfficeImpl struct {
	Path string `json:"path,omitempty"`
}

func newLibreOffice(path string) libreOffice {
	return &libreOfficeImpl{Path: path}
}

func (office *libreOfficeImpl) ToPdf(path string) (*[]byte, error) {
	tempDir := os.TempDir()

	command := exec.Command(
		"soffice",
		"--headless",
		"--convert-to",
		"pdf",
		path,
		"--outdir",
		tempDir,
	)

	err := command.Run()
	if err != nil {
		return nil, err
	}

	inputFileName := strings.Split(filepath.Base(path), ".")[0] + ".pdf"
	outputPath := filepath.Join(tempDir, inputFileName)

	_, err = os.Stat(outputPath)
	if !os.IsNotExist(err) {
		bytes, err := os.ReadFile(outputPath)
		if err != nil {
			return nil, err
		}

		err = os.Remove(outputPath)
		if err != nil {
			return nil, err
		}

		return &bytes, nil
	}

	return nil, err
}
