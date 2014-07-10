package main

import(
	"github.com/oneplus1000/ineed"
	"fmt"
)

func main(){
	var n ineed.Need
	err := n.Init("/data/CODES/Tmp/Test01")
	if err != nil {
		fmt.Printf("error:%s\n",err.Error())
	}
	err = n.Run([]string{"status","Test01"})
	if err != nil {
		fmt.Printf("error:%s\n",err.Error())
	}
}
