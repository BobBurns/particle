package particle

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
	"fmt"
	"bytes"
)

type Result struct {
	OK bool `json:"ok"`
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}


// publishes a particle.io event stream and returns status
func (e *Event) Publish (token string) (*Result, error) {
	out := new(Result)

	var client *http.Client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{tr, nil, nil, 0 * time.Second}

	u, err := url.Parse(URL)
	if err != nil {
		return out, err
	}

	q := url.Values{}
	q.Set("name", e.Name)
	q.Set("data", e.Data.Data)
	if e.Data.Private {
		q.Set("private", "true")
	}


	q.Set("ttl", string(e.Data.TTL))
	q.Set("access_token", token)

	resp, err := client.PostForm(u.String(), q)
	if err != nil {
		fmt.Println("error: client.Get")
		return out, err
	}

	reader := bufio.NewReader(resp.Body)

	// check for :ok as first event on stream

	body := make([]byte, 1024)
	_, err = reader.Read(body)

	if err != nil {
		fmt.Println("error: reader.Read")
		return out, err
	}
	resp.Body.Close()

	body = bytes.Trim(body, "\x00")

	err = json.Unmarshal(body, &out)
	if err != nil {
	  fmt.Println("error: json.Unmarshal")
	  return out, err
	}

	return out, nil
}
