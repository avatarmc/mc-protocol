package protocol

import (
	"bytes"
	"compress/zlib"
	"io"
)

type PacketEncoder struct {
	Writer               io.Writer
	CompressionThreshold int
	ZlibWriter           *zlib.Writer
}

func (e *PacketEncoder) EncodePacket(packet Packet) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	// Contents of the packet (ID + Data)
	if err := WriteVarInt(buf, VarInt(packet.id())); err != nil {
		return nil, err
	}
	if err := packet.write(buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func (e *PacketEncoder) WritePacket(packet Packet) error {
	buf, err := e.EncodePacket(packet)
	if err != nil {
		return err
	}

	uncompessedSize := 0
	extra := 0
	// Only compress if compression is enabled and the packet is large enough
	if e.CompressionThreshold >= 0 && buf.Len() > e.CompressionThreshold {
		var err error
		nBuf := &bytes.Buffer{}
		if e.ZlibWriter == nil {
			e.ZlibWriter, _ = zlib.NewWriterLevel(nBuf, zlib.BestSpeed)
		} else {
			e.ZlibWriter.Reset(nBuf)
		}
		uncompessedSize = buf.Len()

		if _, err = buf.WriteTo(e.ZlibWriter); err != nil {
			return err
		}
		if err = e.ZlibWriter.Close(); err != nil {
			return err
		}
		buf = nBuf
	}

	// Account for the compression header if enabled
	if e.CompressionThreshold >= 0 {
		extra = varIntSize(VarInt(uncompessedSize))
	}

	// Write the length prefix followed by the buffer
	if err := WriteVarInt(e.Writer, VarInt(buf.Len()+extra)); err != nil {
		return err
	}

	// Write the uncompressed packet size
	if e.CompressionThreshold >= 0 {
		if err := WriteVarInt(e.Writer, VarInt(uncompessedSize)); err != nil {
			return err
		}
	}

	_, err = buf.WriteTo(e.Writer)
	return err
}
