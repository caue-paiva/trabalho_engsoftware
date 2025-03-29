package orcidapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const orcidApiSecretsPath = "https://orcid.org/oauth/token"

type Secrets struct {
	ClientId     string
	ClientSecret string
	AccessToken string
}

func readEnvSecrets(jsonFilePath string) (string, string, error) {
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return "", "", errors.New("failed to open json file with Secrets")
	}

	var apiSecrets struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}

	err = json.Unmarshal(jsonData, &apiSecrets)
	if err != nil {
		return "", "", fmt.Errorf("failed to unmarshal json into Secrets struct, error: %v", err)
	}

	return apiSecrets.ClientId, apiSecrets.ClientSecret, nil
}

// given
func GetAcessToken(jsonFilePath string) (Secrets, error) {
	id, secret, err := readEnvSecrets(jsonFilePath)
	if err != nil {
		return Secrets{}, fmt.Errorf("failed to load secrets from json file, error: %v", err)
	}

	// Create url.Values for form-urlencoded data
	formValues := url.Values{
		"client_id":     []string{id},
		"client_secret": []string{secret},
		"grant_type":    []string{"client_credentials"},
		"scope":         []string{"/read-public"},
	}

	// Use MakePostRequest with form-urlencoded content type
	resp, err := makePostRequest(orcidApiSecretsPath, "application/x-www-form-urlencoded", []byte(formValues.Encode()))
	if err != nil {
		return Secrets{},fmt.Errorf("error sending access token request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Secrets{}, fmt.Errorf("api request failed with status code: %d, status: %s", resp.StatusCode, resp.Status)
	}

	// Read response body
	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return Secrets{}, fmt.Errorf("error reading access token response: %v", err)
	}

	accessToken := result["access_token"].(string)

	return Secrets{
		ClientId: id,
		ClientSecret: secret,
		AccessToken: accessToken,
	}, nil
}
