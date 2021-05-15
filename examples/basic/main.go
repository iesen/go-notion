package main

import (
	"fmt"

	"github.com/iesen/go-notion"
)

func main() {
	_ = notion.NewClient()
	fmt.Println("Hey")
}
