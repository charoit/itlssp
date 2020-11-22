package ssp

import (
	"bytes"
	"reflect"
	"testing"
)

var tablePack = []struct {
	buf []byte
	exp []byte
}{
	{[]byte{0x21}, []byte{xSTX, 0x00, 0x01, 0x21, 0xC6, 0x08}},
	{[]byte{0x20}, []byte{xSTX, 0x80, 0x01, 0x20, 0xC0, 0x02}},
	{[]byte{0x0C}, []byte{xSTX, 0x00, 0x01, 0x0C, 0x28, 0x08}},
	{[]byte{0x06, 0x06}, []byte{xSTX, 0x80, 0x02, 0x06, 0x06, 0x24, 0x14}},
	{[]byte{0x06, 0x07}, []byte{xSTX, 0x00, 0x02, 0x06, 0x07, 0x1E, 0x14}},
}

func TestUnitPack(t *testing.T) {

	u := &device{seq: 0x80}
	for _, v := range tablePack {
		if r := u.pack(v.buf); !reflect.DeepEqual(r, v.exp) {
			t.Errorf("pack failed, expected %X, got %X", v.exp, r)
		}
	}
}

func TestUnitUnpack(t *testing.T) {
	u := &device{seq: 0x80}
	for _, v := range tablePack {
		r, e := u.unpack(v.exp)
		if e != nil {
			t.Fatal(e)
		}
		if !reflect.DeepEqual(r, v.buf) {
			t.Errorf("unpack failed, expected %X, got %X", v.buf, r)
		}
	}
}

func TestUnitRead(t *testing.T) {

	u := &device{seq: 0x80}
	var table2 = []struct {
		src []byte
	}{
		{[]byte{0x00}},
		{[]byte{0x00, 0xAA}},
		{[]byte{0x00, 0xBB, 0xCC}},
	}

	for _, v := range table2 {
		b, e := u.read(bytes.NewBuffer(v.src))
		if e != nil || !reflect.DeepEqual(b, v.src) {
			t.Errorf("readPort bytes failed, expected %s, got %s", string(v.src), string(b))
		}
	}
}

func TestUnitCheckSTX(t *testing.T) {
	var table = []struct {
		buf []byte
		exp []byte
	}{
		{[]byte{0x00, 0x01, 0xF5, 0xA9}, []byte{0x00, 0x01, 0xF5, 0xA9}},
		{[]byte{0x00, 0x01, xSTX, 0xA9}, []byte{0x00, 0x01, xSTX, xSTX, 0xA9}},
		{[]byte{0x00, 0x01, 0xF5, xSTX}, []byte{0x00, 0x01, 0xF5, xSTX, xSTX}},
		{[]byte{xSTX, 0x01, 0xF5, 0xF5}, []byte{xSTX, xSTX, 0x01, 0xF5, 0xF5}},
		{[]byte{0x00, xSTX, xSTX, 0x09}, []byte{0x00, xSTX, xSTX, xSTX, xSTX, 0x09}},
		{[]byte{0x00, xSTX, xSTX, xSTX}, []byte{0x00, xSTX, xSTX, xSTX, xSTX, xSTX, xSTX}},
	}

	u := &device{seq: 0x80}
	for _, v := range table {
		if r := u.checkSTX(v.buf); !reflect.DeepEqual(r, v.exp) {
			t.Errorf("checkSTX failed, expected %v, got %v", v.exp, r)
		}
	}
}

func TestUnitGetSEQ(t *testing.T) {
	u := &device{seq: 0x80}
	if seq := u.getSEQ(); seq != 0 {
		t.Errorf("getSEQ failed, expected 0x00, got 0x%02X", seq)
	}
	if seq := u.getSEQ(); seq != 0x80 {
		t.Errorf("getSEQ failed, expected 0x80, got 0x%02X", seq)
	}
}

func TestUnitCheckResponse(t *testing.T) {
	//var tableResp = []struct{
	//	src byte
	//	exp SSPResponse
	//}{
	//	{0xAA, SspResponseOk},
	//	{byte(SspResponseOk), SspResponseOk},
	//	{byte(SspResponseFail), SspResponseFail},
	//	{byte(SspResponseCannotProcess), SspResponseCannotProcess},
	//}
	//
	//u := &device{seq:0x80}
	//for _, v := range tableResp {
	//	if e := u.checkResponse(v.src); !reflect.DeepEqual(r, v.exp) {
	//		t.Errorf("checkSTX failed, expected %v, got %v", v.exp, r)
	//	}
	//}
}
