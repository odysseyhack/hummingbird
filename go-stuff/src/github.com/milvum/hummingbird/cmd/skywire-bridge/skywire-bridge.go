package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	proto "github.com/golang/protobuf/proto"
	pb "github.com/milvum/hummingbird/proto"
	"github.com/skycoin/skywire/pkg/app"
	"github.com/skycoin/skywire/pkg/cipher"
)

var (
	bridgeApp   *app.App
	connsMu     sync.Mutex
	connections map[cipher.PubKey]net.Conn
)

func handleConn(conn net.Conn) {
	// Hard assume the only thing we will be sending and receiving
	// is bytes that are _actually_ MessageRequest as bytestrings.
	for {
		buf := make([]byte, 10*1024) // XXX: 10kb big enough?
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Failed to read package: ", err, " but did read ", n, " bytes")
			return
		}
		serialized := pb.MessageRequest{}
		err = proto.Unmarshal(buf, &serialized)
		if err != nil {
			log.Println("unmarshaling error: ", err)
		}
		// TODO: Do something cool with this (forward/drop/uplink, I guess)
	}
}

// We should probably add some application-level security here, but oh well...
// Basically, accept any incoming connection
func listenLoop() {
	for {
		conn, err := bridgeApp.Accept()
		if err != nil {
			log.Println("failed to accept conn: ", err)
			return
		}

		raddr := conn.RemoteAddr().(*app.Addr)
		connsMu.Lock()
		connections[raddr.PubKey] = conn
		connsMu.Unlock()

		go handleConn(conn)
	}
}

func openConnection(pk cipher.PubKey) {
	connsMu.Lock()
	conn := connections[pk]
	connsMu.Unlock()

	if conn == nil {
		var err error
		addr := &app.Addr{PubKey: pk, Port: 1}
		conn, err = bridgeApp.Dial(addr)
		if err != nil {
			// TODO some kind of awesome error thing
			log.Println("Failed to dial to ", addr)
			return
		}

		connsMu.Lock()
		connections[pk] = conn
		connsMu.Unlock()

		go handleConn(conn)
	}
}

func broadcast(messageRequest pb.MessageRequest) {
	data, err := proto.Marshal(&messageRequest)
	if err != nil {
		log.Println("marshaling error: ", err)
		return
	}

	connsMu.Lock()
	// This is _probably_ an extreme anti-pattern.
	defer connsMu.Unlock()
	for pk, conn := range connections {

		if _, err := conn.Write(data); err != nil {
			log.Println("Error writing data to connection for pk: ", pk)
		} else {
			log.Println("Good write to connection for pk: ", pk)
		}
	}
}

func main() {
	config := &app.Config{AppName: "bridge", AppVersion: "1.0", ProtocolVersion: "0.0.1"}
	bridgeApp, err := app.Setup(config)
	if err != nil {
		log.Fatal("Setup failure: ", err)
	}
	defer bridgeApp.Close()

	connections = make(map[cipher.PubKey]net.Conn)
	fmt.Println("hello world")
}
