package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os/exec"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	pin, formattedPin := generatePin()
	info := accessory.Info{
		Name:         "USB Smart Switch",
		Manufacturer: "Dustin Mowcomber",
	}
	log.Printf("starting HomeKit device %q", info.Name)
	log.Printf("pin: %s", formattedPin)

	// add hc lightbulb
	acc := accessory.NewLightbulb(info)
	acc.Lightbulb.On.OnValueRemoteUpdate(func(isOn bool) {
		if isOn {
			on()
		} else {
			off()
		}
	})
	// default to on state
	acc.Lightbulb.On.SetValue(true)

	// run hc server
	t, err := hc.NewIPTransport(hc.Config{Pin: pin}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}
	hc.OnTermination(func() {
		<-t.Stop()
	})
	t.Start()
}

func off() {
	if err := exec.Command("uhubctl", "-a", "off", "-l", "1-1").Run(); err != nil {
		log.Printf("failed to turn off usb: %s", err.Error())
	}
}

func on() {
	if err := exec.Command("uhubctl", "-a", "on", "-l", "1-1").Run(); err != nil {
		log.Printf("failed to turn on usb: %s", err.Error())
	}
}

func generatePin() (string, string) {
	rand.Seed(time.Now().UnixNano())
	// generate an 8 digit number
	num := randNDigits(8)
	pin := fmt.Sprintf("%d", num)
	// ignoring the error as I've tested the above pin generation to always be 8 digits
	formattedPin, _ := hc.ValidatePin(pin)
	return pin, formattedPin
}

func randNDigits(digits int) int {
	rand.Seed(time.Now().UnixNano())
	low := int(math.Pow(10, float64(digits)-1))
	high := int(math.Pow(10, float64(digits)))
	return low + rand.Intn(high-low)
}
