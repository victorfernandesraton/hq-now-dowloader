package builder

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
