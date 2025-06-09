package main

import (
	"fmt"

	"github.com/frida/frida-go/frida"
)

func main() {
	manager := frida.NewDeviceManager()
	devices, err := manager.EnumerateDevices()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[*] Frida version: %s\n", frida.Version())
	fmt.Println("[*] Devices: ")
	for _, device := range devices {
		fmt.Printf("[*] Found device %s (%s)\n", device.Name(), device.ID())
	}
}
