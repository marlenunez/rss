/*
From post: "Golang : Decode XML data from RSS feed"
https://www.socketloop.com/tutorials/golang-decode-xml-data-from-rss-feed

Clean p tag: jaytaylor/html2text
https://github.com/jaytaylor/html2text
*/
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/jaytaylor/html2text"
	"io/ioutil"
	"net/http"
	"os"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

func main() {

	// response, err := http.Get("http://www.thestar.com.my/RSS/Metro/Community/")
	response, err := http.Get("http://es-us.deportes.yahoo.com/b%C3%A1squet/?format=rss")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	XMLdata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rss := new(Rss)

	buffer := bytes.NewBuffer(XMLdata)

	decoded := xml.NewDecoder(buffer)

	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("---------[rss]----------\n")
	fmt.Printf("Title : %s\n", rss.Channel.Title)
	fmt.Printf("Description : %s\n", rss.Channel.Description)
	fmt.Printf("Link : %s\n", rss.Channel.Link)

	total := len(rss.Channel.Items)

	fmt.Printf("Total items : %v\n\n", total)

	for i := 0; i < total; i++ {
		fmt.Printf("---------[%d]----------\n\n", i+1)
		fmt.Printf("%s\n\n", rss.Channel.Items[i].Title)
		fmt.Printf("%s\n\n", strip(rss.Channel.Items[i].Description))
		fmt.Printf("%s\n\n", rss.Channel.Items[i].Link)
	}

}

func strip(str string) string {
	descri, err := html2text.FromString(str)
	if err != nil {
		panic(err)
	}
	return descri
}
