package csv

import (
	"encoding/csv"
	"os"

	beercli "github.com/CodelyTV/test_project/apihttp/internal"
	errorsBeer "github.com/CodelyTV/test_project/apihttp/internal/errors"
)

//FileCsv create, read, write feli csv
func FileCsv(beers []beercli.Beer, nameFile string) error {

	//create csv
	csvFile, err := os.Create("./data/" + nameFile + ".csv")
	if err != nil {
		return errorsBeer.WrapDataUnreacheable(err, "Error Creating File CSV")

	}

	//Initialize csv write
	csvWrite := csv.NewWriter(csvFile)

	//Write all the records
	for _, beer := range beers {
		csvWrite.Write(beer.Gencsv())

	}

	csvWrite.Flush()
	csvFile.Close()

	return errorsBeer.WrapDataUnreacheable(err, "Error function FileCsv")
}
