package itlssp

import (
	"github.com/pkg/errors"
	"github.com/tarm/serial"
)

type generic struct {
	unit
}

func NewGeneric(c *serial.Config) *generic {
	return &generic{
		unit: &device{
			seq:  0x80,
			conf: c,
		},
	}
}

func (this *generic) Reset() error {
	buf := []byte{byte(SspCmdReset)}
	_, err := this.unit.SendCommand(buf)
	return errors.WithStack(err)
}

func (this *generic) Sync() error {
	buf := []byte{byte(SspCmdSync)}
	_, err := this.unit.SendCommand(buf)
	return errors.WithStack(err)
}

func (this *generic) FirmwareVersion() error {
	buf := []byte{byte(SspCmdFirmwareVersion)}
	_, err := this.unit.SendCommand(buf)
	return errors.WithStack(err)
}

func (this *generic) HostProtocolVersion() error {
	buf := []byte{byte(SspCmdHostProtocolVersion)}
	_, err := this.unit.SendCommand(buf)
	return errors.WithStack(err)
}

func (this *generic) SetupRequest() error {
	buf := []byte{byte(SspCmdSetupRequest)}
	_, err := this.SendCommand(buf)

	return errors.WithStack(err)
}
