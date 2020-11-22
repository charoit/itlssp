package itlssp

import (
	"testing"
)

var tableCRC = []struct {
	buf []byte
	crc uint16
}{
	{[]byte{0x00, 0x01, 0x21}, 0x08C6},
	{[]byte{0x80, 0x01, 0x21}, 0x82C5},
	{[]byte{0x80, 0x01, 0x20}, 0x02C0},
	{[]byte{0x00, 0x01, 0x0C}, 0x0828},
	{[]byte{0x00, 0x02, 0x06, 0x06}, 0x941B},
	{[]byte{0x80, 0x02, 0x06, 0x07}, 0x9421},
}

func TestCrc16(t *testing.T) {
	for _, v := range tableCRC {
		if crc := crc16Hash(v.buf); crc != v.crc {
			t.Errorf("crc16Hash failed, expected 0x%04X, got 0x%04X", v.crc, crc)
		}
	}
}

func TestCrc16Bytes(t *testing.T) {
	for _, v := range tableCRC {
		l := byte(v.crc & 0xFF)
		h := byte((v.crc >> 8) & 0xFF)
		b := crc16Bytes(v.buf)
		if b[0] != l || b[1] != h {
			t.Errorf("crc16Bytes failed, expected %04X, got %02X%02X", v.crc, h, l)
		}
	}
}
