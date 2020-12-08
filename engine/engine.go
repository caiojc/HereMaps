package engine

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//GetMapWithPolygon : Retorna o mapa com o Polygon desenhado
func GetMapWithPolygon(output string, geometry string) {
	endpoint, _ := url.Parse("https://image.maps.api.here.com/mia/1.6/region")
	queryParams := endpoint.Query()
	queryParams.Set("app_id", "APP-ID-HERE")
	queryParams.Set("app_code", "APP-CODE-HERE")
	queryParams.Set("ppi", "320")
	queryParams.Set("w", "1280")
	queryParams.Set("h", "720")
	queryParams.Set("z", "11")
	queryParams.Set("a0", geometry)
	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		f, _ := os.Create(output)
		data, _ := ioutil.ReadAll(response.Body)
		f.Write(data)
		defer f.Close()
	}
}
