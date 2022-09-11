package commands

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/victorfernandesraton/hq-now-dowloader/builder"
	"github.com/victorfernandesraton/hq-now-dowloader/extract"
)

const ErrorChapterNotFound = "Chapter not found in HQ"

func GeAllChapters(id int) error {
	var wg sync.WaitGroup

	hqInfo, err := extract.GetHqInfo(id)
	if err != nil {
		return err
	}
	log.Printf("find data for hq %v with %v chapters\n", hqInfo.Name, len(hqInfo.Capitulos))

	wg.Add(len(hqInfo.Capitulos))
	for _, chapter := range hqInfo.Capitulos {
		go func(chapter extract.HqChapter) {

			images, err := GetImagesByteByChapter(&chapter)
			if err != nil {
				panic(err)
			}
			builder := &builder.BulderPdf{
				Output: fmt.Sprintf("%v-%v.pdf", hqInfo.Name, chapter.Number),
			}
			if err := builder.Execute(images); err != nil {
				panic(err)
			}
			wg.Done()
		}(chapter)
	}
	wg.Wait()

	return nil
}

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

func GetByChapter(id int, chapter string) error {
	var wg sync.WaitGroup
	var chapters []extract.HqChapter
	hqInfo, err := extract.GetHqInfo(id)
	if err != nil {
		return err
	}

	for _, v := range hqInfo.Capitulos {
		if v.Number == chapter {
			chapters = append(chapters, v)
		}
	}

	if len(chapters) == 0 {
		return errors.New(ErrorChapterNotFound)
	}

	log.Printf("find data for hq %v with chapter %v\n", hqInfo.Name, chapter)

	wg.Add(len(chapters))
	for _, chapter := range chapters {
		go func(chapter extract.HqChapter) {

			images, err := GetImagesByteByChapter(&chapter)
			if err != nil {
				panic(err)
			}
			builder := &builder.BulderPdf{
				Output: fmt.Sprintf("%v-%v.pdf", hqInfo.Name, chapters[0].Number),
			}
			if err := builder.Execute(images); err != nil {
				panic(err)
			}
			wg.Done()
		}(chapter)
	}
	wg.Wait()
	return nil
}
