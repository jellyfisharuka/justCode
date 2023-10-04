package main

import (
	"encoding/xml"
	"fmt"
)

type Author struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`
}
type Book struct {
	XMLName xml.Name `xml:"name"`
	Title   string   `xml:"title"`
	Author  Author   `xml:"author"`
}
type Jobs struct {
	XMLName xml.Name
}

func main() {
	author := Author{Name: "Aruzhan"}
	book := Book{
		Title:  "Aruzhansabaq",
		Author: author,
	}
	xmlInfo, err := xml.MarshalIndent(book, "", "   ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(string(xmlInfo))
}
