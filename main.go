package main

import (
	"fmt"
	"log"

	w "github.com/juanmachuca95/convertidor_webp_go.git/webp"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	pdfName := "test.pdf"
	imageName := "conceptos.jpg"

	if err := ConvertPdfToJpg(pdfName, imageName); err != nil {
		log.Fatal(err)
	}

	imageJpeg := "2_1_conceptos.jpg"
	nameWebp := "2_1_conceptos.webp"
	w.ConvertToWebp(imageJpeg, nameWebp)

	log.Println("Programa terminado")
}

// ConvertPdfToJpg will take a filename of a pdf file and convert the file into an
// image which will be saved back to the same location. It will save the image as a
// high resolution jpg file with minimal compression.
func ConvertPdfToJpg(pdfName string, imageName string) error {
	// Setup
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Must be *before* ReadImageFile
	// Make sure our image is high quality
	if err := mw.SetResolution(300, 300); err != nil {
		return err
	}

	// Load the image file into imagick
	if err := mw.ReadImage(pdfName); err != nil {
		log.Println("error al leer archivo ", err)
		return err
	}

	log.Println(mw.GetNumberImages())
	for i := 1; i < int(mw.GetNumberImages()); i++ {

		// Select only first page of pdf
		mw.SetIteratorIndex(i)

		// Must be *after* ReadImageFile
		// Flatten image and remove alpha channel, to prevent alpha turning black in jpg
		if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_BACKGROUND); err != nil {
			return err
		}

		if err := mw.SetImageCompression(imagick.COMPRESSION_JPEG); err != nil {
			fmt.Printf("%v", err)
		}

		// Set any compression (100 = max quality)
		if err := mw.SetCompressionQuality(80); err != nil {
			return err
		}

		// Convert into JPG
		if err := mw.SetFormat("jpg"); err != nil {
			return err
		}

		// Save File
		imageName = fmt.Sprintf("%d_%s", i, imageName)
		mw.WriteImage(imageName)
	}

	return nil

}
