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

	"git.d.foundation/datcom/backend/src/domain"
)

const baseurl = "https://ke6qcmol32.execute-api.ap-southeast-1.amazonaws.com/stag/verify-email"
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

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

	if resp.StatusCode == 400 {
		return nil, domain.UserNotExistInFT
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var ftresp domain.FTResp
	_ = decoder.Decode(&ftresp)
	return &ftresp, nil
}

func (s *Service) RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func (s *Service) GetUserDataFromGoogle(googleOauthConfig *oauth2.Config, code string) (*domain.GoogleUserInfo, error) {
	// Use code to get token and get user info from Google.

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

	var userData domain.GoogleUserInfo
	json.Unmarshal(contents, &userData)
	return &userData, nil
}