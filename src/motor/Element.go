package main

import "fmt"

type Element struct {
	Name string
	Id int
	links []*Link
}


func (e *Element) Print()  {
	fmt.Println(e.Name)
}

func (e *Element) PrintDebug()  {
	fmt.Println(e.Name + " (" , e.Id , ")")
}

func (e *Element) AddALink(l *Link)  {
	e.links = append(e.links, l)
}
