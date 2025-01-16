package main

import (
	"fmt"
	"sync"
)

func main() {
	AddConfig(LogLevelINFO, SinkTypeCONSOLE)
	WG := new(sync.WaitGroup)
	WG.Add(5)

	msg1 := NewMessage(LogLevelFATAL, "Panic occured while adding promo", "CART")
	go msg1.Log(WG)

	msg2 := NewMessage(LogLevelINFO, "Successfully added item", "CART")
	go msg2.Log(WG)

	msg3 := NewMessage(LogLevelWARN, "Item stack empty", "CART")
	go msg3.Log(WG)

	msg4 := NewMessage(LogLevelERROR, "Failed to add item to cart", "CART")
	go msg4.Log(WG)

	msg5 := NewMessage(LogLevelDEBUG, "Debugging add item to cart flow", "CART")
	go msg5.Log(WG)

	logConfig.SetLogLevel(LogLevelWARN, WG)
	WG.Add(5)

	msg1 = NewMessage(LogLevelFATAL, "Panic occured while payment", "ORDER")
	go msg1.Log(WG)

	msg2 = NewMessage(LogLevelINFO, "Successfully placed order", "ORDER")
	go msg2.Log(WG)

	msg3 = NewMessage(LogLevelWARN, "Item stack empty, no items to checkout", "ORDER")
	go msg3.Log(WG)

	msg4 = NewMessage(LogLevelERROR, "Failed to CHECKOUT", "ORDER")
	go msg4.Log(WG)

	msg5 = NewMessage(LogLevelDEBUG, "Debugging checkout flow", "ORDER")
	go msg5.Log(WG)

	WG.Wait()

	fmt.Println("---------------------------------------------------------------------")
	for i := 5; i > 0; i-- {
		fmt.Println(GetLogLevel(i))
		//logConfig.SinkLocation[i].WG.Wait()
		logConfig.SinkLocation[i].ShowMessages()
		fmt.Println("################################################################")
	}

}
