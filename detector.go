package itlssp

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/tarm/serial"
	comport "go.bug.st/serial"
)

var ErrNoDeviceFound = errors.New("No device found!")

type UnitType byte

const (
	Validator   UnitType = 0x00
	SMARTHopper UnitType = 0x03
	SMARTPayout UnitType = 0x06
	NV11        UnitType = 0x07
)

func (u UnitType) String() string {
	switch u {
	case Validator:
		return "Validator"
	case SMARTHopper:
		return "SMART Hopper"
	case SMARTPayout:
		return "SMART Payout"
	case NV11:
		return "NV11"
	default:
		return "Unknown Type"
	}
}

type SSPConnection struct {
	Name string
	Addr byte
}

func (p *SSPConnection) String() string {
	return fmt.Sprintf(`{"Name":"%s","Addr":%d}`, p.Name, p.Addr)
}

func AvailablePorts() (ports []*SSPConnection) {
	names, err := comport.GetPortsList()
	if err == nil {
		for _, v := range names {
			ports = append(ports, &SSPConnection{Name: v, Addr: 0})
		}
	}
	return ports
}

type Unit struct {
	Type     UnitType
	Version  string
	Currency string
	Channels byte
}

type SSPDevice struct {
	Port *SSPConnection
	Unit *Unit
}

func SearchSSPDevices() (devices []*SSPDevice) {
	for _, port := range AvailablePorts() {
		if dvc, err := detect(port); err == nil {
			devices = append(devices, dvc)
		}
	}
	return devices
}

func detect(port *SSPConnection) (device *SSPDevice, err error) {
	cfg := &serial.Config{
		Name:        port.Name,
		Baud:        9600,
		ReadTimeout: time.Microsecond * 500,
		Size:        8,
		Parity:      0,
		StopBits:    2,
	}
	var com *serial.Port
	if com, err = serial.OpenPort(cfg); err != nil {
		return nil, errors.WithStack(err)
	}
	defer com.Close()

	if _, err = com.Write([]byte{xSTX, 0x80, 0x01, 0x05, 0x1D, 0x82}); err != nil {
		return nil, errors.WithStack(err)
	}

	var buf []byte
	if buf, err = readPort(com); err != nil {
		return nil, errors.WithStack(err)
	}

	var zero byte = 4
	if len(buf) < 35 || buf[zero-1] != 0xF0 {
		return nil, ErrNoDeviceFound
	}

	device = &SSPDevice{
		Port: port,
		Unit: &Unit{
			Type:     UnitType(buf[zero]),
			Version:  string(buf[zero+1 : zero+5]),
			Currency: string(buf[zero+5 : zero+8]),
			Channels: buf[zero+11],
		},
	}
	return device, nil
}

func readPort(r io.Reader) ([]byte, error) {
	var line []byte
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line = append(line, scanner.Bytes()...)
	}
	return line, errors.WithStack(scanner.Err())
}
