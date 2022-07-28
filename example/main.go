package main

import (
	"fmt"

	"github.com/sofyan48/go-viro"
)

func main() {
	client := viro.NewViro("testing", "API_KEY")
	res, err := client.SMS().Single("62812479xxxx", "text to send").Send()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))
}
