package list

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/Team-AppScripts/golang/Helper/WebRequest"
)

func getPage(url string) []string {

	var body io.Reader = WebRequest.Request(url)

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	var openList []string = nil

	doc.Find(".sc-jbKcbu > a").Each(func(idx int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		openList = append(openList, href)
	})

	return openList
}

func chooseSonarLint(url string, r *http.Request) map[string]map[string]string {
	// choose Request target url.
	var responseAry []string = getPage(r.URL.Query().Get(url))
	var temp io.Reader = nil
	responseMaps := make(map[string]map[string]string)
	tmpMaps := make(map[string]string)

	for idx := range responseAry {
		temp = WebRequest.Request(responseAry[idx])
		doc, _ := goquery.NewDocumentFromReader(temp)
		if doc.Find("img[src=\"/logos/SonarLint-black.svg\"]") != nil {
			tmpMaps["action"] = doc.Find("main .sc-hMqMXs").Text()
			tmpMaps["Non_compliant_code"] = doc.Find("main .sc-cvbbAY pre").Eq(0).Text()
			tmpMaps["compliant_code"] = doc.Find("main .sc-hMqMXs").Eq(1).Text()
		}

		responseMaps[url] = tmpMaps
		//response = append(response, doc.)
	}

	return responseMaps
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	var responseAry map[string]map[string]string = chooseSonarLint(r.URL.Query().Get("ref"), r)
	fmt.Fprint(w, responseAry)
}
