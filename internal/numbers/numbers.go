package numbers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetData() (*[]Summary, error) {

	cat := []string{"tests", "deaths", "confirmed", "recovered", "PUIs", "PUMs"}
	ch := make(chan Summary, len(cat))

	for _, ca := range cat {
		go func(j string, c chan Summary) {
			d, err := getData(j)
			c <- Summary{Count: d, Type: j, Error: err}
		}(ca, ch)
	}

	l, totalChan := []Summary{}, len(cat)

	for x := 0; totalChan > x; x++ {
		l = append(l, <-ch)
	}

	for _, r := range l {
		if r.Error != nil {
			return nil, r.Error
		}
	}

	return &l, nil
}

func getData(e string) (total int, err error) {
	total = 0
	url := "https://services5.arcgis.com/mnYJ21GiFTR97WFg/arcgis/rest/services/slide_fig/FeatureServer/0/query?f=json&where=1%3D1&returnGeometry=false&spatialRel=esriSpatialRelIntersects&outFields=%2A&outStatistics=%5B%7B%22statisticType%22%3A%22sum%22%2C%22onStatisticField%22%3A%22" + e + "%22%2C%22outStatisticFieldName%22%3A%22value%22%7D%5D&cacheHint=true"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	p := Parent{}

	if err != nil {
		return
	}

	err = json.Unmarshal(body, &p)

	if err != nil {
		return
	}

	for _, s := range (*&p).Features {
		return s.Attributes.Value, nil
	}
	return 0, nil
}
