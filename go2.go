package main

import( 
	"os"
	"fmt"
)

func main(){
	for idx, element := range os.Args[1:]{
		fmt.Println(idx,"=>",element)
	}
	
    //fmt.Println(os.Args[0]) ex1.1
}

