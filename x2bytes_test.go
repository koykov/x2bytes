package x2bytes

import (
	"bytes"
	"testing"
)

type stage struct {
	name string
	val  any
	dst  []byte
}

func TestToBytes(t *testing.T) {
	var (
		buf []byte
		err error

		p, s, eps         = []byte("foobar"), "foobar", []byte("foobar")
		b, i, u, f32, f64 = true, 127, uint(127), float32(3.141592), 3.141592
		i8, i16, i32, i64 = int8(127), int16(127), int32(127), int64(127)
		u8, u16, u32, u64 = uint8(127), uint16(127), uint32(127), uint64(127)
		eb, eiu, ef       = []byte("true"), []byte("127"), []byte("3.141592")
	)
	stages := []stage{
		{name: "bytes", val: p, dst: eps},
		{name: "*bytes", val: &p, dst: eps},
		{name: "string", val: s, dst: eps},
		{name: "*string", val: &s, dst: eps},
		{name: "bool", val: true, dst: eb},
		{name: "*bool", val: &b, dst: eb},
		{name: "int", val: i, dst: eiu},
		{name: "*int", val: &i, dst: eiu},
		{name: "int8", val: i8, dst: eiu},
		{name: "*int8", val: &i8, dst: eiu},
		{name: "int16", val: i16, dst: eiu},
		{name: "*int16", val: &i16, dst: eiu},
		{name: "int32", val: i32, dst: eiu},
		{name: "*int32", val: &i32, dst: eiu},
		{name: "int64", val: i64, dst: eiu},
		{name: "*int64", val: &i64, dst: eiu},
		{name: "uint", val: u, dst: eiu},
		{name: "*uint", val: &u, dst: eiu},
		{name: "uint8", val: u8, dst: eiu},
		{name: "*uint8", val: &u8, dst: eiu},
		{name: "uint16", val: u16, dst: eiu},
		{name: "*uint16", val: &u16, dst: eiu},
		{name: "uint32", val: u32, dst: eiu},
		{name: "*uint32", val: &u32, dst: eiu},
		{name: "uint64", val: u64, dst: eiu},
		{name: "*uint64", val: &u64, dst: eiu},
		{name: "float32", val: f32, dst: ef},
		{name: "*float32", val: &f32, dst: ef},
		{name: "float64", val: f64, dst: ef},
		{name: "*float64", val: &f64, dst: ef},
	}
	x2b := func(t *testing.T, st *stage, buf []byte) {
		if buf, err = ToBytesWR(buf, st.val); err != nil {
			t.Error(err)
		}
		if st.name == "float32" || st.name == "*float32" {
			buf = buf[:8]
		}
		if !bytes.Equal(buf, st.dst) {
			t.Errorf("%s failed: need %s, got %s", st.name, string(buf), string(st.dst))
		}
	}
	for _, st := range stages {
		t.Run(st.name, func(t *testing.T) { x2b(t, &st, buf) })
	}
}

func BenchmarkToBytes(b *testing.B) {
	var (
		buf []byte
		err error

		p, s, eps          = []byte("foobar"), "foobar", []byte("foobar")
		bl, i, u, f32, f64 = true, 127, uint(127), float32(3.141592), 3.141592
		i8, i16, i32, i64  = int8(127), int16(127), int32(127), int64(127)
		u8, u16, u32, u64  = uint8(127), uint16(127), uint32(127), uint64(127)
		eb, eiu, ef        = []byte("true"), []byte("127"), []byte("3.141592")
	)
	stages := []stage{
		{name: "bytes", val: p, dst: eps},
		{name: "*bytes", val: &p, dst: eps},
		{name: "string", val: s, dst: eps},
		{name: "*string", val: &s, dst: eps},
		{name: "bool", val: true, dst: eb},
		{name: "*bool", val: &bl, dst: eb},
		{name: "int", val: i, dst: eiu},
		{name: "*int", val: &i, dst: eiu},
		{name: "int8", val: i8, dst: eiu},
		{name: "*int8", val: &i8, dst: eiu},
		{name: "int16", val: i16, dst: eiu},
		{name: "*int16", val: &i16, dst: eiu},
		{name: "int32", val: i32, dst: eiu},
		{name: "*int32", val: &i32, dst: eiu},
		{name: "int64", val: i64, dst: eiu},
		{name: "*int64", val: &i64, dst: eiu},
		{name: "uint", val: u, dst: eiu},
		{name: "*uint", val: &u, dst: eiu},
		{name: "uint8", val: u8, dst: eiu},
		{name: "*uint8", val: &u8, dst: eiu},
		{name: "uint16", val: u16, dst: eiu},
		{name: "*uint16", val: &u16, dst: eiu},
		{name: "uint32", val: u32, dst: eiu},
		{name: "*uint32", val: &u32, dst: eiu},
		{name: "uint64", val: u64, dst: eiu},
		{name: "*uint64", val: &u64, dst: eiu},
		{name: "float32", val: f32, dst: ef},
		{name: "*float32", val: &f32, dst: ef},
		{name: "float64", val: f64, dst: ef},
		{name: "*float64", val: &f64, dst: ef},
	}
	x2b := func(b *testing.B, st *stage, buf []byte) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if buf, err = ToBytesWR(buf, st.val); err != nil {
				b.Error(err)
			}
			if st.name == "float32" || st.name == "*float32" {
				buf = buf[:8]
			}
			if !bytes.Equal(buf, st.dst) {
				b.Errorf("%s failed: need %s, got %s", st.name, string(buf), string(st.dst))
			}
		}
	}
	for _, st := range stages {
		b.Run(st.name, func(b *testing.B) { x2b(b, &st, buf) })
	}
}
