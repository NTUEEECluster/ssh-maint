package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gliderlabs/ssh"
)

func main() {
	message := flag.String("message", "/etc/ssh-maint/message", "Path to file with the message to write")
	hostLoc := flag.String("host-key", "/etc/ssh-maint/host-key", "Path to the host key")
	listenAddr := flag.String("listen", ":2222", "Address to listen to")
	flag.Parse()

	f, err := os.ReadFile(*message)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to open message:", err.Error())
		os.Exit(1)
	}

	ssh.Handle(func(s ssh.Session) {
		_, _ = s.Write(f)
	})

	log.Fatal(ssh.ListenAndServe(*listenAddr, nil, ssh.HostKeyFile(*hostLoc)))
}
