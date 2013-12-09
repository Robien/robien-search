package main

import "fmt"

type Link struct {
element int
concept int
weight float64
value float64

}

func (l *Link) Create(c *Concept, e *Element){
	e.AddALink(l)
	c.AddALink(l)
	l.element = e.Id
	l.concept = c.Id
}

func (l *Link) Print()  {
	fmt.Println("pomme")
}

