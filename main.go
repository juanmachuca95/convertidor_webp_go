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

	pdfName := "test.pdf"
	imageName := "conceptos.jpg"

	if err := p.ConvertPdfToJpg(pdfName, imageName); err != nil {
		log.Fatal(err)
	}

	log.Println("Programa terminado")
}
