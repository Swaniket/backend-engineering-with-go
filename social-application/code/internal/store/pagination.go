package store

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaginatedFeedQuery struct {
	Limit  int      `json:"limit" validate:"gte=1,lte=20"` // Greater than or equal to 1 & less than equal to 20
	Offset int      `json:"offset" validate:"gte=0"`       // Greater than or equal to 0
	Sort   string   `json:"sort" validate:"oneof=asc desc"`
	Tags   []string `json:"tags" validate:"max=5"` // Search by multiple tags
	Search string   `json:"search" validate:"max=100"`
	Since  string   `json:"since"`
	Until  string   `json:"until"`
}

func (fq PaginatedFeedQuery) ParsePaginationRequest(r *http.Request) (PaginatedFeedQuery, error) {
	queryString := r.URL.Query()

	// Parsing of limit
	limit := queryString.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return fq, nil
		}

		fq.Limit = l
	}

	// Parsing of Offset
	offset := queryString.Get("offset")
	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return fq, nil
		}

		fq.Offset = o
	}

	// Parsing of Sort
	sort := queryString.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}

	// Parsing of tags
	tags := queryString.Get("tags")
	if tags != "" {
		fq.Tags = strings.Split(tags, ",")
	}

	search := queryString.Get("search")
	if search != "" {
		fq.Search = search
	}

	since := queryString.Get("since")
	if since != "" {
		fq.Since = parseTime(since)
	}

	until := queryString.Get("until")
	if until != "" {
		fq.Until = parseTime(until)
	}

	return fq, nil
}

func parseTime(s string) string {
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return ""
	}

	return t.Format(time.DateTime)
}
