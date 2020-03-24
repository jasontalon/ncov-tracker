package individual

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetData() (*[]Attributes, error) {
	url := "https://services5.arcgis.com/mnYJ21GiFTR97WFg/arcgis/rest/services/PH_masterlist/FeatureServer/0/query?f=json&where=1%3D1&outFields=%2A"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	t := StatIndividual{}

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(*&body, &t)
	if err != nil {
		return nil, err
	}

	att := []Attributes{}
	for _, s := range *&t.Features {
		att = append(att, s.Attributes)
	}

	return &att, nil
}
