package main

import (
	"errors"
	"fmt"
	//"io"
	"io/ioutil"
	// "os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	if p.Title == "" {
		return errors.New("Page must have a title")
	}
	fmt.Printf("Save page %s.txt\n", p.Title)
	return ioutil.WriteFile(RESSOURCE_PATH+"/data/"+p.Title+".txt", p.Body, 0600)
}

func load(title string) (*Page, error) {
	fmt.Printf("Load page %s.txt\n", title)
	content, err := ioutil.ReadFile(RESSOURCE_PATH+"/data/"+title + ".txt")
	return &Page{title, content}, err
}

func (p *Page) String() string {
	return fmt.Sprintf("Page: Title[%s] Content[%s]", p.Title, p.Body[:50])
}
