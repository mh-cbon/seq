package seq_test

import (
	"log"
	"os"

	"github.com/mh-cbon/seq"
)

func ExampleRFC7464() {
	var enc = seq.RFC7464(os.Stdout)
	var some = map[string]interface{}{"key": "value"}

	if err := enc.Encode(some); err != nil {
		log.Printf("err=%v", err)
	}
	// Output: {"key":"value"}
	//
}
