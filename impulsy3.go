package impulsy3

import (
	"fmt"

	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/beaglebone"
)

const (
	gen1 = "P9_16"
	gen2 = "P9_14"
	gen3 = "P8_13"
)

var (
	impulses [3]int
	flowsec1 int
	flowsec2 int
	flowsec3 int
)

// func main() {
func RunImpulses() {
	beagleboneAdaptor := beaglebone.NewAdaptor()

	pin1 := gpio.NewButtonDriver(beagleboneAdaptor, gen1)
	pin2 := gpio.NewButtonDriver(beagleboneAdaptor, gen2)
	pin3 := gpio.NewButtonDriver(beagleboneAdaptor, gen3)

	pulseCallback := func(pin string) {
		switch pin {
		case gen1:
			impulses[0]++
		case gen2:
			impulses[1]++
		case gen3:
			impulses[2]++
		}
	}

	pin1.On(gpio.ButtonPush, func(data interface{}) {
		pulseCallback(gen1)
	})
	pin2.On(gpio.ButtonPush, func(data interface{}) {
		pulseCallback(gen2)
	})
	pin3.On(gpio.ButtonPush, func(data interface{}) {
		pulseCallback(gen3)
	})

	robot := gobot.NewRobot("impulsesBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{pin1, pin2, pin3},
		func() {
			licznik := 0
			for {
				startTime := time.Now()
				for time.Since(startTime) < 1*time.Second {
					time.Sleep(10 * time.Millisecond)
				}
				flowsec1 = impulses[0]
				flowsec2 = impulses[1]
				flowsec3 = impulses[2]
				//				fmt.Println("Impulses:", impulses)
				fmt.Println(flowsec1, flowsec2, flowsec3)

				impulses = [3]int{}
				licznik++
				if licznik == 1000 {
					break
				}
			}
		},
	)

	robot.Start()
}
