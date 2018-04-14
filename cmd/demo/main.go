package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/mh-cbon/seq"
)

func main() {

	var many = flag.Bool("many", false, "many messages or not")
	flag.Parse()

	var enc = seq.RFC7464(os.Stdout)

	demoData := map[string]interface{}{
		"key":  "value",
		"key2": "value2",
	}

	ticker := time.NewTicker(time.Millisecond * 125)
	for {
		demoData["time"] = time.Now().UnixNano()
		if err := enc.Encode(demoData); err != nil {
			log.Printf("err=%v", err)
		}
		<-ticker.C
		if *many == false {
			break
		}
	}
}
