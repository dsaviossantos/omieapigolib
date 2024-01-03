package main

import (
	"fmt"
	"os"

	"github.com/dsaviossantos/omieapigolib/pkg/omie_api"
)

func main() {
	omie_api.Config.SetCredentials(os.Getenv("OMIE_API_KEY"), os.Getenv("OMIE_API_SECRET"))
	r := omie_api.Get.ListCompanies()
	fmt.Println(string(r))
}
