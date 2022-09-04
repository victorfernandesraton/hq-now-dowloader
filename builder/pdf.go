package builder

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

func BuildToPdf(images []string, outPath string) error {
	pdfPath := fmt.Sprintf("%s.pdf", outPath)
	fmt.Printf("Generate pd %s\n", pdfPath)
	pdf := fpdf.New("P", "mm", "A4", "")
	for _, v := range images {
		pdf.AddPage()
		pdf.ImageOptions(v, 0, 0, 0, 0, false, fpdf.ImageOptions{}, 0, "")
	}
	return pdf.OutputFileAndClose(pdfPath)
}
