package builder

import "github.com/go-pdf/fpdf"

func BuildToPdf(images []string) error {

	pdf := fpdf.New("P", "mm", "A4", "")
	for _, v := range images {
		pdf.AddPage()
		pdf.Image(v, 0, 0, 100, 200, false, "", 0, "")
	}
	return pdf.OutputFileAndClose("hello.pdf")
}
