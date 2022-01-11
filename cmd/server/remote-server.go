package main

import (
	"log"
	"net"

	"com.lsl/ss/server"
)

const PWD = "cGfU96a9o/sPM0bJ6xRPJiQCOSmfDF9UYNsAKhsxgfK2HWS1oAvQDXM10fVX5oz8g0Rr5WP2WkMGhupKXS83qDrTVSwJU2kwUT6Z7o90w5SJfpxJFnI9+oCwcZYY2QiIf1YixPMSEd2bs4XHk2IERdfc8RcofF6Hu6+Y6LoB1eyxeJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqS4j6eQlYRDYB2w/59btHhydolt5ofBlwFkhOyCOH6zLjeOLxUCnd2hYR63KikGuvOFO3g4TS88rMpBNxjbgzH2Xbs4VSGqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPhvGQ=="

func main() {
	server, err := server.NewServer(PWD, "127.0.0.1:8889")
	if err!=nil{
		log.Fatalln(err)
		return
	}

	server.Listen(func(listenAddr net.Addr) {
		log.Println("服务监听:", listenAddr.String())
	})

}
