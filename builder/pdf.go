package builder

import (
	"bytes"
	"image"
	"log"

	_ "image/jpeg"
	_ "image/png"

	"github.com/signintech/gopdf"
)

type BulderPdf struct {
	Output string
}

func (b *BulderPdf) Execute(images [][]byte) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 0, H: 0}}) //595.28, 841.89 = A4

	log.Printf("create pdf\n")

	for k, imageBytes := range images {
		log.Printf("add page %v of %v\n", k, len(images))

		imageObject, _, err := image.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			return err
		}

		width := float64(imageObject.Bounds().Dx()) / 1.78
		height := float64(imageObject.Bounds().Dy()) / 1.78

		pdf.AddPageWithOption(gopdf.PageOption{
			PageSize: &gopdf.Rect{
				H: height,
				W: width,
			},
		})

		imgH2, err := gopdf.ImageHolderByBytes(imageBytes)
		if err != nil {
			return err
		}
		if err := pdf.ImageByHolder(imgH2, 0, 0, nil); err != nil {
			return err
		}
	}

	return pdf.WritePdf(b.Output)
}
