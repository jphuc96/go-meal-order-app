package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"git.d.foundation/datcom/backend/src/domain"
)

const baseurl = "https://ke6qcmol32.execute-api.ap-southeast-1.amazonaws.com/stag/verify-email"
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v3/userinfo?access_token="

func (s *Service) FortressVerify(email string) (*domain.FTResp, error) {

	req, err := http.NewRequest("GET", baseurl, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("intern", "intern")
	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusBadRequest {
		return nil, domain.UserNotExistInFT
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var ftresp domain.FTResp
	_ = decoder.Decode(&ftresp)
	return &ftresp, nil
}

func (s *Service) ConfigureOAuth2(clientID, clientSecret, redirectURL string) *oauth2.Config {
	if redirectURL == "" {
		redirectURL = "http://localhost:8000/auth/google/callback"
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func (s *Service) RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func (s *Service) GetUserDataFromGoogle(googleOauthConfig *oauth2.Config, code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	return contents, nil
}

func (s *Service) AuthCheck(r *http.Request) error {
	email := r.Header.Get("email")
	accessToken := r.Header.Get("access_token")
	if email == "" {
		return domain.NotProvideEmail
	}
	if accessToken == "" {
		return domain.NotProvideToken
	}

	exist, _ := s.Store.UserStore.ExistByEmailAndToken(email, accessToken)
	if !exist {
		return domain.VerifyUserFailed
	}

	return nil
}
