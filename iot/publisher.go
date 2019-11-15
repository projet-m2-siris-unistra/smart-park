package main

import (
    //"context"
    //"io/ioutil"
    "encoding/json"
    "io"
    "log"
    "net/http"
    //"time"
    "sync"
    "os"
    "os/signal"
    "fmt"
    "math/rand"
    "strings"
    "strconv"
    "time"
)

import nats "github.com/nats-io/nats.go"


// sensor structure

type infos struct {
  Id        int  `json:"id"`
  Occupied  bool `json:"occupied"`
  Battery   int  `json:"battery"`
}


// return true or false

func rand2() bool {
    return rand.Int31()&0x01 == 0
}

var nc *nats.Conn

// message flood

func flood(id int) {
    for {
        bat := rand.Intn(1500)
        occ := rand2()

        test := infos{
            Id: id,
            Occupied: occ,
            Battery: bat,
        }

        requestBody, err := json.Marshal(test)

        if err != nil{
            log.Fatalln(err)
        }

        err = nc.Publish("devices", requestBody)
        if err != nil {
            log.Fatal(err)
        }
        time.Sleep(time.Minute)
    }
}

func main() {
    var err error
    nc, err = nats.Connect("nats://127.0.0.1:4222")
    if err != nil {
        log.Fatal(err)
    }

    ids := os.Getenv("DEVICE_IDS")
    for _, id := range strings.Split(ids, ",") {
        i, err := strconv.Atoi(id)
        if err != nil {
            log.Fatal(err)
        }
        
        go flood(i)
    }


    // handle SIGINT signal
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c


    // close connection
    nc.Close()
}
