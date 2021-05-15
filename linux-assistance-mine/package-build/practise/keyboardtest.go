package main

import (
	_"fmt"
	"os"
)

func main() {
	ref,_ := os.OpenFile("/tmp/test",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	 ref.Write([]byte(os.Args[0]+"  "+os.Args[1]+"\n"))

	ref.Close()

}
