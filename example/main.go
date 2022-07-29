package main

import (
	"fmt"

	"github.com/sofyan48/go-viro"
	"github.com/sofyan48/go-viro/entity"
)

func main() {
	client := viro.NewViro("testing", "API_KEY")
	client.SMS().GetLogs(&entity.LogsParameters{
		To:            "89999",
		BulkID:        []string{"testing", "testing2"},
		MessageID:     []string{"msgID", "msgID2"},
		GeneralStatus: "",
		Limit:         0,
		MCC:           "",
		MNC:           "",
	})
	// single(client)
	// advance(client)
	// multi(client)
}

func single(client *viro.Viro) {
	res, err := client.SMS().Single("62812479xxxx", "text to send").Send()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func multi(client *viro.Viro) {
	post1 := entity.MultiPayload{
		To: []string{
			"081247930699",
			"081247930699",
		},
		Text: "Text To Send",
	}
	postsData := []entity.MultiPayload{}
	postsData = append(postsData, post1)
	postsData = append(postsData, post1)
	res, err := client.SMS().Multi(postsData).Send()
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
