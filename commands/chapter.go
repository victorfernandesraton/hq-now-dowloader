package commands

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/victorfernandesraton/hq-now-dowloader/extract"
)

func CreateByChapter(chapterId int) error {
	var wg sync.WaitGroup
	chapterInfo, err := extract.GetHqPages(chapterId)
	log.Printf("dowload chapter %v of %v\n", chapterInfo.Number, chapterInfo.Name)
	log.Printf("creating chapter %v with %v pages\n", chapterInfo.Number, len(chapterInfo.Pictures))
	serieFolder := fmt.Sprintf("output/%v/", chapterInfo.Name)
	folder := fmt.Sprintf("%v/%v", serieFolder, chapterInfo.Number)
	if err != nil {
		return err
	}
	for k, page := range chapterInfo.Pictures {
		wg.Add(1)
		log.Printf("extract and dowload file %v\n", page.PictureURL)
		log.Printf("find in folder %s\n", serieFolder)
		err = os.MkdirAll(folder, 0755)
		if err != nil {
			wg.Done()
			return err
		}
		go func(picture string, page int) {
			getImages(fmt.Sprintf("%v/%v", folder, page), picture)
			wg.Done()
		}(page.PictureURL, k)
	}
	wg.Wait()
	return GeneratePdf(folder)
}
