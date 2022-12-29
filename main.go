package main

import (
	"log"
	"time"
	"github.com/curzodo/plugo"
)

func main() {
	p, err := plugo.New("ExamplePlugin")

	if err != nil {
		log.Println(err)
	}

	p.Ready("")

	for {
		time.Sleep(5 * time.Second)

		if !p.CheckConnection("redstone") {
			p.Shutdown()
			return
		}

		resp, _ := p.Call("redstone", "Players")

		p.Println(resp[0])
	}
}
