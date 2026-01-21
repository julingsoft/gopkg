package xlog

type LogData struct {
	Method   string
	Url      string
	Headers  map[string]string
	Request  string
	Response interface{}
	Status   int
	Duration int64
}

type LogEntry struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Time    string                 `json:"time"`
	TraceID string                 `json:"trace_id,omitempty"`
	Extra   map[string]interface{} `json:"extra,omitempty"`
}
