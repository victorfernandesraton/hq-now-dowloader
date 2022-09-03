package extract

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type chapterResponse struct {
	Data struct {
		GetHqsByID []struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Synopsis         string `json:"synopsis"`
			EditoraID        int    `json:"editoraId"`
			Status           string `json:"status"`
			PublisherName    string `json:"publisherName"`
			HqCover          string `json:"hqCover"`
			ImpressionsCount int    `json:"impressionsCount"`
			Capitulos        []struct {
				Name   string `json:"name"`
				ID     int    `json:"id"`
				Number string `json:"number"`
			} `json:"capitulos"`
		} `json:"getHqsById"`
	} `json:"data"`
}

type GetHqChapterResponse struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Number string `json:"number"`
}

func GetHqChapter(hqId int) ([]*GetHqChapterResponse, error) {

	url := "https://admin.hq-now.com/graphql"

	payload := strings.NewReader(fmt.Sprintf("{\n\t\"operationName\": \"getHqsById\",\n\t\"variables\": {\n\t\t\"id\": %v\n\t},\n\t\"query\": \"query getHqsById($id: Int!) {\\n  getHqsById(id: $id) {\\n    id\\n    name\\n    synopsis\\n    editoraId\\n    status\\n    publisherName\\n    hqCover\\n    impressionsCount\\n    capitulos {\\n      name\\n      id\\n      number\\n    }\\n  }\\n}\\n\"\n}", hqId))

	body, err := Request(url, payload)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorOnParseJSON, err.Error()))
	}

	var result chapterResponse
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorMarshallJSON, err.Error()))
	}

	var response []*GetHqChapterResponse
	for _, hqs := range result.Data.GetHqsByID {
		for _, cap := range hqs.Capitulos {
			response = append(response, &GetHqChapterResponse{
				Name:   cap.Name,
				ID:     cap.ID,
				Number: cap.Number,
			})
		}
	}
	return response, nil
}
