package ssp

import (
	"bytes"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	var table1 = []struct {
		src string
	}{
		{"test1"},
		{"test2"},
		{"test3"},
	}

	for _, v := range table1 {
		b, e := readPort(bytes.NewBufferString(v.src))
		if e != nil || !reflect.DeepEqual(b, []byte(v.src)) {
			t.Errorf("readPort string failed, expected %s, got %s", v.src, string(b))
		}
	}

	var table2 = []struct {
		src []byte
	}{
		{[]byte{0x00}},
		{[]byte{0x00, 0xAA}},
		{[]byte{0x00, 0xBB, 0xCC}},
	}

	for _, v := range table2 {
		b, e := readPort(bytes.NewBuffer(v.src))
		if e != nil || !reflect.DeepEqual(b, v.src) {
			t.Errorf("readPort bytes failed, expected %s, got %s", string(v.src), string(b))
		}
	}
}
