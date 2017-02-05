package protocol

import (
	"bytes"
	"fmt"
	"io"
)

type Decoder struct {
	Reader io.Reader
}

func (d *Decoder) ReadBody() ([]byte, error) {
	size, err := ReadVarInt(d.Reader)
	if err != nil {
		return nil, err
	} else if size < 0 {
		return nil, errNegativeLength
	}

	buf := make([]byte, size)
	if _, err := io.ReadFull(d.Reader, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (d *Decoder) ReadSpecific(p Packet) error {
	buf, err := d.ReadBody()
	if err != nil {
		return err
	}

	r := bytes.NewReader(buf)
	if packetID, err := ReadVarInt(r); err != nil {
		return err
	} else if int(packetID) != p.id() {
		return fmt.Errorf("Unexpected packet type: read=0x%02X, expected=0x%02X", packetID, p.id())
	}

	return p.read(r)

}
