package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/gcalvoCR/go-learn/cmd/002-binary-encoding/demo"
	"github.com/gcalvoCR/go-learn/cmd/002-binary-encoding/model"
	proto "google.golang.org/protobuf/proto"
)

func main() {
	// client side
	msgToServer := &demo.ClientReq{ID: 1, Type: model.ReqAvg}

	buf := &bytes.Buffer{}

	xml.NewEncoder(buf).Encode(msgToServer)
	fmt.Printf("buf size xml %v\n", buf.Len())

	gob.NewEncoder(buf).Encode(msgToServer)
	fmt.Printf("buf size gob %v\n", buf.Len())

	json.NewEncoder(buf).Encode(msgToServer)
	fmt.Printf("buf size json %v\n", buf.Len())

	// with json, gob, xml, you need to create a new encoder, with proto, there's no need for that
	buf2, err := proto.Marshal(msgToServer)
	if err != nil {
		log.Println("Unable to encode protobuf message")
	}

	fmt.Printf("buf size  proto %v\n", len(buf2))

	// server side
	msgFromClient := &demo.ClientReq{}
	xml.NewDecoder(buf).Decode(msgFromClient)
	fmt.Printf("%v\n", msgFromClient)
	json.NewDecoder(buf).Decode(msgFromClient)
	fmt.Printf("%v\n", msgFromClient)
	gob.NewDecoder(buf).Decode(msgFromClient)
	fmt.Printf("%v\n", msgFromClient)
	proto.Unmarshal(buf2, msgFromClient)
	fmt.Printf("%v\n", msgFromClient)

	sr1 := demo.SearchResp{
		Id:     5,
		Ans:    "this is the answer",
		Source: demo.From_Cache,
	}

	fmt.Println(sr1)

}
