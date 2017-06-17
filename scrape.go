package main

import (

	"fmt"
	"golang.org/x/net/html"
	"strings"
	"log"
	"os"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {


	url := os.Getenv("landingpage")

	r, _ := regexp.Compile(`\/cas\/\d\d\d\d\d\d\d\d\d\d\.html`) // /cas/6180039890.html]


	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()


	page, err  := ioutil.ReadAll(resp.Body)

	Errorhandling(err)

	s := fmt.Sprintf("%s",page)


	doc, err := html.Parse(strings.NewReader(s))


	Errorhandling(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && r.MatchString(a.Val){
					fmt.Println(a.Val)

					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

//handles error
func Errorhandling(err error){

	if err != nil {
		log.Fatal(err)
	}

}