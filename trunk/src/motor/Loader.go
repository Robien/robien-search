package main

import (
    "bufio"
    "os"
	"io"
	"strconv"
	"fmt"
)

func loadElements() {
    // open input file
    fi, err := os.Open("data/elements.dat")
    if err != nil { return }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // make a read buffer
    r := bufio.NewReader(fi)

    // make a buffer to keep chunks that are read

    for {
	    buf := make([]byte, 1)
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		buf = make([]byte, buf[0])
		n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var name string  = string(buf)
		buf = make([]byte, 4)
		n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var id int = int(buf[0]) + int(buf[1]<<8) +  int(buf[2]<<16) +  int(buf[3]<<24)
		if motor.ElementFactory(name).Id != id{
			fmt.Println("sp'a cool")
		}		
    }

 
}

func loadConcepts() {
    // open input file
    fi, err := os.Open("data/concepts.dat")
    if err != nil { return }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // make a read buffer
    r := bufio.NewReader(fi)

    // make a buffer to keep chunks that are read

    for {
	    buf := make([]byte, 1)
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		buf = make([]byte, buf[0])
		n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var name string  = string(buf)
		 buf = make([]byte, 1)
        // read a chunk
        n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		buf = make([]byte, buf[0])
		n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var question string  = string(buf)
		buf = make([]byte, 4)
		n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var id int = int(buf[0]) + int(buf[1]<<8) +  int(buf[2]<<16) +  int(buf[3]<<24)
			
		if motor.ConceptFactory(name, question).Id != id{
			fmt.Println("sp'a cool")
		}	
    }

 
}

func loadLinks() {
    // open input file
    fi, err := os.Open("data/links.dat")
    if err != nil { return }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // make a read buffer
    r := bufio.NewReader(fi)


    for {

		buf := make([]byte, 4)
		n, err := r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var element int32 = int32(buf[0]) + int32(buf[1]<<8) +  int32(buf[2]<<16) +  int32(buf[3]<<24)
		buf = make([]byte, 4)
		n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var concept int32 = int32(buf[0]) + int32(buf[1]<<8) +  int32(buf[2]<<16) +  int32(buf[3]<<24)
		 buf = make([]byte, 1)
        // read a chunk
        n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		buf = make([]byte, buf[0])
		n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var weigthS string  = string(buf)
		 buf = make([]byte, 1)
        // read a chunk
        n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		buf = make([]byte, buf[0])
		n, err = r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
		var valueS string  = string(buf)

		value, _ := strconv.ParseFloat(valueS, 64)
		weigth, _ := strconv.ParseFloat(weigthS, 64)

		motor.LinkFactoryWithWeight(motor.elements[element], motor.concepts[concept], value, weigth)	
    }

 
}