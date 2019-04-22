package domain

type AuthConfig struct {
	RedirectURI string `json:"redirect_uri,omitempty"`
	ClientID    string `json:"client_id,omitempty"`
}

type FTResp struct {
	BasecampID int    `json:"basecamp_id,omitempty"`
	Email      string `json:"email,omitempty"`
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Success    bool   `json:"success,omitempty"`
}

type VerifyResp struct {
	AuthInfo CreateUserInput `json:"auth_info,omitempty"`
}
