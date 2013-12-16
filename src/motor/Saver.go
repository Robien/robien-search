package main

import (
    "bufio"
    "os"
	"strconv"
	"fmt"
)

func saveElements(elems []*Element) {
  
  fmt.Print("sauv Elems ... ")
    // open output file
	os.Remove("data/elements.dat")
    fo, err := os.Create("data/elements.dat")
    if err != nil { panic(err) }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // make a write buffer
    w := bufio.NewWriter(fo)

    // make a buffer to keep chunks that are read

    for _,e := range elems{
	size := make([]byte, 1)
	size[0] = byte(len(e.Name))
    buf := make([]byte, size[0])
	
	 copy(buf, e.Name)
	 if _, err := w.Write(size); err != nil {
            panic(err)
        }
        if _, err := w.Write(buf); err != nil {
            panic(err)
        }
		buf = make([]byte, 4)
		 buf[0], buf[1], buf[2], buf[3] = byte(e.Id & 0xff), byte(e.Id>>8 & 0xff), byte(e.Id>>16 & 0xff), byte(e.Id>>24 & 0xff)
		 if _, err := w.Write(buf); err != nil {
            panic(err)
        }
    }

    if err = w.Flush(); err != nil { panic(err) }
	fmt.Println("finish")
}

func saveConcepts(concepts []*Concept) {
  fmt.Print("sauv concepts ... ")
    // open output file
		os.Remove("data/concepts.dat")
    fo, err := os.Create("data/concepts.dat")
    if err != nil { panic(err) }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // make a write buffer
    w := bufio.NewWriter(fo)

    // make a buffer to keep chunks that are read

    for _,e := range concepts{
	
   	size := make([]byte, 1)
	size[0] = byte(len(e.Name))
    buf := make([]byte, size[0])
	
	 copy(buf, e.Name)
	 if _, err := w.Write(size); err != nil {
            panic(err)
        }
        if _, err := w.Write(buf); err != nil {
            panic(err)
        }

	size[0] = byte(len(e.Question))
    buf = make([]byte, size[0])
	
	 copy(buf, e.Question)
	 if _, err := w.Write(size); err != nil {
            panic(err)
        }
        if _, err := w.Write(buf); err != nil {
            panic(err)
        }
		buf = make([]byte, 4)
		 buf[0], buf[1], buf[2], buf[3] = byte(e.Id & 0xff), byte(e.Id>>8 & 0xff), byte(e.Id>>16 & 0xff), byte(e.Id>>24 & 0xff)
		 if _, err := w.Write(buf); err != nil {
            panic(err)
        }
    }

    if err = w.Flush(); err != nil { panic(err) }
		fmt.Println("finish")
}

func saveLinks(links []*Link) {
  fmt.Print("sauv links ... ")
    // open output file
	os.Remove("data/links.dat")
    fo, err := os.Create("data/links.dat")
    if err != nil { panic(err) }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // make a write buffer
    w := bufio.NewWriter(fo)

    // make a buffer to keep chunks that are read

    for _,l := range links{
	
		buf := make([]byte, 4)
		 buf[0], buf[1], buf[2], buf[3] = byte(l.element & 0xff), byte(l.element>>8 & 0xff), byte(l.element>>16 & 0xff), byte(l.element>>24 & 0xff)
		 if _, err := w.Write(buf); err != nil {
            panic(err)
        }
		buf = make([]byte, 4)
		 buf[0], buf[1], buf[2], buf[3] = byte(l.concept & 0xff), byte(l.concept>>8 & 0xff), byte(l.concept>>16 & 0xff), byte(l.concept>>24 & 0xff)
		 if _, err := w.Write(buf); err != nil {
            panic(err)
        }
		
		s := strconv.FormatFloat(l.weight, 'f', 5, 64)
	
	
	size := make([]byte, 1)
	size[0] = byte(len(s))
	buf = make([]byte, len(s))	
		copy(buf, s)

	 if _, err := w.Write(size); err != nil {
            panic(err)
        }
        if _, err := w.Write(buf); err != nil {
            panic(err)
        }
		
		s = strconv.FormatFloat(l.value, 'f', 5, 64)
			size = make([]byte, 1)
	size[0] = byte(len(s))
	buf = make([]byte, len(s))	
		copy(buf, s)

	 if _, err := w.Write(size); err != nil {
            panic(err)
        }
        if _, err := w.Write(buf); err != nil {
            panic(err)
        }
    }

    if err = w.Flush(); err != nil { panic(err) }
		fmt.Println("finish")
}
