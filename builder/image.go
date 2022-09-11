package builder

import (
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
)

func DowloadFileAsBytes(url string) ([]byte, error) {
	log.Printf("make requeest %v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
