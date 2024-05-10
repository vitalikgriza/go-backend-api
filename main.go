package main

import "fmt"

func main() {
	fmt.Println("Hey")
	server := NewServer(":4000")
	server.Run()
}
