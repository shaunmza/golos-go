package main

import (
	"fmt"

	"github.com/shaunmza/golos-go"
	"github.com/shaunmza/golos-go/operations"
)

var(
  weight int16 = 0
  pkey = "<Privat Posting Key>"
  user = "<UserName>"
  author = "<AuthorName>"
  link = "<Permlink>"
)

func main() {
	cls, _ := golos.NewClient("https://hive.roelandp.nl")
	defer cls.Close()

	cls.SetKeys(&golos.Keys{PKey: []string{pkey}})

	var trx []operations.Operation

	tx := &operations.VoteOperation{
		Voter:    user,
		Author:   author,
		Permlink: link,
		Weight:   weight,
	}
	trx = append(trx, tx)

	resp, err := cls.SendTrx(user, trx)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Answer : ", resp)
	}
}