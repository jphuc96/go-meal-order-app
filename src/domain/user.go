package domain

type CreateUserInput struct {
	Name        string `json:"name,omitempty"`
	GoogleID    string `json:"google_id,omitempty"`
	Email       string `json:"email,omitempty"`
	Token       string `json:"token,omitempty"`
	DeviceToken string `json:"device_token,omitempty"`
	DeviceType  string `json:"device_type,omitempty"`
}
