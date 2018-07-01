package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led7 := gpio.NewLedDriver(r, "7")
	led11 := gpio.NewLedDriver(r, "11")

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			led7.Toggle()
			led11.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led7},
		[]gobot.Device{led11},
		work,
	)

	go robot.Start()
	fmt.Println("start toggle led!")
	time.Sleep(10 * time.Second)

	robot.Stop()
	fmt.Println("stop toggle led!")
	fmt.Println("bye!")
}
