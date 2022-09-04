package convert_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/victorfernandesraton/hq-now-dowloader/convert"
)

func TestValidUrlFromHq(t *testing.T) {
	testCases := []struct {
		desc     string
		url      string
		response *convert.UrlHQResponse
		Err      error
	}{
		{
			desc: "Shoud be a valid url for hq",
			url:  "https://www.hq-now.com/hq/2879/Gata%20Negra%20(2019)",
			response: &convert.UrlHQResponse{
				Name: "Gata Negra (2019)",
				ID:   2879,
			},
			Err: nil,
		},
		{
			desc:     "Shoud be a not a valid url",
			url:      "https://www.hq-now.com/",
			response: nil,
			Err:      errors.New(convert.ErrorInvalidHqUrl),
		},
		{
			desc:     "Shoud be a not a valid domain",
			url:      "https://www.google.com/test/to/path",
			response: nil,
			Err:      fmt.Errorf(convert.ErrorInvalidDomain, convert.DefaultDomain, "www.google.com"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			result, err := convert.ParseUrlFromHQ(tC.url)
			if !cmp.Equal(tC.response, result) {
				t.Fatalf("%v got %v", tC.response, result)
			}
			if err != nil && tC.Err != nil {
				if err.Error() != tC.Err.Error() {
					t.Fatalf("%v got %v", tC.Err, err)
				}
			}
		})
	}
}

func TestUrlFromChapter(t *testing.T) {
	testCases := []struct {
		desc     string
		url      string
		response *convert.UrlChapterResponse
		Err      error
	}{
		{
			desc: "Shoud be a valid url for hq",
			url:  "https://www.hq-now.com/hq-reader/33103/vingadores-2018/chapter/39/page/1",
			response: &convert.UrlChapterResponse{
				Name:    "vingadores-2018",
				ID:      33103,
				Chapter: 39,
			},
			Err: nil,
		},
		{
			desc:     "Shoud be a not a valid url",
			url:      "https://www.hq-now.com/",
			response: nil,
			Err:      errors.New(convert.ErrorInvalidHqUrl),
		},
		{
			desc:     "Shoud be a not a valid domain",
			url:      "https://www.google.com/test/to/path/final/error/ok",
			response: nil,
			Err:      fmt.Errorf(convert.ErrorInvalidDomain, convert.DefaultDomain, "www.google.com"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			result, err := convert.ParseUrlFromChapter(tC.url)
			if !cmp.Equal(tC.response, result) {
				t.Fatalf("%v got %v", tC.response, result)
			}
			if err != nil && tC.Err != nil {
				if err.Error() != tC.Err.Error() {
					t.Fatalf("%v got %v", tC.Err, err)
				}
			}
		})
	}
}
