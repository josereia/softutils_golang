package mimetype

import "errors"

// Mime types enum.
type MimeTypesEnum uint8

const (
	MP4MimeType = iota + 1
	JPGMimeType
	JPEGMimeType
	PNGMimeType
	TXTMimeType
	DOCMimeType
	DOCXMimeType
	PPTMimeType
	PPTXMimeType
	XLSMimeType
	XLSXMimeType
	PDFMimeType
	StreamMimeType
)

// A MimeType provides all methods.
type MimeType struct{}

// Get mime type by mime types enum.
// Return: string, error. 
func (helper *MimeType) GetMimeType(
	mimeType MimeTypesEnum,
) (string, error) {
	value, ok := mimeTypes[mimeType]
	if !ok {
		return "", errors.New("mime type not found")
	}

	return value, nil
}

// Get mime type by file extension.
// Return: MimeTypesEnum, string.
func (helper *MimeType) GetMimeTypeByExtension(
	extension string,
) (MimeTypesEnum, string) {
	switch extension {
	case "mp4":
		return MP4MimeType, mimeTypes[MP4MimeType]
	case "jpg":
		return JPGMimeType, mimeTypes[JPGMimeType]
	case "jpeg":
		return JPEGMimeType, mimeTypes[JPEGMimeType]
	case "png":
		return PNGMimeType, mimeTypes[PNGMimeType]
	case "txt":
		return TXTMimeType, mimeTypes[TXTMimeType]
	case "doc":
		return DOCMimeType, mimeTypes[DOCMimeType]
	case "docx":
		return DOCXMimeType, mimeTypes[DOCXMimeType]
	case "ppt":
		return PPTMimeType, mimeTypes[PPTMimeType]
	case "pptx":
		return PPTXMimeType, mimeTypes[PPTXMimeType]
	case "xls":
		return XLSMimeType, mimeTypes[XLSMimeType]
	case "xlsx":
		return XLSXMimeType, mimeTypes[XLSXMimeType]
	case "pdf":
		return PDFMimeType, mimeTypes[PDFMimeType]
	default:
		return StreamMimeType, mimeTypes[StreamMimeType]
	}
}

// Compare mime type name with enums.
// Return: bool, error
func (helper *MimeType) CompareMimeType(
	name string,
	enums ...MimeTypesEnum,
) (bool, error) {
	for i := 0; i < len(enums); i++ {
		enum := enums[i]

		value, ok := mimeTypes[enum]
		if !ok {
			return false, errors.New("mime type not found")
		}

		if value == name {
			return true, nil
		}
	}

	return false, nil
}

var mimeTypes = map[MimeTypesEnum]string{
	MP4MimeType:    "video/mp4",
	JPGMimeType:    "image/jpeg",
	JPEGMimeType:   "image/jpeg",
	PNGMimeType:    "image/png",
	TXTMimeType:    "text/plain",
	DOCMimeType:    "application/msword",
	DOCXMimeType:   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	PPTMimeType:    "application/vnd.ms-powerpoint",
	PPTXMimeType:   "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	XLSMimeType:    "application/vnd.ms-excel",
	XLSXMimeType:   "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	PDFMimeType:    "application/pdf",
	StreamMimeType: "application/octet-stream",
}
