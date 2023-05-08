package main

import (
	"encoding/json"
	"fmt"
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
)

type handler struct {
	*driver.Http
	*driver.GptDriver
}

func newHandler(http *driver.Http, gpt *driver.GptDriver) *handler {
	return &handler{http, gpt}
}

func (h handler) CommentReport(payload string) error {
	m, err := getMessage(payload)
	if err != nil {
		return err
	}

	resp, err := h.GptDriver.RequestMessage(fmt.Sprintf(driver.ReportSystemPromptTemplate, m.Mission, m.Vision, m.Value), m.Body)
	if err != nil {
		return err
	}

	d, err := json.Marshal(entity.ReviewReportRequest{ReviewBody: resp.Choices[0].Message.Content})
	if err != nil {
		return err
	}

	if _, err = h.Http.Put(fmt.Sprintf("http://backend:8080/reports/%s", m.ID), d); err != nil {
		return err
	}
	return nil
}

func getMessage(payload string) (*entity.PubSubMessage, error) {
	b := []byte(payload)
	var msg *entity.PubSubMessage
	if err := json.Unmarshal(b, &msg); err != nil {
		return nil, err
	}
	return msg, nil
}
