package main

import (
    "github.com/golang/protobuf/proto"
    "net/http"
    "fmt"
    "bytes"
    // Client "./clientStructure.pb.go"
)

func main() {
    myClient := Client {
        Id: proto.Int32(526), 
        Name: proto.String("John Doe"), 
        Email: proto.String("johndoe@example.com"), 
        Country: proto.String("US"),
    }
    clientInbox := make([]*Client_Mail, 0, 20)
    clientInbox = append(clientInbox,
        &Client_Mail {
            RemoteEmail: proto.String("jannetdoe@example.com"), 
            Body: proto.String("Hello. Greetings. Bye."),
        },
        &Client_Mail {
            RemoteEmail: proto.String("WilburDoe@example.com"), 
            Body: proto.String("Bye, Greetings, hello."),
        })

    myClient.Inbox = clientInbox
    
    data, err := proto.Marshal(&myClient)
    if err != nil {
        fmt.Println(err)
        return
    }

    // fmt.Printf("%+v", myClient)
    fmt.Printf("POST:\n%+v",&myClient)

    _, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))

    if err != nil {
        fmt.Println(err)
        return
    }
}