package internal

import "strconv"

//GenericValueUnit formate json
type GenericValueUnit struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

//MashTemp formate json soon Methods
type MashTemp struct {
	Temp     GenericValueUnit `json:"temp"`
	Duration int              `json:"duration"`
}

//Fermentation formate json soon Methods
type Fermentation struct {
	Temp GenericValueUnit `json:"temp"`
}

//Method formate json
type Method struct {
	MashTemp     MashTemp     `json:"mash_temp"`
	Fermentation Fermentation `json:"fermentation"`
	Twist        string       `json:"twist"`
}

//Beer formate json
type Beer struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Tagline          string `json:"tagline"`
	FirstBrewed      string `json:"first_brewed"`
	Description      string `json:"description"`
	ImageURL         string `json:"image_url"`
	Abv              string `json:"code_search_url"`
	Ibu              int    `json:"ibu"`
	TargetFG         int    `json:"target_fg"`
	TargetOG         int    `json:"target_og"`
	Ebc              int    `json:"ebc"`
	Srm              int    `json:"srm"`
	Ph               int    `json:"ph"`
	AttenuationLevel int    `json:"attenuation_level"`
}

//BeerRepo definition of methods to access a data beer
type BeerRepo interface {
	GetBeer() ([]Beer, error)
}

//Gencsv Gencsv Gencsv
func (d *Beer) Gencsv() []string {
	return []string{
		strconv.Itoa(d.ID),
		d.Name,
		d.Tagline,
		d.FirstBrewed,
		d.Description,
		d.ImageURL,
		d.Abv,
	}

}
