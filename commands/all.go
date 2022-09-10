package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/victorfernandesraton/hq-now-dowloader/builder"
	"github.com/victorfernandesraton/hq-now-dowloader/extract"
)

const ErrorChapterNotFound = "Chapter not found in HQ"

func getImages(folder string, url string) error {
	return builder.DownloadFile(fmt.Sprintf("%v.jpg", folder), url)
}

func ProcessConcurrentImages(folder string, chapters []extract.HqChapter) {
	var wg sync.WaitGroup

	wg.Add(len(chapters))
	for _, chapter := range chapters {
		log.Printf("process image chapter %v of %v", chapter.Number, len(chapters))
		filename := fmt.Sprintf("%v/%v", folder, chapter.Number)
		go func(ch extract.HqChapter) {
			if err := GetPages(filename, ch); err != nil {
				panic(err)
			}
			wg.Done()
		}(chapter)
	}
	wg.Wait()
	return
}

func GetPages(foldername string, chapter extract.HqChapter) error {
	var wg sync.WaitGroup

	log.Printf("creating chapter %v(%v) with %v pages\n", chapter.Number, chapter.ID, len(chapter.Pictures))

	wg.Add(len(chapter.Pictures))
	for k, page := range chapter.Pictures {
		log.Printf("extract and dowload file %v\n", page.PictureURL)
		log.Printf("find in folder %s", foldername)
		if err := os.MkdirAll(foldername, 0755); err != nil {
			panic(err)
		}
		go func(picture string, page int) {
			getImages(fmt.Sprintf("%v/%v", foldername, page), picture)
			wg.Done()
		}(page.PictureURL, k)
	}
	wg.Wait()
	return GeneratePdf(foldername)
}

func GeAllChapters(id int) error {
	hqInfo, err := extract.GetHqInfo(id)
	if err != nil {
		return err
	}
	log.Printf("find data for hq %v with %v chapters\n", hqInfo.Name, len(hqInfo.Capitulos))

	folder := fmt.Sprintf("output/%v", hqInfo.Name)
	log.Printf("processsing chapters for %s, output in folder %s", hqInfo.Name, folder)
	ProcessConcurrentImages(folder, hqInfo.Capitulos)
	return nil
}

func GetByChapter(id int, chapter string) error {
	hqInfo, err := extract.GetHqInfo(id)
	var chapters []extract.HqChapter
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

	folder := fmt.Sprintf("output/%v", hqInfo.Name)
	log.Printf("processsing chapters for %s, output in folder %s", hqInfo.Name, folder)
	ProcessConcurrentImages(folder, chapters)
	return nil
}

func GeneratePdf(serieFolder string) error {
	files, err := builder.FindFiles(serieFolder)
	if err != nil {
		return err
	}
	return builder.BuildToPdf(files, fmt.Sprintf("%s/output", serieFolder))
}
