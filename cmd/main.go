package main

import (
	"bloom"
	"fmt"
	"log"
)


func main() {
	bf, err := bloom.New(5)
	if err != nil {
		log.Fatal(err)
	}

	if err := bf.Put([]byte("foo")); err != nil {
		log.Fatal(err)
	}
	fmt.Println(bf.MightContain([]byte("foo")))

	bf.Put([]byte("bar"))
	fmt.Println(bf.MightContain([]byte("bar")))

	fmt.Println(bf.MightContain([]byte("baz")))
	fmt.Println(bf.MightContain([]byte("qux")))
}
