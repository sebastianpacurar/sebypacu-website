package API

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchCountries(location string) ([]byte, error) {
	URL := fmt.Sprintf("https://restcountries.eu/rest/v2/%s", location)

	res, err := http.Get(URL)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Eroare la http.Get(\"%s\")", URL), err.Error())
		return []byte{}, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatalln("Eroare la defer func pe res.Body.Close()", err.Error())
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Eroare la procesarea res.Body in byte[]", err.Error())
		return []byte{}, err
	}

	return body, nil
}
