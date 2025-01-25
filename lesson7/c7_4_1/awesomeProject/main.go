package main

import "os"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func main() {
	p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
	p1.save()
}
