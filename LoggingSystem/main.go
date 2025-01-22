package main

import (
	"logSys/constants"
)

func main() {
	cartlogger := NewLogger(NewConfig(constants.LogLevelINFO, constants.SinkTypeCONSOLE), "CART")
	orderlogger := NewLogger(NewConfig(constants.LogLevelDEBUG, constants.SinkTypeCONSOLE), "ORDER")

	cartlogger.Fatal("Panic occured while adding promo")
	cartlogger.Info("Successfully added item")
	cartlogger.Warn("Item stack empty")

	cartlogger.Config.SetLogLevel(constants.LogLevelWARN)
	cartlogger.Error("Failed to add item to cart")
	cartlogger.Debug("Debugging add item to cart flow")

	orderlogger.Fatal("Panic occured while payment")
	orderlogger.Info("Successfully placed order")
	orderlogger.Warn("Item stack empty, no items to checkout")
	orderlogger.Error("Failed to CHECKOUT")
	orderlogger.Debug("Debugging checkout flow")

	cartlogger.Show()
	orderlogger.Show()

}
