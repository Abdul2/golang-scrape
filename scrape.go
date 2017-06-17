package main

import (

	"fmt"
	"golang.org/x/net/html"
	"strings"
	"log"
	"os"
	"io/ioutil"
	"net/http"
	//"regexp"
	"regexp"
)

func main() {


	url := os.Getenv("landingpage")

	r, _ := regexp.Compile(`\/cas\/\d\d\d\d\d\d\d\d\d\d\.html`) // /cas/6180039890.html]


	values := htmlvalueextractor(url,"href","a")


	fmt.Printf("%v",values)

	for _, v:= range values{

		if r.MatchString(string(v)) {fmt.Printf("%s\n",v)}

	}



}

//returns a slice of data caller is looking for. for a given url, node type and key it retruns a slice of values
//associated with the keys
// for example <a href=2 > and <a href=3> where a=node and href=key it will return slice of 2,3
//the caller is responsible for processing the returned data
func Errorhandling(err error){

	if err != nil {
		log.Fatal(err)
	}

}



func htmlvalueextractor (url string, key string, node string) []string{


b := make([]string, 1)

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


	//variable as a function
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == node {
			for _, a := range n.Attr {
				if a.Key == key {
					fmt.Println(a.Val)
					st := fmt.Sprintf("%s",a.Val)
					b = append(b,st)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

return b

}


