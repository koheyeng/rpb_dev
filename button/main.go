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
	button18 := gpio.NewButtonDriver(r, "12", 0)
	led40 := gpio.NewLedDriver(r, "40")

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			led40.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{button18, led40},
		work,
	)

	go robot.Start()
	robot.Stop()
	fmt.Println("bye!")
}
