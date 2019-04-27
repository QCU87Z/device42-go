package device42

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

	"testing"
)

func TestClient_GetPasswordByDevice(t *testing.T) {
	client := NewBasicAuthClient("https://10.11.12.239/api/1.0","admin", "adm!nd42")
	expected := "applepie"
	c1, _ := client.GetPasswordByDevice("imac")
	actual := c1.Passwords[0].Password

	if actual != expected {
		t.Errorf("Expected %s got %s",expected, actual)
	}
}

func TestClient_GetPasswordById(t *testing.T) {
	client := NewBasicAuthClient("https://10.11.12.239/api/1.0","admin", "adm!nd42")
	expected := "applepie"
	c1, _ := client.GetPasswordById(1)
	actual := c1.Passwords[0].Password

	if actual != expected {
		t.Errorf("Expected %s got %s",expected, actual)
	}
}

func TestClient_DoRequest(t *testing.T) {


	testServer := httptest.NewServer(

		// NewServer takes a handler.
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Inside the handler we define our canned responses,
			// switching on URL and then http method
			//switch r.URL.Path {
			//case "/reports":
			//	switch r.Method {
			//	case "POST":
			//
			//		// Here we read the body that has been posted to our test server
			//		// and save it to a variable, we can assert against this variable later.
			//		body, _ := ioutil.ReadAll(r.Body)
			//		must(t, json.Unmarshal(body, &posted))
			//		httputil.SendJSON(w, 200, m{})
			//
			//		// Finally provide some defaults, I generally just use a 404
			//	default:
			//		httputil.SendJSON(w, 404, nil)
			//	}
			//default:
			//	httputil.SendJSON(w, 404, nil)
			//}
			user, password, ok := r.BasicAuth()
			if ok != true {
				t.Error("Invalid http basic auth credentials")
			}
			if user != "admin" {
				t.Error("Incorrect http basic username")
			}
			if password != "adm!nd42" {
				t.Error("Incorrect http basic password")
			}

			body, _ := ioutil.ReadAll(r.Body)
			err := json.Unmarshal(body, &x123)
			if err != nil {
				t.Error(err)
			}
			httputil.SendJSON(w, 200, m{})
		}),
	)
	client := NewBasicAuthClient(testServer.URL,"admin", "adm!nd42")
	req, err := http.NewRequest("GET", client.baseURL, nil)
	req.SetBasicAuth(client.Username, client.Password)
	bytes, err := client.DoRequest(req)
	if err != nil {
		t.Error(err)
	}
	var data PasswordAPI
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		t.Error(err)
	}

	defer testServer.Close()
}