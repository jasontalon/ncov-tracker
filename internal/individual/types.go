package individual

type StatIndividual struct {
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
	FID        int         `json:"FID"`
	FID1       interface{} `json:"FID_1"`
	PHMasterl  string      `json:"PH_masterl"`
	Edad       interface{} `json:"edad"`
	Kasarian   string      `json:"kasarian"`
	Nationalit string      `json:"nationalit"`
	Residence  string      `json:"residence"`
	TravelHx   string      `json:"travel_hx"`
	Symptoms   interface{} `json:"symptoms"`
	Confirmed  string      `json:"confirmed"`
	Facility   string      `json:"facility"`
	Latitude   float64     `json:"latitude"`
	Longitude  float64     `json:"longitude"`
	Status     interface{} `json:"status"`
	EpiLink    interface{} `json:"epi_link"`
	Petsa      interface{} `json:"petsa"`
}
type Features struct {
	Attributes Attributes `json:"attributes"`
}
