package WebRequest

import (
	"io"
	"log"
	"net/http"
)

func Request(url string) io.ReadCloser {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	return res.Body
}
