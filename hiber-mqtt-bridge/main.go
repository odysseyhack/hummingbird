package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"

	"go.bug.st/serial.v1"
	"go.bug.st/serial.v1/enumerator"
)

type configfile struct {
	serialports
}

type serialports struct {
	Serialports []serialport `json:"serialports"`
}

type serialport struct {
	Name      string `json:"name"`
	VendorID  string `json:"vendor_id"`
	ProductID string `json:"product_id"`
	Baudrate  int    `json:"baudrate"`
}

var (
	portName          = ""
	portID            = 0
	deviceID          = ""
	disableSerialRead = false
	c                 = make(chan string)
	d                 = make(chan string)
	config            configfile
)

func findSerialPort() {
	counter := 0
	for {
		counter++
		ports, err := enumerator.GetDetailedPortsList()
		if err != nil {
			log.Fatal(err)
		}
		for _, port := range ports {
			if port.IsUSB {
				fmt.Printf("   USB name   %s\n", port.Name)
				fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
				fmt.Printf("   USB serial %s\n\n", port.SerialNumber)
				// loop through array serialports in config.json, check combination vendor_id and product_id
				for i, device := range config.Serialports {
					if strings.ToLowerSpecial(nil, port.VID) == strings.ToLowerSpecial(nil, device.VendorID) && strings.ToLowerSpecial(nil, port.PID) == strings.ToLowerSpecial(nil, device.ProductID) {
						portID = i + 1
						portName = port.Name
						color.Green("-- UART -- %s found at %s with index %v", device.Name, portName, portID)
						break
					}
				}
			}
		}
		if portName == "" {
			if counter == 1 {
				color.White("-- IO -- Press [CTRL+C] to close program")
				color.White("-- IO -- Plug-in Hiberfaker USB device to continue...")
			}
		} else {
			break
		}
		time.Sleep(time.Second * 1)
	}
}

func readSerial(port serial.Port) {
	// listen for messages on serial port
	buff := make([]byte, 256)
	for {
		if !disableSerialRead {
			fmt.Println("Reading input")
			n, err := port.Read(buff)
			if err != nil {
				color.Red("Something went wrong, exiting...")
				panic(err)
			}
			c <- string(buff[:n])
			time.Sleep(time.Second * 1 / 3)
		}
	}
}

func hiberCommand(cmd string) {
	d <- cmd
}

func publishReadSerial() {
	for {
		msg := <-c
		if len(msg) >= 2 {
			// Output to terminal
			color.White("%s", msg)
		}
	}
}

func writeSerial(serialPort serial.Port) {
	for {
		msg := <-d
		if msg != "" {
			// disable reading
			disableSerialRead = true
		}
		fmt.Println(msg)
		_, err := serialPort.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		disableSerialRead = false
		time.Sleep(time.Second * 1)
	}
}

func initConfiguration() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		color.Red("-- Error -- Could not resolve current path")
	}
	color.Green("-- Init -- Current path is: %s", dir)

	jsonFile, err := os.Open(fmt.Sprintf("%s/config.json", dir))
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println(err)
	}
	color.Green("-- IO -- Successfully imported config.json")
}

func main() {
	color.Yellow("-- Init -- Hiber Control script for Hummingbird v0.0.1")

	// Load configuration file from ./config.json
	initConfiguration()

	// Find a suitable connected USB device based on config
	// Only when suitable connected USB device is found, continue with code
	findSerialPort()
	// -> portName = array Index in config.json/serialports

	var options = &serial.Mode{
		BaudRate: config.Serialports[portID].Baudrate,
	}

	port, err := serial.Open(portName, options)
	if err != nil {
		log.Fatal(err)
		color.Red("-- UART -- Port %s could not be opened", portName)
	} else {
		color.Green("-- UART -- Port %s opened with baudrate %v", portName, config.Serialports[portID].Baudrate)
	}
	defer port.Close()

	// listen for messages on serial port, put in c channel
	go readSerial(port)

	go publishReadSerial()

	// write incoming data from d channel
	go writeSerial(port)

	go hiberCommand("get_modem_info\r\n")

	var input string
	fmt.Scanln(&input)
}
