package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"

	"go.bug.st/serial.v1"
	"go.bug.st/serial.v1/enumerator"

	pb "../go-stuff/src/github.com/milvum/hummingbird/proto"
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
	serverPort        = flag.Int("port", 8081, "The server port")
	portName          = ""
	portID            = 0
	deviceID          = ""
	disableSerialRead = false
	c                 = make(chan string)
	d                 = make(chan string)
	config            configfile
	counter           = 0
	exitChannel       chan struct{}
	dir               string
)

func checkSerialPort(port *enumerator.PortDetails) {
	if port.IsUSB {
		fmt.Printf("   USB name   %s\n", port.Name)
		fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
		fmt.Printf("   USB serial %s\n\n", port.SerialNumber)
		// loop through array serialports in config.json, check combination vendor_id and product_id
		for i, device := range config.Serialports {
			if strings.ToLowerSpecial(nil, port.VID) == strings.ToLowerSpecial(nil, device.VendorID) && strings.ToLowerSpecial(nil, port.PID) == strings.ToLowerSpecial(nil, device.ProductID) {
				portID = i
				portName = port.Name
				color.Green("-- UART -- %s found at %s with index %v", device.Name, portName, portID)
				break
			}
		}
	}
}

func findSerialPort() bool {
	for {
		counter++
		ports, _ := enumerator.GetDetailedPortsList()
		for _, port := range ports {
			checkSerialPort(port)
		}
		if portName != "" {
			break
		}
		time.Sleep(time.Second * 1)
	}
	return true
}

func readSerial(port serial.Port) {
	// listen for messages on serial port
	buff := make([]byte, 256)
	for {
		if !disableSerialRead {
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
			color.Green("%s", msg)
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
		color.Yellow(msg)
		_, err := serialPort.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		disableSerialRead = false
		time.Sleep(time.Second * 1)
	}
}

func getPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		color.Red("-- Error -- Could not resolve current path")
	}
	color.Green("-- Init -- Current path is: %s", dir)
	return dir
}

func initConfiguration(dir string) bool {
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
	return true
}

func executeCommands() {
	time.Sleep(time.Second * 3)
	go hiberCommand("get_modem_info\r\n")
	time.Sleep(time.Second * 2)
	go hiberCommand("set_payload(7)\r\n")
	time.Sleep(time.Second * 1)
	go hiberCommand("testing")
	time.Sleep(time.Second * 10)
	executeCommands()
}

func openSerialPort() serial.Port {
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
	return port
}

type hiberbridgeServer struct {
}

func (s *hiberbridgeServer) BatchStatuses(stream pb.HiberBridge_BatchStatusesServer) error {
	point, err := stream.Recv()
	if err == io.EOF {
		return stream.SendAndClose(&pb.Reply{Success: true})
	}
	if err != nil {
		return err
	}
	fmt.Println(point.BinnedLocation.Lon)
	fmt.Println(point.BinnedLocation.Lat)
	return nil
}

func newServer() *hiberbridgeServer {
	s := &hiberbridgeServer{}
	return s
}

func startServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *serverPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHiberBridgeServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func main() {
	color.Yellow("-- Init -- Hiber Control script for Hummingbird v0.0.1")

	// Load configuration file from ./config.json
	dir = getPath()
	if initConfiguration(dir) {
		color.Green("-- IO -- Successfully imported config.json")
	}

	// Find a suitable connected USB device based on config
	// Only when suitable connected USB device is found, continue with code
	if findSerialPort() {
		color.Green("-- Serial -- Found serial port")
	}
	// -> portName = array Index in config.json/serialports

	// open serial port
	port := openSerialPort()
	defer port.Close()

	go startServer()

	// listen for messages on serial port, put in c channel
	go readSerial(port)

	go publishReadSerial()

	// write incoming data from d channel
	go writeSerial(port)

	go executeCommands()

	<-exitChannel
}
