package commands

import (
	"fmt"
	"log"
	"sync"

	"github.com/pkg/errors"
	"github.com/victorfernandesraton/hq-now-dowloader/builder"
	"github.com/victorfernandesraton/hq-now-dowloader/extract"
)

const ErrorChapterNotFound = "Chapter not found in HQ"

func GetImagesByteByChapter(chapter *extract.HqChapter) ([][]byte, error) {
	var wg sync.WaitGroup
	images := make([][]byte, len(chapter.Pictures))
	wg.Add(len(chapter.Pictures))
	for k, picture := range chapter.Pictures {
		go func(picture string, page int) {
			image, err := builder.DowloadFileAsBytes(picture)
			if err != nil {
				panic(err)
			}
			images[page] = image
			wg.Done()
		}(picture.PictureURL, k)
	}
	wg.Wait()

	return images, nil
}

func BuildChapter(chapter extract.HqChapter, outputFile string) error {
	builder := &builder.BulderPdf{
		Output: outputFile,
	}
	images, err := GetImagesByteByChapter(&chapter)
	if err != nil {
		return err
	}
	return builder.Execute(images)
}

func GeAllChapters(id int) error {

	hqInfo, err := extract.GetHqInfo(id)
	if err != nil {
		return err
	}
	log.Printf("find data for hq %v with %v chapters\n", hqInfo.Name, len(hqInfo.Capitulos))

	var wg sync.WaitGroup
	wg.Add(len(hqInfo.Capitulos))
	for _, chapter := range hqInfo.Capitulos {
		go func(chapter extract.HqChapter) {
			if err := BuildChapter(chapter, fmt.Sprintf("%v-%v.pdf", hqInfo.Name, chapter.Number)); err != nil {
				panic(err)
			}
			wg.Done()
		}(chapter)
	}
	wg.Wait()

	return nil
}

func GetByChapter(id int, chapter string) error {

	hqInfo, err := extract.GetHqInfo(id)
	if err != nil {
		return err
	}

	log.Printf("find data for hq %v with chapter %v\n", hqInfo.Name, chapter)

	for _, v := range hqInfo.Capitulos {
		if v.Number == chapter {
			return BuildChapter(v, fmt.Sprintf("%v-%v.pdf", hqInfo.Name, v.Number))
		}
	}

	return errors.New(ErrorChapterNotFound)
}
