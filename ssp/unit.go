package ssp

import (
	"bufio"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/tarm/serial"
	"io"
	reflect "reflect"
)

var (
	ErrSspWriteCommand = errors.New("Error write SSP command")
)

type unit interface {
	Open(*serial.Config) error
	Close() error
	SendCommand(data []byte) ([]byte, error)
}

type device struct {
	seq  byte
	conf *serial.Config
	port *serial.Port
}

// Open serial port
func (this *device) Open(cfg *serial.Config) error {
	var err error
	if this.port, err = serial.OpenPort(cfg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Close serial port
func (this *device) Close() error {
	return this.port.Close()
}

// SendCommand sends data and checks error response
func (this *device) SendCommand(data []byte) ([]byte, error) {
	var err error
	var pkg []byte
	var buf []byte
	if pkg, err = this.send(this.pack(data)); err != nil {
		return nil, errors.WithStack(err)
	}
	if buf, err = this.unpack(pkg); err != nil {
		return nil, errors.WithStack(err)
	}
	if err = this.checkResponse(buf); err != nil {
		return nil, errors.WithStack(err)
	}
	return buf, nil
}

// send write buffer of bytes to serial port and return reading data
func (this *device) send(buff []byte) ([]byte, error) {
	n, err := this.port.Write(buff)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if n != len(buff) {
		return nil, ErrSspWriteCommand
	}
	return this.read(this.port)
}

// checkResponse check the answer for errors
func (this *device) checkResponse(data []byte) error {
	code := SSPResponse(data[0])
	if code != SspResponseOk {
		if code == SspResponseCannotProcess {
			if data[1] == 0x03 {
				return errors.New("Validator has responded with \"Busy\", command cannot be processed at this time")
			} else {
				return errors.Errorf("Command response is CANNOT PROCESS COMMAND, error code - 0x%02X", data[1])
			}
		}
		return errors.New(code.String())
	}
	return nil
}

// read bytes from serial port
func (this *device) read(r io.Reader) ([]byte, error) {

	var buff []byte
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		buff = append(buff, scanner.Bytes()...)
	}
	log.Debug().Msgf("read: %X - %X", buff, scanner.Bytes())
	return buff, errors.WithStack(scanner.Err())
}

// unpack data from the received packet
func (this *device) unpack(data []byte) ([]byte, error) {
	if len(data) < 6 {
		return nil, errors.Errorf("Invalid data packet size (%d): %X", len(data), data)
	}
	if data[0] != xSTX {
		return nil, errors.Errorf("Invalid data packet format: %X", data)
	}

	crc := data[len(data)-2:] // crc16
	if !reflect.DeepEqual(crc, crc16Bytes(data[1:len(data)-2])) {
		return nil, errors.Errorf("Invalid packet checksum 0x%04X", crc)
	}

	return data[3 : len(data)-2], nil
}

// pack data into a package for sending
func (this *device) pack(data []byte) []byte {
	res := append([]byte{xSTX, this.getSEQ(), byte(len(data))}, data...)
	return append(res, crc16Bytes(res[1:])...)
}

// checkSTX checks the buffer for an entry value
// Byte stuffing is used to encode any xSTX bytes that are included in the data to be transmitted. If 0x7F (xSTX)
// appears in the data to be transmitted then it should be replaced by 0x7F, 0x7F.
// Byte stuffing is done after the CRC is calculated, the CRC its self can be byte stuffed. The maximum length of
// data is 0xFF bytes.
func (this *device) checkSTX(data []byte) []byte {
	idx := 0
	var res []byte
	for i := 0; i < len(data); i++ {
		if data[i] == xSTX {
			res = append(res, data[idx:i+1]...)
			res = append(res, xSTX)
			idx = i + 1
		}
	}
	if idx <= len(data) {
		res = append(res, data[idx:]...)
	}
	return res
}

// getSEQ receive next value SEQ
// The sequence flag is used to allow the slave to determine whether a packet is a re-transmission due to its last
// reply being lost. Each time the master sends a new packet to a slave it alternates the sequence flag. If a slave
// receives a packet with the same sequence flag as the last one, it does not execute the command but simply
// repeats it's last reply. In a reply packet the address and sequence flag match the command packet.
func (this *device) getSEQ() byte {
	val := this.seq << 7
	if val == 0 {
		this.seq = 1
	} else {
		this.seq = val
	}
	return val
}
