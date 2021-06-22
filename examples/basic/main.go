package main

import (
	"fmt"

	"github.com/iesen/go-notion"
)

func main() {
	client := notion.NewClient()
	fmt.Println("Hey client is", client)
}
