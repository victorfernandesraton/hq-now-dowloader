package extract

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type (
	chapterResponse struct {
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
	GetHqChapterResponse struct {
		Name   string `json:"name"`
		ID     int    `json:"id"`
		Number string `json:"number"`
	}
	GetHqResponse struct {
		Name     string `json:"name"`
		ID       int    `json:"id"`
		Cover    string `json:"cover"`
		Chapters []GetHqChapterResponse
	}
)

func GetHqChapter(hqId int) (*GetHqResponse, error) {

	payload := strings.NewReader(fmt.Sprintf("{\n\t\"operationName\": \"getHqsById\",\n\t\"variables\": {\n\t\t\"id\": %v\n\t},\n\t\"query\": \"query getHqsById($id: Int!) {\\n  getHqsById(id: $id) {\\n    id\\n    name\\n    synopsis\\n    editoraId\\n    status\\n    publisherName\\n    hqCover\\n    impressionsCount\\n    capitulos {\\n      name\\n      id\\n      number\\n    }\\n  }\\n}\\n\"\n}", hqId))

	body, err := Request(payload)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorOnParseJSON, err.Error()))
	}

	var result chapterResponse
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorMarshallJSON, err.Error()))
	}

	var resposne GetHqResponse
	for _, hqs := range result.Data.GetHqsByID {
		if hqs.ID == hqId {
			resposne.Name = hqs.Name
			resposne.ID = hqs.ID
			resposne.Cover = hqs.HqCover
		}
		for _, cap := range hqs.Capitulos {
			resposne.Chapters = append(resposne.Chapters, GetHqChapterResponse{
				Name:   cap.Name,
				ID:     cap.ID,
				Number: cap.Number,
			})
		}
	}
	return &resposne, nil
}
