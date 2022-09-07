package builder

import (
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const FixedWidth = 794

func DownloadFile(filepath string, url string) (err error) {

	log.Printf("make requeest %v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	log.Printf("create file %v\n", filepath)
	ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func FindFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(path, ".pdf") {
			log.Printf("reading data from %s\n", path)
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ResizeImage(path string, output string) (*string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	outputFile, err := os.Create(output)
	defer outputFile.Close()
	if err != nil {
		return nil, err
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	scale := img.Bounds().Dx() / FixedWidth
	height := img.Bounds().Dy() * scale

	dst := image.NewRGBA(image.Rect(0, 0, FixedWidth, height))

	draw.Draw(dst, image.Rect(image.ZP.X, image.ZP.Y, 400, 499), img, img.Bounds().Size(), draw.Over)

	if err := jpeg.Encode(outputFile, dst, &jpeg.Options{}); err != nil {
		return nil, err
	}
	return &output, nil
}
