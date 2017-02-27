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
    	myPromoReq := Promo{}

        data, err := ioutil.ReadAll(r.Body)

        if err != nil {
            fmt.Println(err)
        }

        if err := proto.Unmarshal(data, &myPromoReq); err != nil {
            fmt.Println(err)
        }

        myPromo := &myPromoReq

        fmt.Printf("%+v", myPromo)

        // for _, mail := range myPromo.Inbox {
        //     fmt.Println(mail.RemoteEmail, ":", mail.Body)
        // }
    })
    
    fmt.Println("Listening....")
    http.ListenAndServe(":3000", nil)
}