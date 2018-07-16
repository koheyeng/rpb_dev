package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	button12 := gpio.NewButtonDriver(r, "12")
	led40 := gpio.NewLedDriver(r, "40")

	work := func() {
		button12.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button 12 pressed")
			led40.On()
		})
		button12.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button 12 pressed")
			led40.Off()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led40, button12},
		work,
	)

	go robot.Start()
}
