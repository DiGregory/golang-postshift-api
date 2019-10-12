package postshift

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"

)

type Mail struct {
	Email string
	Key   string
}
type Letter struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Date    string `json:"date"`
	Subject string `json:"subject"`
	From    string `json:"from"`
}


func NewMail(name string, domainName string) (m *Mail, err error) {

	p := url.Values{}

	p.Add("action", "new")
	p.Add("type", "json")
	p.Add("name", name)
	if domainName == "post-shift.ru" || domainName == "postshift.ru" || domainName == "" {
		p.Add("domain", domainName)
	}

	resp, err := SendReq(p)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &m)

	if err != nil {
		return nil, err
	}
	return m, nil

}
func GetMail(key string, id string, p url.Values) (*Letter, error) {
	if p == nil {
		p = url.Values{}
	}
	p.Add("action", "getmail")
	p.Add("type", "json")
	p.Add("id", id)
	p.Add("key", key)
	p.Add("forced","1") //без него текст не виден

	resp, err := SendReq(p)
	if err != nil {
		return nil, err
	}

	var jsonResp *Letter
	if err = json.Unmarshal(resp, &jsonResp); err != nil {

		return nil, err
	}

	return jsonResp, nil

}
func Clear(key string)(clearStatus string,err error){
	p:=url.Values{}
	p.Add("action","clear")
	p.Add("key",key)
	p.Add("type","json")
	resp,err:= SendReq(p)
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
func Update(key string)(lifetime string,err error){
	p := url.Values{}

	p.Add("action", "update")
	p.Add("type", "json")
	p.Add("key", key)


	resp, err := SendReq(p)
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
func Lifetime(key string)(lifetime string,err error){
	p := url.Values{}

	p.Add("action", "livetime")
	p.Add("type", "json")
	p.Add("key", key)


	resp, err := SendReq(p)
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

func GetList(key string) ([]Letter, error) {
	p := url.Values{}

	p.Add("action", "getlist")
	p.Add("type", "json")
	p.Add("key", key)


	resp, err := SendReq(p)
	if err != nil {
		return nil, err
	}

	var jsonResp []Letter
	if err = json.Unmarshal(resp, &jsonResp); err != nil {

		return nil, err
	}

	return jsonResp, nil

}
func DeleteMail(key string) (deleteStatus string, err error) {

	p := url.Values{}

	p.Add("action", "delete")
	p.Add("type", "json")
	p.Add("key", key)

	resp, err := SendReq(p)
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

func SendReq(p url.Values) (response []byte, err error) {
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
	defer resp.Body.Close()

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respJSON, nil
}

