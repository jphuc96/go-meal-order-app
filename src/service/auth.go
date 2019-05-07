package service

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"git.d.foundation/datcom/backend/src/domain"
)

const baseurl = "https://ke6qcmol32.execute-api.ap-southeast-1.amazonaws.com/stag/verify-email"
const tokenInfoURL = "https://oauth2.googleapis.com/tokeninfo?id_token="

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

func (s *Service) RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func (s *Service) AuthCheck(r *http.Request) error {
	accessToken := r.Header.Get("authorization")
	if accessToken == "" {
		return domain.NotProvideAccessToken
	}

	exist, _ := s.Store.UserStore.ExistByToken(accessToken)
	if !exist {
		return domain.VerifyUserFailed
	}

	return nil
}

func (s *Service) GetGoogleUserInfo(idToken string) (*domain.GoogleUser, error) {
	res, err := http.Get(tokenInfoURL + idToken)
	var gu domain.GoogleUser
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(res.Body)
	err = d.Decode(&gu)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &gu, nil
}
