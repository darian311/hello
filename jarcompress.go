package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
	"encoding/csv"


)


func main(){
	fmt.Print(os.Args[1], os.Args[2])

	f, err := os.Create(os.Args[2])
	if err != nil{
		fmt.Printf("%s", err)
        os.Exit(1)		
	}

	w := csv.NewWriter(f)

	defer	f.Close()
	 
	//s := []string{ "Puppy" }
	err = w.Write([]string{ "Puppy" })
	if err != nil{ fmt.Print("Cannot write to file", err)}
	
	w.Flush()

	
}

// for each folder do a go routine. 
func listjars(){

	

}


