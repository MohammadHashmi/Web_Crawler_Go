package main

// Imports
import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// Main Function
func main() {
	bfs_on_links()
}

func bfs_on_links(){
	//Defining data structures needed for traversal
	visited := []string{}
	results := []string{}
	queue := []string{}

	neighbors := findNeighbors("https://httpbin.org/links/10/0")

	//visited = append(visited, neighbors...)
	//results = append(results, neighbors...)
	queue = append(queue, neighbors...)

	fmt.Println(queue)
	x := queue[0]
	visited = append(visited, x)
	queue = queue[1:]

	for queue:
		curr := queue[0]
		queue = queue[1:]
		results = append(results, curr)

		y:= queue[0]
		for y {
			if !slice.Contains(visited, y) {
				append(visited, y)
				append(queue, y)
			}
			
		}



	fmt.Println(queue, x)
}

func findNeighbors(link string) []string {
	neighbors := []string{}

	// Testing parsing on httpbin
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	// Parse the html of the response body
	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}

		// Iterate through every element in the html
	for n := range doc.Descendants() {
		// If the type is "a" and it's a link (href) print the value of the link
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, v := range n.Attr {
				//TODO: Absolute links?
				if v.Key == "href" {
					neighbors = append(neighbors, v.Val)
					break
				}
			}
		}
	}
	return neighbors
}