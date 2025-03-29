package orcidapi

import (
	"encoding/json"
	"fmt"
)

type RecordPerson struct {
	GivenName  string
	FamilyName string
	Biography  string
	Emails     []string
	Addresses  []string
	KeyWords   []string
}

// struct to represent a generic affiliation to some organization according to the orcid API format
type RecordAffiliation struct {
	OrgName    string
	Role       string
	StartYear  int
	StartMonth int
	EndYear    int
	EndMonth   int
	OrgCity    string
	OrgRegion  string
	OrgCountry string
}

type RecordActivities struct {
	DistinctionsAffiliations []string
	EducationAffiliations    []string
}

// Implement the json.Unmarshaler interface for RecordPerson
func (rp *RecordPerson) UnmarshalJSON(data []byte) error {
	// Define an auxiliary struct *within* the method
	// This struct mirrors the JSON structure needed to extract fields
	type JSONValue struct {
		Value string `json:"value"`
	}
	type JSONName struct {
		GivenNames JSONValue `json:"given-names"`
		FamilyName JSONValue `json:"family-name"`
	}
	type JSONListEntry struct {
		Value string `json:"value"`
	}
	type JSONEmails struct {
		Email []JSONListEntry `json:"email"`
	}
	type JSONAddresses struct {
		Address []JSONListEntry `json:"address"`
	}
	type JSONKeywords struct {
		Keyword []JSONListEntry `json:"keyword"`
	}
	type JSONBiography struct {
		Biography string `json:"content"`
	}

	type JSONPerson struct {
		Name      JSONName       `json:"name"`
		Biography *JSONBiography `json:"biography"` // Use pointer for null handling
		Emails    JSONEmails     `json:"emails"`
		Addresses JSONAddresses  `json:"addresses"`
		Keywords  JSONKeywords   `json:"keywords"`
	}
	// Use an alias type to avoid recursion during unmarshalling
	// Define a temporary root object containing the 'person' key
	var Person JSONPerson
	// Unmarshal the raw JSON data into the auxiliary structure
	if err := json.Unmarshal(data, &Person); err != nil {
		return fmt.Errorf("error unmarshalling intermediate structure: %w", err)
	}

	// Map data from the auxiliary struct to the RecordPerson fields (rp)
	rp.GivenName = Person.Name.GivenNames.Value
	rp.FamilyName = Person.Name.FamilyName.Value

	if Person.Biography != nil {
		rp.Biography = Person.Biography.Biography // Access the content field
	} else {
		rp.Biography = "" // Or default value
	}

	//set list fields
	rp.Emails = make([]string, 0, len(Person.Emails.Email))
	for _, entry := range Person.Emails.Email {
		rp.Emails = append(rp.Emails, entry.Value)
	}

	rp.Addresses = make([]string, 0, len(Person.Addresses.Address))
	for _, entry := range Person.Addresses.Address {
		rp.Addresses = append(rp.Addresses, entry.Value)
	}

	rp.KeyWords = make([]string, 0, len(Person.Keywords.Keyword))
	for _, entry := range Person.Keywords.Keyword {
		rp.KeyWords = append(rp.KeyWords, entry.Value)
	}

	return nil // Return nil on success
}
