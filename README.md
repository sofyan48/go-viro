# GOLANG API FOR VIRO SMS 
Unofficial VIRO SMS API

## How To Use

Install package via go module 
``` bash
go get github.com/sofyan48/go-viro
```

Example send SMS
``` golang
package main

import (
	"fmt"

	"github.com/sofyan48/go-viro"
	"github.com/sofyan48/go-viro/entity"
)

func main() {
	client := viro.NewViro("testing", "API_KEY")
	single(client)
	advance(client)
}

func single(client *viro.Viro) {
	res, err := client.SMS().Single("62812479xxxx", "text to send").Send()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func advance(client *viro.Viro) {
	post1 := entity.AdvancePayload{
		To:   "08124793xxxx",
		Text: "Text To Send",
	}
	postsData := []entity.AdvancePayload{}
	postsData = append(postsData, post1)
	postsData = append(postsData, post1)
	res, err := client.SMS().Advance(postsData).Send()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

```