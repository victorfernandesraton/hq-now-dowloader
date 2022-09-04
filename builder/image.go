package builder

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadFile(filepath string, url string) (err error) {

	fmt.Printf("make requeest %v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Printf("create file %v\n", filepath)
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
			fmt.Printf("reading data from %s\n", path)
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
