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

//QueryResponse query response
type QueryResponse []QueryData

//Data query result data
type QueryData struct {
	Target     string          `json:"target"`
	Datapoints [][]interface{} `json:"datapoints"`
}

//QueryTableResponse query table response
type QueryTableResponse []QueryTableData

//QueryTableData query table data
type QueryTableData struct {
	Columns []TableColumn   `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
	Type    string          `json:"type"`
}

//TableColumn table column
type TableColumn struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

//AnnotationRequest annotation request
type AnnotationRequest struct {
	Range struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	} `json:"range"`
	RangeRaw struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"rangeRaw"`
	Annotation Annotation `json:"annotation"`
}

//Annotation annotation
type Annotation struct {
	Name string `json:"name"`

	Datasource string `json:"datasource"`
	IconColor  string `json:"iconColor"`
	Enable     bool   `json:"enable"`
	ShowLine   bool   `json:"showLine"`
	Query      string `json:"query"`
}

//AnnotationResponse annotation response
type AnnotationResponse struct {
	// The original annotation sent from Grafana.
	Annotation Annotation `json:"annotation"`
	// Time since UNIX Epoch in milliseconds. (required)
	Time int64 `json:"time"`
	// The title for the annotation tooltip. (required)
	Title string `json:"title"`
	// Tags for the annotation. (optional)
	Tags string `json:"tags"`
	// Text for the annotation. (optional)
	Text string `json:"text"`
}
