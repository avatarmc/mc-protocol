// Generated by protocol_builder
// Do not edit

package protocol

import (
	"io"
)

func (s *StatusRequest) id() int { return 0 }
func (s *StatusRequest) write(ww io.Writer) (err error) {
	return
}
func (s *StatusRequest) read(rr io.Reader) (err error) {
	return
}

func (s *StatusPing) id() int { return 1 }
func (s *StatusPing) write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp[0] = byte(s.Time >> 56)
	tmp[1] = byte(s.Time >> 48)
	tmp[2] = byte(s.Time >> 40)
	tmp[3] = byte(s.Time >> 32)
	tmp[4] = byte(s.Time >> 24)
	tmp[5] = byte(s.Time >> 16)
	tmp[6] = byte(s.Time >> 8)
	tmp[7] = byte(s.Time >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	return
}
func (s *StatusPing) read(rr io.Reader) (err error) {
	var tmp [8]byte
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	s.Time = int64((uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56))
	return
}

func init() {
	packetCreator[Status][serverbound][0] = func() Packet { return &StatusRequest{} }
	packetCreator[Status][serverbound][1] = func() Packet { return &StatusPing{} }
}
