package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "time"
)

var client http.Client

func main() {
	n := flag.Int("n", 0, "Comic number to fetch (default latest)")
	o := flag.String("o", "text", "Print output in format: text/json")
	s := flag.Bool("s", false, "Save image to current directory")
	t := flag.Int("t", 30, "Client timeout in seconds")
	flag.Parse()

	if *n < 0 {
		log.Fatal("invalid comic number")
	}

	if *o != "text" && *o != "json" {
		log.Fatal("unsupported output format")
	}

	if *t < 0 {
		log.Fatal("invalid timeout")
	}

    client = http.Client{Timeout: time.Duration(*t) * time.Second}

    comic, err := fetchComic(*n)
    if err != nil {
        log.Fatal("error fetching comic", err)
    }

	printOutput(comic, *o)

	if *s {
        saveImage(comic)
	}
}

func fetchComic(n int) (Comic, error) {
    url := "http://xkcd.com/info.0.json"
    if n > 0 {
        url = fmt.Sprintf("http://xkcd.com/%d/info.0.json", n)
    }

    res, err := client.Get(url)
    if err != nil {
        return Comic{}, err
    }
    defer res.Body.Close()

    if res.StatusCode >= http.StatusBadRequest {
        return Comic{}, fmt.Errorf("error fetching data with status code %d", res.StatusCode)
    }

    return NewResponse(res.Body).Comic(), nil
}

func printOutput(c Comic, f string) {
    var output string
    if f == "json" {
        output = string(c.JSON())
    } else {
        output = c.Text()
    }
    fmt.Printf(output)
}

func saveImage(c Comic) error {
    url := c.Image
    res, err := client.Get(url)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    bytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err
    }

    filename := url[strings.LastIndex(url, "/"):]
    err = ioutil.WriteFile("./"+filename, bytes, 0644)
    if err != nil {
        return err
    }

    return nil
}
