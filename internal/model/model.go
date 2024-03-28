package model

import (
	m "github.com/georgedinicola/jeopardy-data-scraper/model"
)

type JeopardyGameBoxScore struct {
	m.JeopardyGameBoxScore
}

type Episode struct {
	m.Episode
}

type Contestant struct {
	m.Contestant
}

type PaginationMetaData struct {
	TotalRecords int  `json:"total_records"`
	CurrentPage  int  `json:"current_page"`
	TotalPages   int  `json:"total_pages"`
	NextPage     *int `json:"next_page,omitempty"`
	PrevPage     *int `json:"prev_page,omitempty"`
}
