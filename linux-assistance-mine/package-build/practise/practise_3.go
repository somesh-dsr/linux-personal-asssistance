package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	displayMap := make(map[string]string)
	displayMap["dsr"] = "GDSR"
	displayMap["dsr2"] = "GDSR2"
	var slice []int8
	slice = append(slice,12,13,14,15)
	//ref,err := os.OpenFile("/tmp/scheduletest",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	//fmt.Println(err)
	timer := time.Now()
	fmt.Println(timer.Hour())
	fmt.Println(timer.Minute())
	fmt.Println(timer.Zone())
	fmt.Println(time.Now())
	fmt.Println(timer.Zone())
	ref,_ := os.Stat("/etc")
	fmt.Println(ref.Sys())
	fmt.Println(ref.Mode())
	fmt.Println(ref)

}
type validate_struct struct {
	data string
	communicate chan struct{}
}
func validate(){
	ref := &validate_struct{
		data: "",
		communicate: make(chan struct{}),
	}
	data := <-ref.communicate
	fmt.Println(data)


}
func (ref *validate_struct)pushdata(){


}
