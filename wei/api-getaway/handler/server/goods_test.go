package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEs(t *testing.T) {
	url := "http://127.0.0.1:8082/es?keyword=十&page=1&size=2"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Njc3MDA1ODQsImlhdCI6MTc2NzY5Njk4NCwidXNlcklkIjoiMSJ9.mq9RBey_0ROQaQetLlecs_vYAEU4Nyh2gtmqP6IsrWQ")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
