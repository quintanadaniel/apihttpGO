package internal_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	beerscli "github.com/CodelyTV/test_project/apihttp/cli"
	errorBeer "github.com/CodelyTV/test_project/apihttp/errors"
	"github.com/CodelyTV/test_project/apihttp/internal/mock"
)

// BeerRepo definiton of methods to access a data beer
type BeerRepo interface {
	GetBeers() ([]Beer, error)
}

func MockBeers() (b []beerscli.Beer) {
	return []beerscli.Beer{
		beerscli.NewBeersRows(1, "pilsen", "test", "beer test", "http://beertest.com"),
	}
}

func TestRunBeers(t *testing.T) {
	repo := new(mock.MyMock)
	repo.On("GetBeersAll", "prueba", 1).Return(MockBeers())
	beers := repo.GetBeersAll("prueba", 1)
	fmt.Println(beers)
	assert.NotNil(t, beers)
}

func TestFetchByID(t *testing.T) {
	tests := map[string]struct {
		repo  beerscli.Beer
		input int
		want  int
		err   error
	}{
		"valid beer":      {repo: buildMockBeers(), input: 1, want: 1, err: nil},
		"not found beer":  {repo: buildMockBeers(), input: 999999, err: errorBeer.NewDataUnreacheable("error")},
		"error with beer": {repo: buildMockBeers(), err: errorBeer.NewDataUnreacheable("error")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T)) {
			service := N
		}
	}
}
