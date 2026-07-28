package auth

import (
	"fmt"
	"net/url"
)

func GetHistory(req HistoryRequest) ([]byte, error) {

	client, err := NewFyersClient()
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(
		"data/history?symbol=%s&resolution=%s&date_format=%s&range_from=%s&range_to=%s&cont_flag=%s",
		url.QueryEscape(req.Symbol),
		url.QueryEscape(req.Resolution),
		url.QueryEscape(req.DateFormat),
		url.QueryEscape(req.RangeFrom),
		url.QueryEscape(req.RangeTo),
		url.QueryEscape(req.ContFlag),
	)

	return client.Get(endpoint)
}
