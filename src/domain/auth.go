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
	AuthInfo CreateaUserOutput `json:"auth_info,omitempty"`
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

type AuthResp struct {
	Token          string         `json:"token"`
	GoogleUserInfo GoogleUserInfo `json:"user_info"`
}
