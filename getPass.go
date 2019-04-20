package device42_go

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Password struct {
	TotalCount int `json:"total_count"`
	Passwords  []struct {
		Username       string        `json:"username"`
		Category       string        `json:"category"`
		DeviceIds      []int         `json:"device_ids"`
		ViewUsers      string        `json:"view_users"`
		ViewGroups     string        `json:"view_groups"`
		LastPwChange   time.Time     `json:"last_pw_change"`
		Notes          string        `json:"notes"`
		Storage        string        `json:"storage"`
		UseOnlyUsers   string        `json:"use_only_users"`
		Label          string        `json:"label"`
		ViewEditGroups string        `json:"view_edit_groups"`
		FirstAdded     time.Time     `json:"first_added"`
		UseOnlyGroups  string        `json:"use_only_groups"`
		StorageID      int           `json:"storage_id"`
		ViewEditUsers  string        `json:"view_edit_users"`
		Password       string        `json:"password"`
		ID             int           `json:"id"`
		CustomFields   []interface{} `json:"custom_fields"`
	} `json:"Passwords"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

const baseURL string = "https://10.11.12.239/api/1.0"

type Client struct {
	Username string
	Password string
}

func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		username,
		password,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	defaultTransport := http.DefaultTransport.(*http.Transport)

	// Create new Transport that ignores self-signed SSL
	httpClientWithSelfSignedTLS := &http.Transport{
		Proxy:                 defaultTransport.Proxy,
		DialContext:           defaultTransport.DialContext,
		MaxIdleConns:          defaultTransport.MaxIdleConns,
		IdleConnTimeout:       defaultTransport.IdleConnTimeout,
		ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
		TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}

	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{Transport: httpClientWithSelfSignedTLS}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func (s *Client) GetPasswordById(id int) (*Password, error) {
	url := fmt.Sprintf(baseURL+"/passwords/?id=%d&plain_text=yes", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Password
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(bytes))
	return &data, nil
}

func (s *Client) GetPasswordByDevice(device string) (*Password, error) {
	url := fmt.Sprintf(baseURL+"/passwords/?device=%s&plain_text=yes", device)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Password
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(bytes))
	return &data, nil
}
