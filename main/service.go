package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// func homepageHandler(w http.ResponseWriter, r *http.Request) {
// 	_, err := fmt.Fprintf(w, "hello world")
// 	checkError(err)
// }

type thumbnailRequest struct {
	Url string `json: "url"`
}

type screenshotAPIRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width"`
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		checkError(err)
		return
	}
	apiRequest := screenshotAPIRequest{
		Token:          "SQQR4NBLUJ4AA5RMQMJMI3SMHEWGJS88",
		Url:            decoded.Url,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
	}

	jsonString, err := json.Marshal(apiRequest)
	checkError(err)

	req, err := http.NewRequest("POST", "https://screenshotapi.net/api/v1/screenshot", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	defer response.Body.Close()

	type screenshotAPIResponse struct {
		Screenshot string `json"screenshot"`
	}

	var apiResponse screenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(err)

	// Pass back the screenshot URL to the frontend.
	_, err = fmt.Fprintf(w, `{ "screenshot": "%s" }`, apiResponse.Screenshot)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	http.HandleFunc("/api/thumbnail", thumbnailHandler)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
