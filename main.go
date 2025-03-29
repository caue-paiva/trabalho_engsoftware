package main

import (
	"fmt"
	"os"

	orcidapi "orcidapi/apimodule"
)

func main() {
	const jsonSecretsPath = "orcidsecrets.json"

	//orcid id secrets stored in json file
	secrets, err := orcidapi.GetAcessToken(jsonSecretsPath)
	if err != nil {
		fmt.Printf("failure to fetch client secrets (id,secret and access token), error: %v", err)
		os.Exit(1)
	}

	testId := "0009-0007-8094-7155"
	apiHandler := orcidapi.ApiHandler{AccessToken: secrets.AccessToken}
	record := apiHandler.NewRecord(testId)

	data, err := record.Person()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%v", data)

}
