package domain

type AuthConfig struct {
	RedirectURL string `json:"redirect_url,omitempty"`
	ClientID    string `json:"client_id,omitempty"`
	State       string `json:"state"`
}

type FTResp struct {
	BasecampID int    `json:"basecamp_id,omitempty"`
	Email      string `json:"email,omitempty"`
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Success    bool   `json:"success,omitempty"`
}

type VerifyResp struct {
	AuthInfo UserOutput `json:"auth_info,omitempty"`
}
type AuthResp struct {
	Token      string     `json:"token"`
	GoogleUser GoogleUser `json:"user_info"`
}

type GoogleUser struct {
	Name          string `json:"name,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	Profile       string `json:"profile,omitempty"`
	Picture       string `json:"picture,omitempty"`
	Email         string `json:"email,omitempty"`
	EmailVerified string `json:"email_verified,omitempty"`
	Locale        string `json:"locale,omitempty"`
}
