package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    )

func main() {
     s := []string  {
        "http://www.cnn.com", 
        "http://www.facebook.com", 
        "http://www.darianhickman.com",
     }
     for  _, link := range s{
        getLink(link)
     }

}

func getLink( link string){
    response, err := http.Get(link)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s: %s\n", link, string(contents[0:140]))
    }
}