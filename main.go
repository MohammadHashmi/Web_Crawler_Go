package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	res, err := http.Get("https://httpbin.org/links/10/0")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	//..body, err := io.ReadAll(res.Body)

	//if err != nil {
	//	log.Fatal(err)
	//}

	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, v := range n.Attr {
				if v.Key == "href" {
					fmt.Println(v.Val)
					break
				}
			}
		}
	}
	fmt.Println("testing")
}
