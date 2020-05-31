package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// APIResponse contains the data returned in the api
type Response struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// NewResponse creates a new api response
func NewResponse(r io.Reader) Response {
	var ar Response
	err := json.NewDecoder(r).Decode(&ar)
	if err != nil {
		log.Fatal("error parsing data", err)
	}
	return ar
}

// Comic returns a Comic from an API response
func (ar Response) Comic() Comic {
	date := fmt.Sprintf("%s-%s-%s", ar.Year, ar.Month, ar.Day)
	return Comic{
		Title:       ar.Title,
		Number:      ar.Num,
		Date:        date,
		Description: ar.Alt,
		Image:       ar.Img,
	}
}
