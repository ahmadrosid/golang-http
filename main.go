package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		enc := json.NewEncoder(writer)
		err := enc.Encode(map[string]interface{}{
			"title": "Some",
			"id":    123,
		})
		if err != nil {
			msg := fmt.Sprintf("err: %s", err.Error())
			_, _ = writer.Write([]byte(msg))
			return
		}
	})
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("error starting the server: %q", err)
	}
}

func getRequest() {
	url := "https://poetrydb.org/title/spring/title"
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("err %q\n", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("err %q\n", err)
	}

	text := string(body[:])
	fmt.Println(text)
}

func postRequest(url string) error {
	body, err := json.Marshal(map[string]interface{}{
		"title": "Some",
		"id":    123,
	})
	if err != nil {
		return err
	}

	payloadBodyRaw := bytes.NewReader(body)
	_, err = http.Post(
		url,
		"application/json",
		payloadBodyRaw,
	)

	if err != nil {
		return err
	}
	return nil
}
