package server

import (
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/sirupsen/logrus"
)

var log = logrus.New().WithFields(logrus.Fields{"mode": "server"})

func handler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(logrus.Fields{
		"local":  conn.LocalAddr().String(),
		"remote": conn.RemoteAddr().String(),
	}).Info("connection established")

	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Error(err)
				break
			}

			log.WithFields(logrus.Fields{"msg": string(msg), "op": op}).Info("client message")
			err = wsutil.WriteServerMessage(conn, ws.OpText, []byte("OK"))
			if err != nil {
				log.Error(err)
			}
		}
	}()
}

// Start runs server runtime
func Start(addr string) {
	log.Info("server start")

	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(handler)))

	log.Info("server stop")
}
