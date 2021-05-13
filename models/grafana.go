package models

import "time"

//SearchRequest search request
type SearchRequest struct {
	Target string `json:"target"`
}

//SearchResponse search array response
type SearchArrayResponse []string

//SearchMapResponse search map response
type SearchMapResponse []SearchData

//SearchData search data
type SearchData struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

//QueryRequest query request
type QueryRequest struct {
	PanelID int `json:"panelId"`
	Range   struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
		Raw  struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"raw"`
	} `json:"range"`
	RangeRaw struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"rangeRaw"`
	Interval   string `json:"interval"`
	IntervalMs int    `json:"intervalMs"`
	Targets    []struct {
		Target string `json:"target"`
		RefID  string `json:"refId"`
		Type   string `json:"type"`
	} `json:"targets"`
	AdhocFilters []struct {
		Key      string `json:"key"`
		Operator string `json:"operator"`
		Value    string `json:"value"`
	} `json:"adhocFilters"`
	Format        string `json:"format"`
	MaxDataPoints int    `json:"maxDataPoints"`
}

//QueryResponse query resquest
type QueryResponse []QueryData

//Data query result data
type QueryData struct {
	Target     string          `json:"target"`
	Datapoints [][]interface{} `json:"datapoints"`
}
