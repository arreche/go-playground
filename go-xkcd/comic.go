package main

import (
	"encoding/json"
	"fmt"
)

// Comic contains the data for the application
type Comic struct {
	Title       string `json:"title"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// Text returns a text representation
func (c Comic) Text() string {
	return fmt.Sprintf("Title: %s\nNo: %d\nDate: %s\nDesc: %s\nImage: %s\n", c.Title, c.Number, c.Date, c.Description, c.Image)
}

// JSON returns a json representation
func (c Comic) JSON() []byte {
	bytes, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return bytes
}
