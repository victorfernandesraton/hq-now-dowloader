package builder

import (
	"log"

	"github.com/signintech/gopdf"
)

type BulderPdf struct {
	Output string
}

func (b *BulderPdf) Execute(images [][]byte) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

	log.Printf("create pdf\n")

	for k, image := range images {
		log.Printf("add page %v of %v\n", k, len(images))

		pdf.AddPage()

		imgH2, err := gopdf.ImageHolderByBytes(image)
		if err != nil {
			return err
		}
		if err := pdf.ImageByHolder(imgH2, 0, 0, nil); err != nil {
			return err
		}
	}

	return pdf.WritePdf(b.Output)
}
