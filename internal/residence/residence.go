package residence

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetData() (*[]Features, error) {

	url := "https://services5.arcgis.com/mnYJ21GiFTR97WFg/arcgis/rest/services/PH_masterlist/FeatureServer/0/query?f=json&returnGeometry=false&spatialRel=esriSpatialRelIntersects&outFields=%2A&groupByFieldsForStatistics=residence&orderByFields=value%20desc&outStatistics=%5B%7B%22statisticType%22%3A%22count%22%2C%22onStatisticField%22%3A%22FID%22%2C%22outStatisticFieldName%22%3A%22value%22%7D%5D&cacheHint=true"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	t := StatResidence{}
	err := json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	return &t.Features, nil
}

func total(t *[]Features) int {
	s := 0
	for _, o := range *t {
		s += o.Attributes.Value
	}
	return s
}
