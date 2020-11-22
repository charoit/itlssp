// Package ITL is s device drivers.
package itlssp

import "github.com/sigurn/crc16"

// crc16Hash calc CRC16 sum
func crc16Hash(buf []byte) uint16 {
	p := crc16.Params{
		Poly:   0x8005,
		Init:   0xFFFF,
		RefIn:  false,
		RefOut: false,
		XorOut: 0x0000,
		Check:  0xBB3D,
		Name:   "CRC-16",
	}
	return crc16.Checksum(buf, crc16.MakeTable(p))
}

// crc16Bytes return two bytes CRC16 {0xLO, 0xHI}
func crc16Bytes(buf []byte) []byte {
	crc := crc16Hash(buf)
	return []byte{
		byte(crc & 0xFF),        // lo byte
		byte((crc >> 8) & 0xFF), // hi byte
	}
}
