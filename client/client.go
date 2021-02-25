package client

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/sirupsen/logrus"
)

var log = logrus.New().WithFields(logrus.Fields{"mode": "client"})

func connect(addr string) net.Conn {
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), addr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func send(conn net.Conn, msg string) {
	err := wsutil.WriteClientMessage(conn, ws.OpText, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func readline(reader *bufio.Reader) string {
	fmt.Print("Send: ")
	line, err := reader.ReadString('\n')

	if err != nil {
		log.Error(err)
		return readline(reader)
	}

	return line[:len(line)-1]
}

// Start runs client runtime
func Start(addr string) {
	log.Info("client start")

	conn := connect(addr)
	reader := bufio.NewReader(os.Stdin)

	for {
		line := readline(reader)

		if line == "q" {
			break
		}

		send(conn, line)
	}

	log.Info("client stop")
}
