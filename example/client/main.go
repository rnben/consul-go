package main

import (
	"fmt"
	"github.com/ruben-zhi/consul-go"
	"log"
)

func main() {
	resolver := consul_go.NewConsulResolver("127.0.0.1:8500", "grpc-go")
	resolve, err := resolver.Resolve("")
	if err != nil {
		log.Fatalln(err)
	}
	next, err := resolve.Next()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(next))
	for _, n := range next {
		fmt.Println(n.Addr)
	}
}
