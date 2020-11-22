package main

import (
	"os"
	"time"

	"github.com/charoit/itlssp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tarm/serial"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Start driver")

	//log.Info().Interface("Ports", itl.AvailablePorts()).Send()
	//log.Info().Interface("Devices", itl.SearchSSPDevices()).Send()

	ports := itlssp.AvailablePorts()
	cfg := &serial.Config{
		Name:        ports[0].Name,
		Baud:        9600,
		ReadTimeout: time.Microsecond * 500,
		Size:        8,
		Parity:      0,
		StopBits:    2,
	}
	gen := itlssp.NewGeneric(cfg)
	if err := gen.Open(cfg); err != nil {
		log.Fatal().Err(err).Send()
	}
	defer gen.Close()

	if err := gen.HostProtocolVersion(); err != nil {
		log.Fatal().Err(err).Send()
	}

	//c := gen.Pkg([]byte{0x06,0x06})
	//fmt.Printf("%X\n",c)
	//c = gen.Pkg([]byte{0x06,0x06})
	//fmt.Printf("%X\n",c)

}
