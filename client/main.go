package main

import (
    "github.com/golang/protobuf/proto"
    "net/http"
    "fmt"
    "bytes"
    // Client "./clientStructure.pb.go"
)

func main() {
    myPromo := Promo {
        Id: proto.Int32(1234), 
        Name: proto.String("Promo Gratis Ongkir"), 
        BaseCode: proto.String("GRATISONGKIR"), 
        IsGlobal: proto.Int32(1),
    }

    promoRules := make([]*Promo_Rule, 0, 20)
    promoRules = append(promoRules,
        &Promo_Rule {
            RuleId : proto.Int32(1),
            RuleName : proto.String("Minimum 200.000 IDR"),
            RuleValue : proto.String("200000"),
        },
        &Promo_Rule {
            RuleId : proto.Int32(2),
            RuleName : proto.String("Free Ongkir Max 20.000 IDR"),
            RuleValue : proto.String("20000"),
        })

    myPromo.Rules = promoRules

    promoTimeWindows := make([]*Promo_TimeWindow, 0, 20)
    promoTimeWindows = append(promoTimeWindows,
        &Promo_TimeWindow {
            TwId : proto.Int32(11),
            TwStart : proto.String("2017-01-01"),
            TwEnd : proto.String("2017-02-01"),
        },
        &Promo_TimeWindow {
            TwId : proto.Int32(12),
            TwStart : proto.String("2017-01-02"),
            TwEnd : proto.String("2017-02-02"),
        })

    myPromo.Timewindows = promoTimeWindows
    
    data, err := proto.Marshal(&myPromo)
    if err != nil {
        fmt.Println(err)
        return
    }

    // fmt.Printf("%+v", myClient)
    fmt.Printf("POST:\n%+v",&myPromo)

    _, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))

    if err != nil {
        fmt.Println(err)
        return
    }
}