package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var globalDB *mgo.Database
var account = "arnold"
var in []chan string
var out []chan result
var maxUser = 100
var maxThread = 10

type currency struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Amount  float64       `bson:"amount"`
	Account string        `bson:"account"`
	Code    string        `bson:"code"`
	Version int           `bson:version`
}

type result struct {
	Account string
	Result  float64
}

func Random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min+1) + min
}

func pay(w http.ResponseWriter, r *http.Request) {
	number := Random(1, maxUser)
	channelNumber := number % maxThread
	account = "user" + strconv.Itoa(number)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		in[channelNumber] <- account
		for {
			select {
			case result := <-out[channelNumber]:
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
	in = make([]chan string, maxThread)
	out = make([]chan result, maxThread)

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

	for i := range in {
		in[i] = make(chan string)
		out[i] = make(chan result)
	}

	// create 100 user
	for i := 0; i < maxUser; i++ {
		account = "user" + strconv.Itoa(i+1)
		user := currency{Account: account, Amount: 1000.00, Code: "USD"}
		if err := globalDB.C("bank").Insert(&user); err != nil {
			panic("insert error")
		}
	}
	for i := range in {
		go func(in *chan string, i int) {
			for {
				select {
				case data := <-*in:
					entry := currency{}
				LOOP:
					err := globalDB.C("bank").Find(bson.M{"account": data}).One(&entry)
					if err != nil {
						panic(err)
					}

					entry.Amount = entry.Amount + 50.00

					err = globalDB.C("bank").Update(bson.M{
						"_id":     entry.ID,
						"version": entry.Version,
					}, bson.M{"$set": map[string]interface{}{
						"amount":  entry.Amount,
						"version": (entry.Version + 1),
					}})

					if err != nil {
						goto LOOP
					}

					out[i] <- result{
						Account: account,
						Result:  entry.Amount,
					}

				}
			}

		}(&in[i], i)
	}

	log.Println("Listen server on " + port + " port")
	http.HandleFunc("/", pay)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
