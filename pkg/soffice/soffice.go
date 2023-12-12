// A bridge between golang and the libre office command line tool.
// This package focuses on file conversions.
package soffice

// A Soffice struct provides LibreOffice setup and methods.
type Soffice struct {
	// LibreOffice installation path.
	Path string `json:"path,omitempty"`
}

// Converts any file to pdf using LibreOffice.
func (office *Soffice) ToPdf(path string) (*[]byte, error) {
	lOffice := newLibreOffice(office.Path)

	bytes, err := lOffice.ToPdf(path)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
