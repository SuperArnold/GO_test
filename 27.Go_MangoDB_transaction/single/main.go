package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var globalDB *mgo.Database
var account = "arnold"
var in chan string
var out chan result

type currency struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Amount  float64       `bson:"amount"`
	Account string        `bson:"account"`
	Code    string        `bson:"code"`
}

type result struct {
	Account string
	Result  float64
}

func pay(w http.ResponseWriter, r *http.Request) {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		in <- account
		for {
			select {
			case result := <-out:
				fmt.Printf("%+v\n", result)
				wg.Done()
				return
			}
		}

	}(&wg)

	wg.Wait()
	io.WriteString(w, "ok")

}

func main() {
	in = make(chan string)
	out = make(chan result)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	session, _ := mgo.Dial("localhost:27017")
	globalDB = session.DB("queue")
	globalDB.C("bank").DropCollection()

	user := currency{Account: account, Amount: 1000.00, Code: "USD"}
	err := globalDB.C("bank").Insert(&user)

	if err != nil {
		panic("insert error")
	}

	go func(in *chan string) {
		for {
			select {
			case data := <-*in:
				entry := currency{}

				err := globalDB.C("bank").Find(bson.M{"account": data}).One(&entry)
				if err != nil {
					panic(err)
				}

				entry.Amount = entry.Amount + 50.00

				err = globalDB.C("bank").UpdateId(entry.ID, &entry)

				if err != nil {
					panic("update error")
				}

				out <- result{
					Account: account,
					Result:  entry.Amount,
				}

			}
		}

	}(&in)

	log.Println("Listen server on " + port + " port")
	http.HandleFunc("/", pay)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
