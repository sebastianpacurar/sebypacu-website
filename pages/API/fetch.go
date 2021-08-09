package API

import (
	"io/ioutil"
	"log"
	"net/http"
)

func FetchData(location string) []byte {
	URL := "https://restcountries.eu/rest/v2/"

	resp, err := http.Get(URL + location)
	if err != nil {
		log.Fatalln(err)
	}


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return body
}
