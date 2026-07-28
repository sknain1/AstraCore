package auth

type HistoryRequest struct {
	Symbol     string `json:"symbol"`
	Resolution string `json:"resolution"`
	DateFormat string `json:"date_format"`
	RangeFrom  string `json:"range_from"`
	RangeTo    string `json:"range_to"`
	ContFlag   string `json:"cont_flag"`
}
