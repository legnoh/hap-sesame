package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/legnoh/go-sesame/sesame"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version      string
	revision     string
	name         string
	serialNumber string
	token        = kingpin.Flag("token", "your sesame api token").Short('t').Required().String()
	deviceId     = kingpin.Flag("device-id", "your sesame device id").Short('d').Default("auto-detected").String()
	pin          = kingpin.Flag("pin", "homekit device pin for connect").Short('p').Default("00102003").String()
	storagePath  = kingpin.Flag("storage-path", "directory for configfiles").Default("hc-sesame").String()
)

func main() {
	kingpin.Version(version + "-" + revision)
	kingpin.Parse()

	c := sesame.NewClient(*token)
	list := c.GetSesameList()

	if *deviceId == "auto-detected" {
		name = list[0].Nickname
		*deviceId = list[0].DeviceId
		serialNumber = list[0].Serial
	} else {
		for i := 0; i < len(list); i++ {
			if *deviceId == list[i].DeviceId {
				name = list[i].Nickname
				serialNumber = list[i].Serial
				break
			}
		}
	}

	info := accessory.Info{
		Name:         name,
		SerialNumber: serialNumber,
		Manufacturer: "CANDY HOUSE, Inc.",
		Model:        "SESAME",
	}

	config := hc.Config{
		Pin:         *pin,
		StoragePath: *storagePath,
	}

	// acc := accessory.NewRockManagement(info)

	// acc.RockManagement.On.OnValueRemoteUpdate(func(on bool) {
	// 	if on == true {
	// 		c.ControlSesame(*deviceId, "unlock")
	// 		log.Println(name + ": Turn unlock")
	// 	} else {
	// 		c.ControlSesame(*deviceId, "lock")
	// 		log.Println(name + ": Turn lock")
	// 	}
	// })

	// t, err := hc.NewIPTransport(config, acc.Accessory)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// hc.OnTermination(func() {
	// 	<-t.Stop()
	// })

	// t.Start()
}
