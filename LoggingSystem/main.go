package main

import (
	"fmt"
)

func main() {
	AddConfig(LogLevelINFO, SinkTypeCONSOLE, "CART")
	AddConfig(LogLevelDEBUG, SinkTypeCONSOLE, "ORDER")

	msg1 := NewMessage(LogLevelFATAL, "Panic occured while adding promo", "CART")
	msg1.Log()

	msg2 := NewMessage(LogLevelINFO, "Successfully added item", "CART")
	msg2.Log()

	msg3 := NewMessage(LogLevelWARN, "Item stack empty", "CART")
	msg3.Log()

	logConfig["CART"].SetLogLevel(LogLevelWARN)

	msg4 := NewMessage(LogLevelERROR, "Failed to add item to cart", "CART")
	msg4.Log()

	msg5 := NewMessage(LogLevelDEBUG, "Debugging add item to cart flow", "CART")
	msg5.Log()

	msg1 = NewMessage(LogLevelFATAL, "Panic occured while payment", "ORDER")
	msg1.Log()

	msg2 = NewMessage(LogLevelINFO, "Successfully placed order", "ORDER")
	msg2.Log()

	msg3 = NewMessage(LogLevelWARN, "Item stack empty, no items to checkout", "ORDER")
	msg3.Log()

	msg4 = NewMessage(LogLevelERROR, "Failed to CHECKOUT", "ORDER")
	msg4.Log()

	msg5 = NewMessage(LogLevelDEBUG, "Debugging checkout flow", "ORDER")
	msg5.Log()

	for key := range logConfig {
		logConfig[key].SinkLocation.WG.Wait()
		fmt.Println("---------------------------------------------------------------------")
		fmt.Println(key)
		logConfig[key].SinkLocation.ShowMessages()
		fmt.Println("################################################################")
	}

}
