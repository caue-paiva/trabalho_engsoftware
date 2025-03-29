package orcidapi

type ApiHandler struct {
	AccessToken string
}

func (api ApiHandler) NewRecord(orcidID string) RecordCaller {
	return RecordCaller{
		orcidID:     orcidID,
		accessToken: api.AccessToken,
	}
}
