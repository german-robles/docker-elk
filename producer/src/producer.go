package main

import (
	"net"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "app-ger-test"}))

	if err != nil {
		log.Fatal(err)
	}
	log.Hooks.Add(hook)
	ctx := log.WithFields(logrus.Fields{
		"name":   "foo",
		"action": "delete_user",
		"for":    "bar",
	})
	ctx.Info("User deleted")
}
