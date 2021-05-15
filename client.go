package notion

import "fmt"

type Client struct {
}

func NewClient() Client {
	fmt.Println("Created")
	return Client{}
}
