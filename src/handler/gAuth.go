package handler

import (
	"encoding/json"
	"net/http"

	"git.d.foundation/datcom/backend/src/domain"

	"github.com/BurntSushi/toml"
)

func (c *CoreHandler) GetGoogleLoginURL(w http.ResponseWriter, r *http.Request) {

	var conf domain.AuthConfig
	_, err := toml.DecodeFile("config/gAuth.toml", &conf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errresp := domain.CreateErrorResp{}
		errresp.Error.Code = "FileNotFound"
		errresp.Error.Message = "No such file or directory"
		resp, _ := json.Marshal(errresp)
		w.Write(resp)
		return
	}

	js, _ := json.Marshal(conf)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func (c *CoreHandler) VerifyGoogleUserLogin(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user domain.CreateUserInput
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errresp := domain.CreateErrorResp{}
		errresp.Error.Code = "DecodeError"
		errresp.Error.Message = "Can't decode struct to Json"
		resp, _ := json.Marshal(errresp)
		w.Write(resp)
		return
	}

	success := fortressVerify(&user)
	if !success {
		w.WriteHeader(http.StatusBadRequest)
		errresp := domain.CreateErrorResp{}
		errresp.Error.Code = "LoginFail"
		errresp.Error.Message = "Validation Fail"
		resp, _ := json.Marshal(errresp)
		w.Write(resp)
		return
	} else {
		var verifyResp domain.VerifyResp
		verifyResp.AuthInfo = user
		resp, _ := json.Marshal(&verifyResp)
		w.Write(resp)
	}
}

const baseurl = "https://ke6qcmol32.execute-api.ap-southeast-1.amazonaws.com/stag/verify-email"

func fortressVerify(u *domain.CreateUserInput) bool {

	req, _ := http.NewRequest("GET", baseurl, nil)
	req.SetBasicAuth("intern", "intern")
	q := req.URL.Query()
	q.Add("email", u.Email)
	req.URL.RawQuery = q.Encode()
	// fmt.Println(req.URL.String())
	client := &http.Client{}
	resp, _ := client.Do(req)

	decoder := json.NewDecoder(resp.Body)
	var ftresp domain.FTResp
	_ = decoder.Decode(&ftresp)
	return ftresp.Success
}
