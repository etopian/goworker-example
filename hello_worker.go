package main

import (

  "github.com/benmanns/goworker"
  "fmt"


)

func init() {


    settings := goworker.WorkerSettings{
        URI:            "redis://172.17.0.1:6379/",
        Connections:    1000,
        Queues:         []string{"urlqueue"},
        UseNumber:      true,
        ExitOnComplete: false,
        Concurrency:    1,
        Namespace:      "resque:",
        Interval:       1,
    }    
    
    goworker.SetSettings(settings)
  goworker.Register("GetUrl", getUrl)
  
  //goworker.Register("GetUrl", getUrl)
}



func getUrl(queue string, args ...interface{}) error {
  fmt.Println("Hello")
  return nil
}
