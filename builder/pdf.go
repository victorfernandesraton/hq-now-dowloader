package builder

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/go-pdf/fpdf"
)

func BuildToPdf(images []string, outPath string) error {
	pdfPath := fmt.Sprintf("%s.pdf", outPath)
	log.Printf("Generate pd %s\n", pdfPath)
	pdf := fpdf.New("P", "mm", "Tabloid", "")
	for _, imgFile := range images {

		if reader, err := os.Open(filepath.Join(imgFile)); err == nil {
			defer reader.Close()
			im, _, err := image.DecodeConfig(reader)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile, err)
				continue
			}
			log.Printf("%s %d %d\n", imgFile, im.Width, im.Height)
			pdf.AddPage()
			pdf.ImageOptions(imgFile, 0, 0, 0, 430, false, fpdf.ImageOptions{
				ReadDpi:               true,
				AllowNegativePosition: true,
			}, 0, "")
		} else {
			return err
		}
	}
	return pdf.OutputFileAndClose(pdfPath)
}
