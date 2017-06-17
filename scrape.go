
package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"regexp"
	"os"
	// "io/ioutil" debug use

)




func f(n *html.Node) {


	r, _ := regexp.Compile(`\/cas\/\d\d\d\d\d\d\d\d\d\d\.html`) // /cas/6180039890.html]

	if n.Type == html.ElementNode && n.Data == "a" {


		for _, t := range n.Attr{


			s := fmt.Sprint(t)

			//fmt.Println(s)

			if r.MatchString(s){

				s1 := r.FindAllString(s, -1)

				fmt.Printf("%s\n",s1)


			}
		}



	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c)
	}
}


func main() {


	//point at landing page

	url := os.Getenv("landingpage")


	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//debug use
	//page, _  := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s",page)


	doc, err := html.Parse(resp.Body)




	if err != nil {

		panic(err)
	}
	//var f func(*html.Node)

	f(doc)



}

