package hospital

type StatHospital struct {
	ObjectIDFieldName     string           `json:"objectIdFieldName"`
	UniqueIDField         UniqueIDField    `json:"uniqueIdField"`
	GlobalIDFieldName     string           `json:"globalIdFieldName"`
	GeometryType          string           `json:"geometryType"`
	SpatialReference      SpatialReference `json:"spatialReference"`
	Fields                []Fields         `json:"fields"`
	ExceededTransferLimit bool             `json:"exceededTransferLimit"`
	Features              []Features       `json:"features"`
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
	Length       int64       `json:"length,omitempty"`
	Domain       interface{} `json:"domain"`
	DefaultValue interface{} `json:"defaultValue"`
}
type Attributes struct {
	Facility  string  `json:"facility"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Count     int     `json:"count_"`
	ObjectID  int     `json:"ObjectId"`
}
type Features struct {
	Attributes Attributes `json:"attributes"`
}
