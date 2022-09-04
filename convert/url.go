package convert

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	ErrorInvalidHqUrl  = "Invalid hq url"
	ErrorInvalidDomain = "Invalid domain name expect %s, got %s"
	DefaultDomain      = "www.hq-now.com"
	SlashSeparator     = "/"
)

type (
	UrlHQResponse struct {
		Name string
		ID   int
	}
	UrlChapterResponse struct {
		Chapter int
		Name    string
		ID      int
	}
)

func ParseUrlFromHQ(url string) (*UrlHQResponse, error) {
	split := strings.Split(url, SlashSeparator)
	if len(split) != 6 {
		return nil, errors.New(ErrorInvalidHqUrl)
	}

	if split[2] != DefaultDomain {
		return nil, fmt.Errorf(ErrorInvalidDomain, DefaultDomain, split[2])
	}

	idNumber, err := strconv.Atoi(split[4])
	if err != nil {
		return nil, err
	}

	return &UrlHQResponse{
		Name: strings.ReplaceAll(split[5], "%20", " "),
		ID:   idNumber,
	}, nil
}

func ParseUrlFromChapter(url string) (*UrlChapterResponse, error) {
	split := strings.Split(url, SlashSeparator)
	if split[2] != DefaultDomain {
		return nil, fmt.Errorf(ErrorInvalidDomain, DefaultDomain, split[2])
	}
	if len(split) != 10 {
		return nil, errors.New(ErrorInvalidHqUrl)
	}

	idNumber, err := strconv.Atoi(split[4])
	if err != nil {
		return nil, err
	}
	idChapter, err := strconv.Atoi(split[7])
	if err != nil {
		return nil, err
	}

	return &UrlChapterResponse{
		Name:    strings.ReplaceAll(split[5], "%20", " "),
		ID:      idNumber,
		Chapter: idChapter,
	}, nil
}
