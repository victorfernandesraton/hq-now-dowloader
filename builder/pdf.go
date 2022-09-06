package builder

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/go-pdf/fpdf"
)

func BuildToPdf(images []string, outPath string) error {
	pdfPath := fmt.Sprintf("%s.pdf", outPath)
	fmt.Printf("Generate pd %s\n", pdfPath)
	pdf := fpdf.New("P", "mm", "A4", "")
	for _, imgFile := range images {

		if reader, err := os.Open(filepath.Join(imgFile)); err == nil {
			defer reader.Close()
			im, _, err := image.DecodeConfig(reader)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile, err)
				continue
			}
			fmt.Printf("%s %d %d\n", imgFile, im.Width, im.Height)
			pdf.AddPage()
			pdf.ImageOptions(imgFile, 0, 0, 0, 0, false, fpdf.ImageOptions{
				ReadDpi: true,
			}, 0, "")
		} else {
			return err
		}
	}
	return pdf.OutputFileAndClose(pdfPath)
}
