package extract

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type (
	response struct {
		Data struct {
			GetHqsByID []struct {
				ID               int         `json:"id"`
				Name             string      `json:"name"`
				Synopsis         string      `json:"synopsis"`
				EditoraID        int         `json:"editoraId"`
				Status           string      `json:"status"`
				PublisherName    string      `json:"publisherName"`
				HqCover          string      `json:"hqCover"`
				ImpressionsCount int         `json:"impressionsCount"`
				Capitulos        []HqChapter `json:"capitulos"`
			} `json:"getHqsById"`
		} `json:"data"`
	}
	HqInfo struct {
		ID               int         `json:"id"`
		Name             string      `json:"name"`
		Synopsis         string      `json:"synopsis"`
		EditoraID        int         `json:"editoraId"`
		Status           string      `json:"status"`
		PublisherName    string      `json:"publisherName"`
		HqCover          string      `json:"hqCover"`
		ImpressionsCount int         `json:"impressionsCount"`
		Capitulos        []HqChapter `json:"capitulos"`
	}
	HqChapter struct {
		Name     string `json:"name"`
		ID       int    `json:"id"`
		Number   string `json:"number"`
		Pictures []struct {
			PictureURL string `json:"pictureUrl"`
		} `json:"pictures"`
	}
)

func GetHqInfo(hqId int) (*HqInfo, error) {
	var result response
	payload := strings.NewReader(fmt.Sprintf("{\n\t\"operationName\": \"getHqsById\",\n\t\"variables\": {\n\t\t\"id\": %v\n\t},\n\t\"query\": \"query getHqsById($id: Int!) {\\n  getHqsById(id: $id) {\\n    id\\n    name\\n    synopsis\\n    editoraId\\n    status\\n    publisherName\\n    hqCover\\n    impressionsCount\\n    capitulos {\\n      name\\n      id\\n      number\\n pictures {\\n      pictureUrl\\n    } }\\n  }\\n}\\n\"\n}", hqId))
	response, err := request(payload)
	if err != nil {
		return nil, errors.New("error in request")
	}

	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorMarshallJSON, err.Error()))
	}
	if len(result.Data.GetHqsByID) > 0 {
		item := result.Data.GetHqsByID[0]
		var chapters []HqChapter
		for _, v := range item.Capitulos {
			chapters = append(chapters, v)
		}

		return &HqInfo{
			ID:               item.ID,
			Name:             item.Name,
			Synopsis:         item.Synopsis,
			EditoraID:        item.EditoraID,
			Status:           item.Status,
			PublisherName:    item.PublisherName,
			HqCover:          item.HqCover,
			ImpressionsCount: item.ImpressionsCount,
			Capitulos:        chapters,
		}, nil
	}
	return nil, errors.New("not found hq")
}
