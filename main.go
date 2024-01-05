package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dsaviossantos/omieapigolib/pkg/omie_api"
)

func main() {
	// [x] Set your credentials
	// [] Rate limit setups
	// [] Verify RL permissions
	// [] Make request
	// [] True: Make request | False: Wait then make request
	// [] Receive response
	// [] Parse response
	// [] Error: Handle OmieError | Success: Continue
	// [] ErrorHandler: What to do next? Wait and try again? Return error?
	// [] Success: Return response

	start := time.Now()
	key := os.Getenv("OMIE_API_KEY")
	secret := os.Getenv("OMIE_API_SECRET")
	omie_api.Config.SetCredentials(key, secret)
	r := omie_api.Get.ListCompanies(1, 50)
	fmt.Println(time.Since(start))
	fmt.Println(r.EmpresasCadastro[0].CodigoEmpresa)
}
