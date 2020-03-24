package hospital

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetData() (*[]Attributes, error) {
	url := "https://services5.arcgis.com/mnYJ21GiFTR97WFg/arcgis/rest/services/conf_fac_tracking/FeatureServer/0/query?f=json&where=1%3D1&returnGeometry=false&spatialRel=esriSpatialRelIntersects&outFields=%2A&orderByFields=count_%20desc&resultOffset=0&resultRecordCount=50&cacheHint=true"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	s := StatHospital{}

	err = json.Unmarshal(body, &s)

	if err != nil {
		return nil, err
	}

	att := []Attributes{}

	for _, f := range *&s.Features {
		att = append(att, f.Attributes)
	}

	return &att, nil
}
