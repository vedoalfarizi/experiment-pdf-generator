package main

import (
	"github.com/experiments/pdfgen/examples/wakaf"
	"log"
)

func main() {
	wakafContent := wakaf.GetPDFContent()

	pdfRes, err := wakafContent.Generate()
	if err != nil {
		log.Fatalf("failed to generate PDF cause :%+v", err)
	}

	err = pdfRes.Save("examples/wakaf/result/wakaf.pdf")
	if err != nil {
		log.Fatalf("failed to save PDF cause :%+v", err)
	}

	log.Println("successfully generate PDF")
}
