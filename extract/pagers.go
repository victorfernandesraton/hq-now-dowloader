package extract

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type (
	pageResponse struct {
		Data struct {
			GetChapterByID struct {
				Name     string `json:"name"`
				Number   string `json:"number"`
				Oneshot  bool   `json:"oneshot"`
				Pictures []struct {
					PictureURL string `json:"pictureUrl"`
				} `json:"pictures"`
				Hq struct {
					ID        int    `json:"id"`
					Name      string `json:"name"`
					Capitulos []struct {
						ID     int    `json:"id"`
						Number string `json:"number"`
					} `json:"capitulos"`
				} `json:"hq"`
			} `json:"getChapterById"`
		} `json:"data"`
	}
	GetHqPagesResponse struct {
		Name     string `json:"name"`
		Number   string `json:"number"`
		Pictures []struct {
			PictureURL string `json:"pictureUrl"`
		} `json:"pictures"`
	}
)

func GetHqPages(hqChaperUniqueId int) (*GetHqPagesResponse, error) {

	url := "https://admin.hq-now.com/graphql"

	payload := strings.NewReader(fmt.Sprintf("{\n\t\"operationName\": \"getChapterById\",\n\t\"variables\": {\n\t\t\"chapterId\": %v\n\t},\n\t\"query\": \"query getChapterById($chapterId: Int!) {\\n  getChapterById(chapterId: $chapterId) {\\n    name\\n    number\\n    oneshot\\n    pictures {\\n      pictureUrl\\n    }\\n    hq {\\n      id\\n      name\\n      capitulos {\\n        id\\n        number\\n      }\\n    }\\n  }\\n}\\n\"\n}", hqChaperUniqueId))

	body, err := Request(url, payload)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorOnParseJSON, err.Error()))
	}

	var result pageResponse
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorMarshallJSON, err.Error()))
	}

	return &GetHqPagesResponse{
		Name:     result.Data.GetChapterByID.Hq.Name,
		Number:   result.Data.GetChapterByID.Number,
		Pictures: result.Data.GetChapterByID.Pictures,
	}, nil
}
