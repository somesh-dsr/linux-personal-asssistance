package main

import (
	"fmt"
	"os"
	"regexp"
)

type p21 struct {
	sample string
}
type p2 struct {
	data string
	*p21
}
func main() {
	ref := &p2{
		data : "sample",
		p21: &p21{
			sample: "sample",
		},
	}
	change(ref)
	fmt.Println(ref.p21.sample,ref)

	regexpRef,_ := regexp.Compile(`(?i)(schedule|set)\s*(a|an|the)?\s*(meeting|reminder)\s*\s*(at|by)\s*(\d\d:\d\d)`)
	res := regexpRef.FindAllStringSubmatch(`schedule meeting at 10:30`,-1)
	fmt.Println(res[0][len(res[0])-1])
	path,err := os.Getwd()
	fmt.Println(path,err)
}
func change(samp interface{}){
	res := samp.(*p2)
	fmt.Println(res.data)
	fmt.Println(res)
	res.data = "DSR"
	res.p21.sample="sample_DSR"

}