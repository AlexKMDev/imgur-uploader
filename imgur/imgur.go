package imgur

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	baseURL = "https://api.imgur.com/3"
)

type (
	Data struct {
		ID          string   `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		DateTime    int      `json:"datetime"`
		Type        string   `json:"type"`
		Animated    bool     `json:"animated"`
		Width       int      `json:"width"`
		Height      int      `json:"height"`
		Size        int      `json:"size"`
		Views       int      `json:"views"`
		Bandwidth   int      `json:"bandwidth"`
		Vote        string   `json:"vote"`
		Favorite    bool     `json:"favorite"`
		Nsfw        string   `json:"nsfw"`
		Section     string   `json:"section"`
		AccountURL  string   `json:"account_url"`
		AccountID   int      `json:"account_id"`
		IsAd        bool     `json:"is_ad"`
		InMostViral bool     `json:"in_most_viral"`
		Tags        []string `json:"tags"`
		AdType      int      `json:"ad_type"`
		AdURL       string   `json:"ad_url"`
		InGallery   bool     `json:"in_gallery"`
		DeleteHash  string   `json:"deletehash"`
		Name        string   `json:"name"`
		Link        string   `json:"link"`
	}

	Error struct {
		Error   string `json:"error"`
		Request string `json:"request"`
		Method  string `json:"method"`
	}

	Answer struct {
		Data    json.RawMessage `json:"data"`
		Success bool            `json:"success"`
		Status  int             `json:"status"`
	}
)

var (
	ClientID = ""
)

func Upload(filename string) (*Data, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/image", baseURL), body)

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", ClientID))
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	var answer Answer
	if err = json.Unmarshal(respBody, &answer); err != nil {
		return nil, err
	}

	var data Data
	var error Error

	if err = json.Unmarshal(answer.Data, &data); err != nil {
		if err = json.Unmarshal(answer.Data, &error); err != nil {
			return nil, err
		}

		return nil, errors.New(error.Error)
	}

	return &data, nil
}

func Delete(hash string) (*bool, error) {
	body := &bytes.Buffer{}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/image/%s", baseURL, hash), body)
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", ClientID))

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	var answer Answer
	if err = json.Unmarshal(respBody, &answer); err != nil {
		return nil, err
	}

	var data bool
	var error Error

	if err = json.Unmarshal(answer.Data, &data); err != nil {
		if err = json.Unmarshal(answer.Data, &error); err != nil {
			return nil, err
		}

		return nil, errors.New(error.Error)
	}

	return &data, nil
}
