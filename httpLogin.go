package httpLogin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(url, username, password string) bool {
	fmt.Println("accessing: ", url)
	//If you need proxy uncomment the following
	//tr := &http.Transport{
	//	Proxy: nil,
	//}
	//client := &http.Client{Transport: tr}
	client := &http.Client{}

	var jsonStr = []byte(`{"username":"` + username + `", "password":"` + password + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

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
