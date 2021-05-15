
package main

import (
	"fmt"
	processor "linux-assistance/engine"
	events "linux-assistance/handlers"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
		case "welcome": processor.ProcessRequests("welcome")
		case "reminder" : processor.ProcessRequests("reminder")
		case "generic": processor.ProcessRequests("generic")
	default:
		fmt.Println("Under implementation")
		return
		//default:processor.ProcessRequests("generic")
	}
	events.ActivityStatus = make(chan struct{})
	//for {
		//time.Sleep(10*time.Second)
		select {
		case <- time.After(2*time.Minute):
			if checkUpdateChannel(){
				return
			}
		case <-events.ActivityStatus:
			if checkUpdateChannel(){
				return
			}
		case <-events.Update:
			return
		}
	//}

}
func checkUpdateChannel()bool{
	if events.Update == nil {
		return true
	}
	return false
}
