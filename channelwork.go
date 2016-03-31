package main

import(
	"fmt"
	_"os"
	_"sort"
)

func main(){
	s := []string  {
        "http://www.cnn.com", 
        "http://www.facebook.com", 
        "http://www.darianhickman.com",
     }
	//ci := make(chan int)            // unbuffered channel of integers
	//cj := make(chan int, 0)         // unbuffered channel of integers
	//cs := make(chan *os.File, 100)  // buffered channel of pointers to Files	
	_ = s

	c := make(chan int)  // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
	    fmt.Printf("Channel pushing\n")
	    c <- 1  // Send a signal; value does not matter.
	}()
	fmt.Printf("Make something happen then see what the channel returns\n")
	<-c   // Wait for sort to finish; discard sent value.
	fmt.Printf("Channel returned.\n")
	
}
