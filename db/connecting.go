package mgo


import (
	"errors"
	"time"
	//"context"
	"go.mongodb.org/mongo-driver/mongo"
)

const mgoNum int = 3
const mgoURL string = "mongodb://root:weel99699@192.168.43.113:27017/?connect=direct&compressors=disabled&gssapiServiceName=mongodb"
//var Clients [mgoNum]*mongo.Client
var clients chan *mongo.Client

func init(){
	var c *mongo.Client
	clients = make(chan *mongo.Client, mgoNum)
	for i:=0 ; i<mgoNum ; i++ {
		c = connecting(mgoURL)
		if err := ping(c); err != nil {
			panic(err)
		} else {
			clients <- c
		}
	}
}

func GetClient() (*mongo.Client, error) {
	select {
		case c := <- clients :
			return c, nil
		case <- time.After(time.Second * 10) : 
			return nil, errors.New("Client waiting timeout.")
	}
}

func ReturnClient(c *mongo.Client) error {
	select {
		case clients <- c:
			return nil
		case <- time.After(time.Second * 10) :
			return errors.New("Client return error.")
	}
}









