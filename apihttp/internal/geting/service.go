package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	beercli "github.com/CodelyTV/test_project/apihttp/internal/cli"
)

const (
	idFlag       = "id"
	nameFileFlag = "nameFile"
	urlBeer      = "https://api.punkapi.com/v2/beers"
)

//Service provider find the beers
type Service interface {
	GetBeersAll() []beercli.Beer

	GetBeersID(id string) beercli.Beer
}

type service struct {
	bR beercli.BeerRepo
}

//NewService creates and adding service with the necesary dependens
func NewService(r beercli.BeerRepo) Service {
	return &service{r}
}

//GetBeersAll function beers
func (s *service) GetBeersAll() []beercli.Beer {
	var beers []beercli.Beer

	resp, _ := http.Get(urlBeer)

	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &beers)

	return s.bR.GetBeers
}

//GetBeersID function that find id beers
func (s *service) GetBeersID(idbeers string) []beercli.Beer {
	var idsbeers []beercli.Beer
	resp, _ := http.Get(urlBeer + "/" + idbeers)

	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &idsbeers)

	return beercli.Beer
}
