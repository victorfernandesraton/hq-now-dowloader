package commands

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/victorfernandesraton/hq-now-dowloader/builder"
	"github.com/victorfernandesraton/hq-now-dowloader/extract"
)

func getImages(folder string, url string) error {
	return builder.DownloadFile(fmt.Sprintf("%v.jpg", folder), url)
}

func GetPages(chapter *extract.GetHqChapterResponse) error {
	var wg sync.WaitGroup
	data, err := extract.GetHqPages(chapter.ID)
	if err != nil {
		return err
	}

	for k, page := range data.Pictures {
		wg.Add(1)
		fmt.Printf("extract and dowload file %v\n", page.PictureURL)
		folder := fmt.Sprintf("output/%v/%v", data.Name, chapter.Number)
		err = os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
		go func(picture string, page int) {
			getImages(fmt.Sprintf("%v/%v", folder, page), picture)
			wg.Done()
		}(page.PictureURL, k)
	}
	wg.Wait()
	return nil
}

func CreateAllChapters(id int) error {
	var wg sync.WaitGroup

	chapters, err := extract.GetHqChapter(id)
	if err != nil {
		return err
	}

	for _, chapter := range chapters {
		wg.Add(1)
		go func(chapter *extract.GetHqChapterResponse) {
			GetPages(chapter)
			wg.Done()
		}(chapter)
	}
	wg.Wait()
	return nil
}

func CreateByOneChapter(id int, chapter int) error {
	var current *extract.GetHqChapterResponse

	chapters, err := extract.GetHqChapter(id)
	if err != nil {
		return err
	}

	for _, item := range chapters {
		num, err := strconv.Atoi(item.Number)
		if err != nil {
			return err
		}
		if item.ID == chapter || num == chapter {
			current = item
		}
	}

	if err = GetPages(current); err != nil {
		return err
	}
	return nil
}
