package main

import (
    "github.com/golang/protobuf/proto"
    "net/http"
    "fmt"
    "io/ioutil"
    // "reflect"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	myClientReq := Client{}

        data, err := ioutil.ReadAll(r.Body)

        if err != nil {
            fmt.Println(err)
        }

        if err := proto.Unmarshal(data, &myClientReq); err != nil {
            fmt.Println(err)
        }

        myClient := &myClientReq

        fmt.Printf("%+v", myClient)

        for _, mail := range myClient.Inbox {
            fmt.Println(mail.RemoteEmail, ":", mail.Body)
        }
    })
    
    fmt.Println("Listening....")
    http.ListenAndServe(":3000", nil)
}