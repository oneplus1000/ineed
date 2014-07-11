package main

import(
	"github.com/oneplus1000/ineed"
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	var err error
	var dir string
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	//fmt.Printf("%s\n",dir)
	var n ineed.Need
	err = n.Init(dir)
	if err != nil {
		fmt.Printf("error:%s\n",err.Error())
	}

	var args []string
	for i,arg := range os.Args {
		//fmt.Printf("%s\n",arg)
		if i > 0 {
			args = append(args,arg)
		}
	}

	err = n.Run(args)
	if err != nil {
		fmt.Printf("error:%s\n",err.Error())
	}
}
