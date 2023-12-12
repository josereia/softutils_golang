# SoftUtils Golang

A set of tools for your golang applications.

## PDF

A package for pdf file creation by html template file.

**How to use:**

```go
package main

import "github.com/softyesti/softutils_golang/pkg/pdf"

func main() {
    pdfi := &pdf.Pdf{
        TemplatePath: "path"
	TemplateData: interface{}
    }

    bytes, err := pdfi.Generate()
    if err != nil {
        panic(err)
    }
}
```

## Soffice

A package that provides a bridge between golang and the libre office command line tool. This package focuses on file conversions.

**How to use:**

```go
package main

import "github.com/softyesti/softutils_golang/pkg/soffice"

func main() {
    office := &soffice.Soffice{
        Path: "path" // optional custom libre office installation path
    }

    bytes, err := office.ToPdf("file_path")
    if err != nil {
        panic(err)
    }
}

```

## MimeType

A package that provides mimetype utils like enum, find and compare functions.

**How to use:**

```go
package main

import "github.com/softyesti/softutils_golang/pkg/mimetype"

func main() {
    helper:= &mimetype.MimeType{}

    mimeTypeName, err := helper.GetMimeType(mimetype.PDFMimeType)
    if err != nil {
        panic(err)
    }

    mimeTypeEnum, mimeTypeName := helper.GetMimeTypeByExtension("pdf")

    isDocx, err := helper.CompareMimeType("application/msword", helper.DOCMimeType)
    if err != nil {
        panic(err)
    }
}

```
