package numbers

type Parent struct {
	ObjectIDFieldName string        `json:"objectIdFieldName"`
	UniqueIDField     UniqueIDField `json:"uniqueIdField"`
	GlobalIDFieldName string        `json:"globalIdFieldName"`
	Fields            []Fields      `json:"fields"`
	Features          []Features    `json:"features"`
}
type UniqueIDField struct {
	Name               string `json:"name"`
	IsSystemMaintained bool   `json:"isSystemMaintained"`
}
type Fields struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Alias        string      `json:"alias"`
	SQLType      string      `json:"sqlType"`
	Domain       interface{} `json:"domain"`
	DefaultValue interface{} `json:"defaultValue"`
}
type Attributes struct {
	Value int `json:"value"`
}
type Features struct {
	Attributes Attributes `json:"attributes"`
}

type Summary struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
	Error error  `json:"error"`
}
