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
	led40 := gpio.NewLedDriver(r, "40")

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			led40.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led40},
		work,
	)

	go robot.Start()
	time.Sleep(60 * time.Second)
	robot.Stop()
	fmt.Println("bye!")
}
