// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nbt

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type TypeID int

const (
	TagEnd TypeID = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
)

var (
	ErrInvalidCompound = errors.New("invalid compound")
)

type Compound struct {
	Name  string
	Items map[string]interface{}
}

func NewCompound() *Compound {
	return &Compound{
		Items: make(map[string]interface{}),
	}
}

func (c *Compound) Serialize(w io.Writer) error {
	err := writeString(w, c.Name)
	if err != nil {
		return err
	}
	return c.serialize(w)
}

func (c *Compound) serialize(w io.Writer) error {
	for k, v := range c.Items {
		if err := writeByte(w, byte(getTypeID(v))); err != nil {
			return err
		}
		if err := writeString(w, k); err != nil {
			return err
		}
		if err := writeType(w, v); err != nil {
			return err
		}
	}
	return writeByte(w, byte(TagEnd))
}

func writeType(w io.Writer, v interface{}) error {
	switch v := v.(type) {
	case int8:
		return writeByte(w, byte(v))
	case int16:
		return binary.Write(w, binary.BigEndian, v)
	case int32:
		return binary.Write(w, binary.BigEndian, v)
	case int64:
		return binary.Write(w, binary.BigEndian, v)
	case float32:
		return binary.Write(w, binary.BigEndian, v)
	case float64:
		return binary.Write(w, binary.BigEndian, v)
	case []byte:
		if err := binary.Write(w, binary.BigEndian, int32(len(v))); err != nil {
			return err
		}
		_, err := w.Write(v)
		return err
	case string:
		return writeString(w, v)
	case *List:
		return v.serialize(w)
	case *Compound:
		return v.serialize(w)
	case []int32:
		if err := binary.Write(w, binary.BigEndian, int32(len(v))); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, v)
	}
	panic("unhandled type")
}

func getTypeID(i interface{}) TypeID {
	switch i.(type) {
	case int8:
		return TagByte
	case int16:
		return TagShort
	case int32:
		return TagInt
	case int64:
		return TagLong
	case float32:
		return TagFloat
	case float64:
		return TagDouble
	case []byte:
		return TagByteArray
	case string:
		return TagString
	case *List:
		return TagList
	case *Compound:
		return TagCompound
	case []int32:
		return TagIntArray
	}
	panic(fmt.Sprintf("invalid type %T", i))
}

func (c *Compound) Deserialize(r io.Reader) error {
	var err error
	c.Name, err = readString(r)
	if err != nil {
		return err
	}
	return c.deserialize(r)
}

func (c *Compound) deserialize(r io.Reader) error {
	for {
		id, err := readByte(r)
		if err != nil {
			return err
		}
		// End of compound
		if TypeID(id) == TagEnd {
			break
		}
		name, err := readString(r)
		if err != nil {
			return err
		}
		c.Items[name], err = readType(r, TypeID(id))
		if err != nil {
			return err
		}
	}
	return nil
}

type List struct {
	Type     TypeID
	Elements []interface{}
}

func (l *List) serialize(w io.Writer) error {
	if err := writeByte(w, byte(l.Type)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, int32(len(l.Elements))); err != nil {
		return err
	}
	for _, e := range l.Elements {
		if err := writeType(w, e); err != nil {
			return err
		}
	}
	return nil
}

func (l *List) deserialize(r io.Reader) error {
	t, err := readByte(r)
	if err != nil {
		return err
	}
	l.Type = TypeID(t)
	var le int32
	err = binary.Read(r, binary.BigEndian, &le)
	if err != nil {
		return err
	}
	if le < 0 {
		return errors.New("negative length for list")
	}
	l.Elements = make([]interface{}, le)
	for i := 0; i < int(le); i++ {
		l.Elements[i], err = readType(r, l.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

func readType(r io.Reader, id TypeID) (interface{}, error) {
	switch id {
	case TagByte:
		v, err := readByte(r)
		return int8(v), err
	case TagShort:
		var v int16
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TagInt:
		var v int32
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TagLong:
		var v int64
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TagFloat:
		var v float32
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TagDouble:
		var v float64
		err := binary.Read(r, binary.BigEndian, &v)
		return v, err
	case TagByteArray:
		var l int32
		err := binary.Read(r, binary.BigEndian, &l)
		if err != nil {
			return nil, err
		}
		if l < 0 {
			return nil, errors.New("negative length for byte array")
		}
		v := make([]byte, l)
		_, err = io.ReadFull(r, v)
		return v, err
	case TagString:
		return readString(r)
	case TagList:
		l := &List{}
		err := l.deserialize(r)
		return l, err
	case TagCompound:
		c := NewCompound()
		err := c.deserialize(r)
		return c, err
	case TagIntArray:
		var l int32
		err := binary.Read(r, binary.BigEndian, &l)
		if err != nil {
			return nil, err
		}
		if l < 0 {
			return nil, errors.New("negative length for int array")
		}
		v := make([]int32, l)
		err = binary.Read(r, binary.BigEndian, v)
		return v, err
	}
	return nil, fmt.Errorf("invalid type %d", id)
}

func writeByte(w io.Writer, b byte) error {
	if bw, ok := w.(io.ByteWriter); ok {
		return bw.WriteByte(b)
	}
	var buf [1]byte
	buf[0] = b
	_, err := w.Write(buf[:1])
	return err
}

func readByte(r io.Reader) (byte, error) {
	if br, ok := r.(io.ByteReader); ok {
		return br.ReadByte()
	}
	var buf [1]byte
	_, err := r.Read(buf[:1])
	return buf[0], err
}

func writeString(w io.Writer, str string) error {
	b := []byte(str)
	err := binary.Write(w, binary.BigEndian, int16(len(b)))
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func readString(r io.Reader) (string, error) {
	var l int16
	err := binary.Read(r, binary.BigEndian, &l)
	if err != nil {
		return "", nil
	}
	if l < 0 {
		return "", errors.New("negative length for string")
	}
	buf := make([]byte, int(l))
	_, err = io.ReadFull(r, buf)
	return string(buf), err
}
