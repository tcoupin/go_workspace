package main

import "os"
import "fmt"
import "cesar/crypto"
import "strconv"

func main(){
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s DECALAGE 'TXT'\n",os.Args[0])
		return
	}
	delta,e := strconv.Atoi(os.Args[1])
	if e != nil {
		fmt.Printf("%s\n",e)
		return
	}
	text  := string(os.Args[2])
	fmt.Printf("Cesar: %s\n",crypto.Cesar(text,delta))
}