package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func BlurImage(inputPath, outputPath string, blurRadius float64) error {
	// Open the image file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	// Blur the image
	blurredImg := imaging.Blur(img, blurRadius)

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Save the blurred image to the output file
	err = jpeg.Encode(outputFile, blurredImg, nil)
	if err != nil {
		return err
	}

	fmt.Println("Image blurred successfully.")
	return nil
}

func main() {
	inputPath := "assets/input/image.JPG"          // Specify the path to your input image
	outputPath := "assets/output/blurredimage.jpg" // Specify the path for the output blurred image
	blurRadius := 50.0                             // Set your desired blur radius

	err := BlurImage(inputPath, outputPath, blurRadius)
	if err != nil {
		log.Fatal(err)
	}
}
