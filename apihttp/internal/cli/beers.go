package cli

import (
	"log"

	"github.com/CodelyTV/test_project/apihttp/internal/geting"

	"github.com/spf13/cobra"
)

//CobraFn function definition of Run comand cobra
type CobraFn func(cmd *cobra.Command, args []string)

const (
	idFlag       = "id"
	nameFileFlag = "nameFile"
)

//InitBeer param for console
func InitBeer(service geting.Service) *cobra.Command {
	beerCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print all beers",
		Run:   runBeers(services),
	}

	beerCmd.Flags().StringP(idFlag, "i", "", "Id of the beer")
	beerCmd.Flags().StringP(nameFileFlag, "c", "", "Name CSV File")

	return beerCmd
}

func runBeers(service geting.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetString(idFlag)
		nameFile, _ := cmd.Flags().GetString(nameFileFlag)

		if id != "" && nameFile != "" {
			beer, err = service.GetBeersID(id)
			if err != nil {
				log.Fatal(err)
			}

		} else {
			beers, err = service.GetBeersAll()
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
