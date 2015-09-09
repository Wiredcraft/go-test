package httpLogin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(url, username, password string) bool {
	fmt.Println("accessing: ", url)
	tr := &http.Transport{
		Proxy: nil,
	}
	var jsonStr = []byte(`{"username":"` + username + `", "password":"` + password + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))
	return res.StatusCode == http.StatusOK
}
