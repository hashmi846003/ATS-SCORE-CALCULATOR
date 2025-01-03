package parser

import (
    "bytes"
    "fmt"
    "os"

    "github.com/unidoc/unioffice/document"
    "github.com/unidoc/unipdf/v3/extractor"
    "github.com/unidoc/unipdf/v3/model"
)

func ParseDOCX(filepath string) (string, error) {
    // Open the DOCX file
    doc, err := document.Open(filepath)
    if err != nil {
        return "", err
    }
    defer doc.Close()

    // Extract text from the document
    var text string
    for _, para := range doc.Paragraphs() {
        for _, run := range para.Runs() {
            text += run.Text()
        }
        text += "\n"
    }
    return text, nil
}

func ParsePDF(filepath string) (string, error) {
    // Open the PDF file
    f, err := os.Open(filepath)
    if err != nil {
        return "", err
    }
    defer f.Close()

    pdfReader, err := model.NewPdfReader(f)
    if err != nil {
        return "", err
    }

    numPages, err := pdfReader.GetNumPages()
    if err != nil {
        return "", err
    }

    var buf bytes.Buffer
    for i := 1; i <= numPages; i++ {
        page, err := pdfReader.GetPage(i)
        if err != nil {
            return "", err
        }

        ex, err := extractor.New(page)
        if err != nil {
            return "", err
        }

        text, err := ex.ExtractText()
        if err != nil {
            return "", err
        }

        buf.WriteString(text)
        buf.WriteString("\n")
    }

    return buf.String(), nil
}

func main() {
    // Example usage
    docText, err := ParseDOCX("example.docx")
    if err != nil {
        fmt.Println("Error parsing DOCX:", err)
    } else {
        fmt.Println("DOCX Content:")
        fmt.Println(docText)
    }

    pdfText, err := ParsePDF("example.pdf")
    if err != nil {
        fmt.Println("Error parsing PDF:", err)
    } else {
        fmt.Println("PDF Content:")
        fmt.Println(pdfText)
    }
}
