package WebRequest

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Request(url string) io.ReadCloser {
	// Request the HTML page.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	res, _ := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	} else {
		fmt.Println(res.Request.Body)
	}

	return res.Body
}

func TestCode(url string) {
	var testReader io.Reader = Request(url)
	buf := new(bytes.Buffer)
	buf.ReadFrom(testReader)
	fmt.Print(buf.String())
}
