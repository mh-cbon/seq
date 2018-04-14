# sequence encoder &nbsp;[![GoDoc](https://godoc.org/github.com/mh-cbon/seq?status.svg)](https://godoc.org/github.com/mh-cbon/seq)

Almost implements `RFC7464` to realize `application/json-seq` streaming.

see: https://tools.ietf.org/html/rfc7464#section-2

# demo

```sh
$ go run cmd/demo/main.go
{"key":"value","key2":"value2","time":1523718306771451031}

$ go run cmd/demo/main.go | hexdump -C
00000000  1e 7b 22 6b 65 79 22 3a  22 76 61 6c 75 65 22 2c  |.{"key":"value",|
00000010  22 6b 65 79 32 22 3a 22  76 61 6c 75 65 32 22 2c  |"key2":"value2",|
00000020  22 74 69 6d 65 22 3a 31  35 32 33 37 31 38 33 31  |"time":152371831|
00000030  39 39 32 37 35 38 30 32  39 38 7d 0a              |9927580298}.|
0000003c
```

### TODO

    $ astitodo -f json . | jq -r '.[] | "\(.Filename):\(.Line)\n\t \(.Message[0])"'
    encoder.go:33
    	 It does not escape separators from the encoded data.
