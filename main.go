package main

import (
	"fmt"
	"log"
	"os"
)

// var endpoints []Endpoint
var endpoints map[string]string

func init() {
	l := log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	l.Println("Initialize API Client...")
	l.Println("Check status API...")
}

func main() {
	endpoints = getEndpoints()
	fmt.Println(endpoints["characters"])
}
