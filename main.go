package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/egnimos/golang-protobuf/src/simplepb"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	writeToFile("simple.bin", sm)

	sm2 := simplepb.SimpleMessage{}
	readToFile("simple.bin", &sm2)
	fmt.Println(sm2)
}

/***********WRITE AND READ TO JSON************/
// func toJSON() string {

// }

/********WRITE AND READ TO THE DISK*********/
func readToFile(fName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	//unmarshal
	if err2 := proto.Unmarshal(in, pb); err2 != nil {
		log.Fatalln("couldn't put the files into protocol buffers struct", err)
		return err2
	}

	return nil
}

func writeToFile(fName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("can't serialise to bytes", err)
		return err
	}

	//write to the file
	if err := ioutil.WriteFile(fName, out, 0644); err != nil {
		log.Fatalln("can't write to the file", err)
		return err
	}

	fmt.Println("wrtitten to the file")

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm.GetId())

	return &sm
}
