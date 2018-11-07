package main

import (
	"fmt"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"github.com/grokify/alexa-skill-ringcentral-meetings-util/rcroomsoffice365"
)

type opts struct {
}

func main() {
	opts := rcroomsoffice365.RingCentralRoomsOffice365Data{}

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(opts)

	cmds := rcroomsoffice365.RingCentralRoomsOffice365(opts)
	fmt.Println("-----BEGIN POWER SHELL COMMANDS-----")
	fmt.Println(cmds)
	fmt.Println("-----END POWER SHELL COMMANDS-----")
}
