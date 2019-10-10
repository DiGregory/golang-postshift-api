package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"
)

type mail struct {
	Email string
	Key   string
}
type letter struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Date    string `json:"date"`
	Subject string `json:"subject"`
	From    string `json:"from"`
}


func newMail(name string, domainName string) (m *mail, err error) {

	p := url.Values{}

	p.Add("action", "new")
	p.Add("type", "json")
	p.Add("name", name)
	if domainName == "post-shift.ru" || domainName == "postshift.ru" || domainName == "" {
		p.Add("domain", domainName)
	}

	resp, err := sendReq(p)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &m)

	if err != nil {
		return nil, err
	}
	return m, nil

}
func getMail(key string, id string, p url.Values) (*letter, error) {
	if p == nil {
		p = url.Values{}
	}
	p.Add("action", "getmail")
	p.Add("type", "json")
	p.Add("id", id)
	p.Add("key", key)
	p.Add("forced","1") //без него текст не виден

	resp, err := sendReq(p)
	if err != nil {
		return nil, err
	}

	var jsonResp *letter
	if err = json.Unmarshal(resp, &jsonResp); err != nil {

		return nil, err
	}

	return jsonResp, nil

}
func clear(key string)(clearStatus string,err error){
	p:=url.Values{}
	p.Add("action","clear")
	p.Add("key",key)
	p.Add("type","json")
	resp,err:=sendReq(p)
	if err!=nil{
		return "",nil
	}
	var jsonResp map[string]string
	if err = json.Unmarshal(resp, &jsonResp); err != nil {

		return "", err
	}
	if jsonResp["clear"]!="ok"{
		jsonResp["clear"]="badKey"
	}

	return jsonResp["clear"], nil

}
func update(key string)(lifetime string,err error){
	p := url.Values{}

	p.Add("action", "update")
	p.Add("type", "json")
	p.Add("key", key)


	resp, err := sendReq(p)
	if err != nil {
		return "", err
	}

	var jsonResp map[string]string
	if err = json.Unmarshal(resp, &jsonResp); err != nil {

		return "", err
	}
	if _,ok:=jsonResp["error"];ok{
		return jsonResp["error"],nil
	}

	return jsonResp["livetime"], nil
}
func lifetime(key string)(lifetime string,err error){
	p := url.Values{}

	p.Add("action", "livetime")
	p.Add("type", "json")
	p.Add("key", key)


	resp, err := sendReq(p)
	if err != nil {
		return "", err
	}

	var jsonResp map[string]string
	if err = json.Unmarshal(resp, &jsonResp); err != nil {

		return "", err
	}
	if _,ok:=jsonResp["error"];ok{
		return jsonResp["error"],nil
	}

	return jsonResp["livetime"], nil
}

func getList(key string) ([]letter, error) {
	p := url.Values{}

	p.Add("action", "getlist")
	p.Add("type", "json")
	p.Add("key", key)


	resp, err := sendReq(p)
	if err != nil {
		return nil, err
	}

	var jsonResp []letter
	if err = json.Unmarshal(resp, &jsonResp); err != nil {

		return nil, err
	}

	return jsonResp, nil

}
func deleteMail(key string) (deleteStatus string, err error) {

	p := url.Values{}

	p.Add("action", "delete")
	p.Add("type", "json")
	p.Add("key", key)

	resp, err := sendReq(p)
	if err != nil {
		return "", err
	}

	var jsonResp map[string]string
	if err = json.Unmarshal(resp, &jsonResp); err != nil {

		return "", err
	}
	if jsonResp["delete"]!="ok"{
		jsonResp["delete"]="badKey"
	}

	return jsonResp["delete"], nil

}

func sendReq(p url.Values) (response []byte, err error) {
	myUrl, err := url.Parse("https://post-shift.ru/api.php?")
	if err != nil {
		return nil, err
	}

	if p == nil {
		p = url.Values{}
	}
	myUrl.RawQuery = p.Encode()
	resp, err := http.PostForm(myUrl.String(), p)
	if err != nil {
		return nil, err
	}

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respJSON, nil
}


