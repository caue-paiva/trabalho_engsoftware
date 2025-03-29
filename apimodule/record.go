package orcidapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RecordCaller struct {
	orcidID     string
	accessToken string
}

type RecordPerson struct {
}

// formats url and request headers for API GET call, public methods then only provide the specific endpoint for the records API
func (rec RecordCaller) makeCall(endpoint string) (*http.Response, error) {
	url := fmt.Sprintf("https://pub.orcid.org/v3.0/%s/%s", rec.orcidID, endpoint)
	header := http.Header{}
	header.Set("Accept", "application/vnd.orcid+json")
	header.Set("Authorization", fmt.Sprintf("Bearer %s", rec.accessToken))

	return makeGetRequest(url, header)
}

// get full information from an orcid record (minimal parsing)
func (rec RecordCaller) GetFull() (map[string]any, error) {
	resp, err := rec.makeCall("record")
	if err != nil {
		return nil, fmt.Errorf("error making ORCID API call: %w", err)
	}

	//this func streams API body response to the JSON file

	var data map[string]any
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return data, nil
}

func (rec RecordCaller) TempSaveOnFile() {
	resp, err := rec.makeCall("record")
	if err != nil {
		//return nil, fmt.Errorf("error making ORCID API call: %w", err)
	}

	//this func streams API body response to the JSON file

	//create file
	file, err := os.Create(fmt.Sprintf("%s_record.json", rec.orcidID))
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	//streams response body data to new file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error writing response to file: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Successfully saved record to %s_record.json\n", rec.orcidID)
}
