package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beercli "github.com/CodelyTV/test_project/apihttp/internal"
	"github.com/CodelyTV/test_project/apihttp/internal/errors"
)

const (
	urlProd    = "/v2/beer"
	ontarioURL = "https://api.punkapi.com"
)

type beerRepo struct {
	url string
}

//NewOntarioRepository feth beer data from csv
func NewOntarioRepository() beercli.BeerRepo {
	return &beerRepo{url: ontarioURL}
}

func (b *beerRepo) GetBeers() (beers []beercli.Beer, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, urlProd))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", urlProd)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading response from %s ", urlProd)
	}

	err = json.Unmarshal(content, &beers)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}

	return
}
