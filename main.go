package main

import (
	"log"

	p "github.com/juanmachuca95/convertidor_webp_go/pdf"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {

	// Setup
	imagick.Initialize()
	defer imagick.Terminate()

	docName := "test_word.doc"
	pdfName := "test.pdf"
	docImageName := "doc.jpg"
	pdfImageName := "pdf.jpg"

	// Word implementation
	if err := p.ConvertPdfToJpg(docName, docImageName); err != nil {
		log.Fatal(err)
	}

	// PDF implementation
	if err := p.ConvertPdfToJpg(pdfName, pdfImageName); err != nil {
		log.Fatal(err)
	}

	log.Println("Programa terminado")
}
