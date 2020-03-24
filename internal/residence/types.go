package residence

type StatResidence struct {
	ObjectIDFieldName string           `json:"objectIdFieldName"`
	UniqueIDField     UniqueIDField    `json:"uniqueIdField"`
	GlobalIDFieldName string           `json:"globalIdFieldName"`
	GeometryType      string           `json:"geometryType"`
	SpatialReference  SpatialReference `json:"spatialReference"`
	Fields            []Fields         `json:"fields"`
	Features          []Features       `json:"features"`
}
type UniqueIDField struct {
	Name               string `json:"name"`
	IsSystemMaintained bool   `json:"isSystemMaintained"`
}
type SpatialReference struct {
	Wkid       int `json:"wkid"`
	LatestWkid int `json:"latestWkid"`
}
type Fields struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Alias        string      `json:"alias"`
	SQLType      string      `json:"sqlType"`
	Domain       interface{} `json:"domain"`
	DefaultValue interface{} `json:"defaultValue"`
	Length       int         `json:"length,omitempty"`
}
type Attributes struct {
	Value     int    `json:"value"`
	Residence string `json:"residence"`
}
type Features struct {
	Attributes Attributes `json:"attributes"`
}
