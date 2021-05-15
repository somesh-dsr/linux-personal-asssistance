package main

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"time"
)

var sample chan struct{}
type sample_struct struct {
	Data string
}
type testiterface interface {
	display()
}
type sample2 struct {
	sample_struct
	sampledata string
}

func (s sample2) display() {
	//panic("implement me")
}

type sample3 struct {
	sample_struct
}

func (s sample3) display() {
	//panic("implement me")
}
func (s sample3) hello() {
//	panic("implement me")
}

func main(){
	cmd := exec.Command("espeak", "Hai DSR")
	fmt.Println("Executing espeak")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Executing espeak done")
	sample = make(chan struct{})
	go test()
	go test()
	//for {
		select {
			case <-sample:
				fmt.Println("get")
		//default:
			//fmt.Println("Got default")
		}
	//}
	var ti testiterface
	ti = &sample2{

		sample_struct: sample_struct{

			Data: "data sample2",
		},
	}

	res := reflect.ValueOf(ti).Elem()
	fmt.Println(res.FieldByName("Data"))

	fmt.Printf("%v\n",res)
	ti.display()
	ti = new(sample3)
	ti.display()
	//var check chan struct{}
	out,err := exec.Command("notify-send","hai dsr").CombinedOutput()
	fmt.Println(string(out),err)


}
func test(){
	time.Sleep(1*time.Second)
	fmt.Println("sending")
	sample<- struct{}{}
}

