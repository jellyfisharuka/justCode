package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*import ( USING XML. simple example
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
*/

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int64  `json:"age"`
	Address  UserAddress
}

type UserAddress struct {
	City   City   `json:"city"`
	Region Region `json:"region"`
	Street string `json:"street"`
}
type City struct {
	Name    string `json:"cityName"`
	ZipCode string `json:"zipCode"`
}
type Region struct {
	Name string `json:"regionName"`
	Code string `json:"regionCode"`
}

func main() {

	jsonBytes, err := os.Open("user_data.json")
	if err != nil {
		panic(err)
	}
	defer jsonBytes.Close()
	var user1 User
	decodeUser := json.NewDecoder(jsonBytes)
	err = decodeUser.Decode(&user1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded user data from json %+v\n", user1)

}

///тут можно и использовать Marshal, UnMarshal. Но Decoder полезно для чтения больших JSON-файлов и позволяет читать данные покадрово и обрабатывать их по мере их поступления
