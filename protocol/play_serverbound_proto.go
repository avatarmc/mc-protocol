// Generated by protocol_builder
// Do not edit

package protocol

import (
	"io"
	"io/ioutil"
	"math"
)

func (t *TeleportConfirm) id() int { return 0x00 }
func (t *TeleportConfirm) write(ww io.Writer) (err error) {
	if err = WriteVarInt(ww, t.TeleportID); err != nil {
		return
	}
	return
}
func (t *TeleportConfirm) read(rr io.Reader) (err error) {
	if t.TeleportID, err = ReadVarInt(rr); err != nil {
		return
	}
	return
}

func (t *TabComplete) id() int { return 0x01 }
func (t *TabComplete) write(ww io.Writer) (err error) {
	var tmp [8]byte
	if err = WriteString(ww, t.Text); err != nil {
		return
	}
	if err = WriteBool(ww, t.AssumeComman); err != nil {
		return
	}
	if err = WriteBool(ww, t.HasTarget); err != nil {
		return
	}
	if t.HasTarget == true {
		tmp[0] = byte(t.Target >> 56)
		tmp[1] = byte(t.Target >> 48)
		tmp[2] = byte(t.Target >> 40)
		tmp[3] = byte(t.Target >> 32)
		tmp[4] = byte(t.Target >> 24)
		tmp[5] = byte(t.Target >> 16)
		tmp[6] = byte(t.Target >> 8)
		tmp[7] = byte(t.Target >> 0)
		if _, err = ww.Write(tmp[:8]); err != nil {
			return
		}
	}
	return
}
func (t *TabComplete) read(rr io.Reader) (err error) {
	var tmp [8]byte
	if t.Text, err = ReadString(rr); err != nil {
		return
	}
	if t.AssumeComman, err = ReadBool(rr); err != nil {
		return
	}
	if t.HasTarget, err = ReadBool(rr); err != nil {
		return
	}
	if t.HasTarget == true {
		if _, err = rr.Read(tmp[:8]); err != nil {
			return
		}
		t.Target = (Position(tmp[7]) << 0) | (Position(tmp[6]) << 8) | (Position(tmp[5]) << 16) | (Position(tmp[4]) << 24) | (Position(tmp[3]) << 32) | (Position(tmp[2]) << 40) | (Position(tmp[1]) << 48) | (Position(tmp[0]) << 56)
	}
	return
}

func (c *ChatMessage) id() int { return 0x02 }
func (c *ChatMessage) write(ww io.Writer) (err error) {
	if err = WriteString(ww, c.Message); err != nil {
		return
	}
	return
}
func (c *ChatMessage) read(rr io.Reader) (err error) {
	if c.Message, err = ReadString(rr); err != nil {
		return
	}
	return
}

func (c *ClientStatus) id() int { return 0x03 }
func (c *ClientStatus) write(ww io.Writer) (err error) {
	if err = WriteVarInt(ww, c.ActionID); err != nil {
		return
	}
	return
}
func (c *ClientStatus) read(rr io.Reader) (err error) {
	if c.ActionID, err = ReadVarInt(rr); err != nil {
		return
	}
	return
}

func (c *ClientSettings) id() int { return 0x04 }
func (c *ClientSettings) write(ww io.Writer) (err error) {
	var tmp [1]byte
	if err = WriteString(ww, c.Locale); err != nil {
		return
	}
	tmp[0] = byte(c.ViewDistance >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	tmp[0] = byte(c.ChatMode >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	if err = WriteBool(ww, c.ChatColors); err != nil {
		return
	}
	tmp[0] = byte(c.DisplayedSkinParts >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	if err = WriteVarInt(ww, c.MainHand); err != nil {
		return
	}
	return
}
func (c *ClientSettings) read(rr io.Reader) (err error) {
	var tmp [1]byte
	if c.Locale, err = ReadString(rr); err != nil {
		return
	}
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.ViewDistance = (byte(tmp[0]) << 0)
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.ChatMode = (byte(tmp[0]) << 0)
	if c.ChatColors, err = ReadBool(rr); err != nil {
		return
	}
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.DisplayedSkinParts = (byte(tmp[0]) << 0)
	if c.MainHand, err = ReadVarInt(rr); err != nil {
		return
	}
	return
}

func (c *ConfirmTransactionServerbound) id() int { return 0x05 }
func (c *ConfirmTransactionServerbound) write(ww io.Writer) (err error) {
	var tmp [2]byte
	tmp[0] = byte(c.ID >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	tmp[0] = byte(c.ActionNumber >> 8)
	tmp[1] = byte(c.ActionNumber >> 0)
	if _, err = ww.Write(tmp[:2]); err != nil {
		return
	}
	if err = WriteBool(ww, c.Accepted); err != nil {
		return
	}
	return
}
func (c *ConfirmTransactionServerbound) read(rr io.Reader) (err error) {
	var tmp [2]byte
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.ID = (byte(tmp[0]) << 0)
	if _, err = rr.Read(tmp[:2]); err != nil {
		return
	}
	c.ActionNumber = int16((uint16(tmp[1]) << 0) | (uint16(tmp[0]) << 8))
	if c.Accepted, err = ReadBool(rr); err != nil {
		return
	}
	return
}

func (e *EnchantItem) id() int { return 0x06 }
func (e *EnchantItem) write(ww io.Writer) (err error) {
	var tmp [1]byte
	tmp[0] = byte(e.ID >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	tmp[0] = byte(e.Enchantment >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	return
}
func (e *EnchantItem) read(rr io.Reader) (err error) {
	var tmp [1]byte
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	e.ID = (byte(tmp[0]) << 0)
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	e.Enchantment = (byte(tmp[0]) << 0)
	return
}

func (c *ClickWindow) id() int { return 0x07 }
func (c *ClickWindow) write(ww io.Writer) (err error) {
	var tmp [2]byte
	tmp[0] = byte(c.ID >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	tmp[0] = byte(c.Slot >> 8)
	tmp[1] = byte(c.Slot >> 0)
	if _, err = ww.Write(tmp[:2]); err != nil {
		return
	}
	tmp[0] = byte(c.Button >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	tmp[0] = byte(c.ActionNumber >> 8)
	tmp[1] = byte(c.ActionNumber >> 0)
	if _, err = ww.Write(tmp[:2]); err != nil {
		return
	}
	tmp[0] = byte(c.Mode >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	if err = c.ClickedItem.Serialize(ww); err != nil {
		return
	}
	return
}
func (c *ClickWindow) read(rr io.Reader) (err error) {
	var tmp [2]byte
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.ID = (byte(tmp[0]) << 0)
	if _, err = rr.Read(tmp[:2]); err != nil {
		return
	}
	c.Slot = int16((uint16(tmp[1]) << 0) | (uint16(tmp[0]) << 8))
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.Button = (byte(tmp[0]) << 0)
	if _, err = rr.Read(tmp[:2]); err != nil {
		return
	}
	c.ActionNumber = int16((uint16(tmp[1]) << 0) | (uint16(tmp[0]) << 8))
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.Mode = (byte(tmp[0]) << 0)
	if err = c.ClickedItem.Deserialize(rr); err != nil {
		return
	}
	return
}

func (c *CloseWindow) id() int { return 0x08 }
func (c *CloseWindow) write(ww io.Writer) (err error) {
	var tmp [1]byte
	tmp[0] = byte(c.ID >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	return
}
func (c *CloseWindow) read(rr io.Reader) (err error) {
	var tmp [1]byte
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.ID = (byte(tmp[0]) << 0)
	return
}

func (p *PluginMessageServerbound) id() int { return 0x09 }
func (p *PluginMessageServerbound) write(ww io.Writer) (err error) {
	if err = WriteString(ww, p.Channel); err != nil {
		return
	}
	if _, err = ww.Write(p.Data); err != nil {
		return
	}
	return
}
func (p *PluginMessageServerbound) read(rr io.Reader) (err error) {
	if p.Channel, err = ReadString(rr); err != nil {
		return
	}
	if p.Data, err = ioutil.ReadAll(rr); err != nil {
		return
	}
	return
}

func (u *UseEntity) id() int { return 0x0A }
func (u *UseEntity) write(ww io.Writer) (err error) {
	var tmp [4]byte
	if err = WriteVarInt(ww, u.TargetID); err != nil {
		return
	}
	if err = WriteVarInt(ww, u.Type); err != nil {
		return
	}
	if u.Type == 2 {
		tmp0 := math.Float32bits(u.TargetX)
		tmp[0] = byte(tmp0 >> 24)
		tmp[1] = byte(tmp0 >> 16)
		tmp[2] = byte(tmp0 >> 8)
		tmp[3] = byte(tmp0 >> 0)
		if _, err = ww.Write(tmp[:4]); err != nil {
			return
		}
		tmp1 := math.Float32bits(u.TargetY)
		tmp[0] = byte(tmp1 >> 24)
		tmp[1] = byte(tmp1 >> 16)
		tmp[2] = byte(tmp1 >> 8)
		tmp[3] = byte(tmp1 >> 0)
		if _, err = ww.Write(tmp[:4]); err != nil {
			return
		}
		tmp2 := math.Float32bits(u.TargetZ)
		tmp[0] = byte(tmp2 >> 24)
		tmp[1] = byte(tmp2 >> 16)
		tmp[2] = byte(tmp2 >> 8)
		tmp[3] = byte(tmp2 >> 0)
		if _, err = ww.Write(tmp[:4]); err != nil {
			return
		}
	}
	if u.Type == 0 || u.Type == 2 {
		if err = WriteVarInt(ww, u.Hand); err != nil {
			return
		}
	}
	return
}
func (u *UseEntity) read(rr io.Reader) (err error) {
	var tmp [4]byte
	if u.TargetID, err = ReadVarInt(rr); err != nil {
		return
	}
	if u.Type, err = ReadVarInt(rr); err != nil {
		return
	}
	if u.Type == 2 {
		var tmp0 uint32
		if _, err = rr.Read(tmp[:4]); err != nil {
			return
		}
		tmp0 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
		u.TargetX = math.Float32frombits(tmp0)
		var tmp1 uint32
		if _, err = rr.Read(tmp[:4]); err != nil {
			return
		}
		tmp1 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
		u.TargetY = math.Float32frombits(tmp1)
		var tmp2 uint32
		if _, err = rr.Read(tmp[:4]); err != nil {
			return
		}
		tmp2 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
		u.TargetZ = math.Float32frombits(tmp2)
	}
	if u.Type == 0 || u.Type == 2 {
		if u.Hand, err = ReadVarInt(rr); err != nil {
			return
		}
	}
	return
}

func (k *KeepAliveServerbound) id() int { return 0x0B }
func (k *KeepAliveServerbound) write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp[0] = byte(k.ID >> 56)
	tmp[1] = byte(k.ID >> 48)
	tmp[2] = byte(k.ID >> 40)
	tmp[3] = byte(k.ID >> 32)
	tmp[4] = byte(k.ID >> 24)
	tmp[5] = byte(k.ID >> 16)
	tmp[6] = byte(k.ID >> 8)
	tmp[7] = byte(k.ID >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	return
}
func (k *KeepAliveServerbound) read(rr io.Reader) (err error) {
	var tmp [8]byte
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	k.ID = int64((uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56))
	return
}

func (p *Player) id() int { return 0x0C }
func (p *Player) write(ww io.Writer) (err error) {
	if err = WriteBool(ww, p.OnGround); err != nil {
		return
	}
	return
}
func (p *Player) read(rr io.Reader) (err error) {
	if p.OnGround, err = ReadBool(rr); err != nil {
		return
	}
	return
}

func (p *PlayerPosition) id() int { return 0x0D }
func (p *PlayerPosition) write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp0 := math.Float64bits(p.X)
	tmp[0] = byte(tmp0 >> 56)
	tmp[1] = byte(tmp0 >> 48)
	tmp[2] = byte(tmp0 >> 40)
	tmp[3] = byte(tmp0 >> 32)
	tmp[4] = byte(tmp0 >> 24)
	tmp[5] = byte(tmp0 >> 16)
	tmp[6] = byte(tmp0 >> 8)
	tmp[7] = byte(tmp0 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp1 := math.Float64bits(p.Y)
	tmp[0] = byte(tmp1 >> 56)
	tmp[1] = byte(tmp1 >> 48)
	tmp[2] = byte(tmp1 >> 40)
	tmp[3] = byte(tmp1 >> 32)
	tmp[4] = byte(tmp1 >> 24)
	tmp[5] = byte(tmp1 >> 16)
	tmp[6] = byte(tmp1 >> 8)
	tmp[7] = byte(tmp1 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp2 := math.Float64bits(p.Z)
	tmp[0] = byte(tmp2 >> 56)
	tmp[1] = byte(tmp2 >> 48)
	tmp[2] = byte(tmp2 >> 40)
	tmp[3] = byte(tmp2 >> 32)
	tmp[4] = byte(tmp2 >> 24)
	tmp[5] = byte(tmp2 >> 16)
	tmp[6] = byte(tmp2 >> 8)
	tmp[7] = byte(tmp2 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	if err = WriteBool(ww, p.OnGround); err != nil {
		return
	}
	return
}
func (p *PlayerPosition) read(rr io.Reader) (err error) {
	var tmp [8]byte
	var tmp0 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp0 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.X = math.Float64frombits(tmp0)
	var tmp1 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp1 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.Y = math.Float64frombits(tmp1)
	var tmp2 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp2 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.Z = math.Float64frombits(tmp2)
	if p.OnGround, err = ReadBool(rr); err != nil {
		return
	}
	return
}

func (p *PlayerPositionLook) id() int { return 0x0E }
func (p *PlayerPositionLook) write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp0 := math.Float64bits(p.X)
	tmp[0] = byte(tmp0 >> 56)
	tmp[1] = byte(tmp0 >> 48)
	tmp[2] = byte(tmp0 >> 40)
	tmp[3] = byte(tmp0 >> 32)
	tmp[4] = byte(tmp0 >> 24)
	tmp[5] = byte(tmp0 >> 16)
	tmp[6] = byte(tmp0 >> 8)
	tmp[7] = byte(tmp0 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp1 := math.Float64bits(p.Y)
	tmp[0] = byte(tmp1 >> 56)
	tmp[1] = byte(tmp1 >> 48)
	tmp[2] = byte(tmp1 >> 40)
	tmp[3] = byte(tmp1 >> 32)
	tmp[4] = byte(tmp1 >> 24)
	tmp[5] = byte(tmp1 >> 16)
	tmp[6] = byte(tmp1 >> 8)
	tmp[7] = byte(tmp1 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp2 := math.Float64bits(p.Z)
	tmp[0] = byte(tmp2 >> 56)
	tmp[1] = byte(tmp2 >> 48)
	tmp[2] = byte(tmp2 >> 40)
	tmp[3] = byte(tmp2 >> 32)
	tmp[4] = byte(tmp2 >> 24)
	tmp[5] = byte(tmp2 >> 16)
	tmp[6] = byte(tmp2 >> 8)
	tmp[7] = byte(tmp2 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp3 := math.Float32bits(p.Yaw)
	tmp[0] = byte(tmp3 >> 24)
	tmp[1] = byte(tmp3 >> 16)
	tmp[2] = byte(tmp3 >> 8)
	tmp[3] = byte(tmp3 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	tmp4 := math.Float32bits(p.Pitch)
	tmp[0] = byte(tmp4 >> 24)
	tmp[1] = byte(tmp4 >> 16)
	tmp[2] = byte(tmp4 >> 8)
	tmp[3] = byte(tmp4 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	if err = WriteBool(ww, p.OnGround); err != nil {
		return
	}
	return
}
func (p *PlayerPositionLook) read(rr io.Reader) (err error) {
	var tmp [8]byte
	var tmp0 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp0 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.X = math.Float64frombits(tmp0)
	var tmp1 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp1 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.Y = math.Float64frombits(tmp1)
	var tmp2 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp2 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.Z = math.Float64frombits(tmp2)
	var tmp3 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp3 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.Yaw = math.Float32frombits(tmp3)
	var tmp4 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp4 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.Pitch = math.Float32frombits(tmp4)
	if p.OnGround, err = ReadBool(rr); err != nil {
		return
	}
	return
}

func (p *PlayerLook) id() int { return 0x0F }
func (p *PlayerLook) write(ww io.Writer) (err error) {
	var tmp [4]byte
	tmp0 := math.Float32bits(p.Yaw)
	tmp[0] = byte(tmp0 >> 24)
	tmp[1] = byte(tmp0 >> 16)
	tmp[2] = byte(tmp0 >> 8)
	tmp[3] = byte(tmp0 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	tmp1 := math.Float32bits(p.Pitch)
	tmp[0] = byte(tmp1 >> 24)
	tmp[1] = byte(tmp1 >> 16)
	tmp[2] = byte(tmp1 >> 8)
	tmp[3] = byte(tmp1 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	if err = WriteBool(ww, p.OnGround); err != nil {
		return
	}
	return
}
func (p *PlayerLook) read(rr io.Reader) (err error) {
	var tmp [4]byte
	var tmp0 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp0 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.Yaw = math.Float32frombits(tmp0)
	var tmp1 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp1 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.Pitch = math.Float32frombits(tmp1)
	if p.OnGround, err = ReadBool(rr); err != nil {
		return
	}
	return
}

func (p *PlayerVehicleMove) id() int { return 0x10 }
func (p *PlayerVehicleMove) write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp0 := math.Float64bits(p.X)
	tmp[0] = byte(tmp0 >> 56)
	tmp[1] = byte(tmp0 >> 48)
	tmp[2] = byte(tmp0 >> 40)
	tmp[3] = byte(tmp0 >> 32)
	tmp[4] = byte(tmp0 >> 24)
	tmp[5] = byte(tmp0 >> 16)
	tmp[6] = byte(tmp0 >> 8)
	tmp[7] = byte(tmp0 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp1 := math.Float64bits(p.Y)
	tmp[0] = byte(tmp1 >> 56)
	tmp[1] = byte(tmp1 >> 48)
	tmp[2] = byte(tmp1 >> 40)
	tmp[3] = byte(tmp1 >> 32)
	tmp[4] = byte(tmp1 >> 24)
	tmp[5] = byte(tmp1 >> 16)
	tmp[6] = byte(tmp1 >> 8)
	tmp[7] = byte(tmp1 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp2 := math.Float64bits(p.Z)
	tmp[0] = byte(tmp2 >> 56)
	tmp[1] = byte(tmp2 >> 48)
	tmp[2] = byte(tmp2 >> 40)
	tmp[3] = byte(tmp2 >> 32)
	tmp[4] = byte(tmp2 >> 24)
	tmp[5] = byte(tmp2 >> 16)
	tmp[6] = byte(tmp2 >> 8)
	tmp[7] = byte(tmp2 >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp3 := math.Float32bits(p.Yaw)
	tmp[0] = byte(tmp3 >> 24)
	tmp[1] = byte(tmp3 >> 16)
	tmp[2] = byte(tmp3 >> 8)
	tmp[3] = byte(tmp3 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	tmp4 := math.Float32bits(p.Pitch)
	tmp[0] = byte(tmp4 >> 24)
	tmp[1] = byte(tmp4 >> 16)
	tmp[2] = byte(tmp4 >> 8)
	tmp[3] = byte(tmp4 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	return
}
func (p *PlayerVehicleMove) read(rr io.Reader) (err error) {
	var tmp [8]byte
	var tmp0 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp0 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.X = math.Float64frombits(tmp0)
	var tmp1 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp1 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.Y = math.Float64frombits(tmp1)
	var tmp2 uint64
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	tmp2 = (uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56)
	p.Z = math.Float64frombits(tmp2)
	var tmp3 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp3 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.Yaw = math.Float32frombits(tmp3)
	var tmp4 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp4 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.Pitch = math.Float32frombits(tmp4)
	return
}

func (s *SteerBoat) id() int { return 0x11 }
func (s *SteerBoat) write(ww io.Writer) (err error) {
	if err = WriteBool(ww, s.LeftPaddle); err != nil {
		return
	}
	if err = WriteBool(ww, s.RightPaddle); err != nil {
		return
	}
	return
}
func (s *SteerBoat) read(rr io.Reader) (err error) {
	if s.LeftPaddle, err = ReadBool(rr); err != nil {
		return
	}
	if s.RightPaddle, err = ReadBool(rr); err != nil {
		return
	}
	return
}

func (c *CraftRecipeRequest) id() int { return 0x12 }
func (c *CraftRecipeRequest) write(ww io.Writer) (err error) {
	var tmp [1]byte
	tmp[0] = byte(c.WindowID >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	if err = WriteVarInt(ww, c.RecipeID); err != nil {
		return
	}
	if err = WriteBool(ww, c.MakeAll); err != nil {
		return
	}
	return
}
func (c *CraftRecipeRequest) read(rr io.Reader) (err error) {
	var tmp [1]byte
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.WindowID = (byte(tmp[0]) << 0)
	if c.RecipeID, err = ReadVarInt(rr); err != nil {
		return
	}
	if c.MakeAll, err = ReadBool(rr); err != nil {
		return
	}
	return
}

func (c *ClientAbilities) id() int { return 0x13 }
func (c *ClientAbilities) write(ww io.Writer) (err error) {
	var tmp [4]byte
	tmp[0] = byte(c.Flags >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	tmp0 := math.Float32bits(c.FlyingSpeed)
	tmp[0] = byte(tmp0 >> 24)
	tmp[1] = byte(tmp0 >> 16)
	tmp[2] = byte(tmp0 >> 8)
	tmp[3] = byte(tmp0 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	tmp1 := math.Float32bits(c.WalkingSpeed)
	tmp[0] = byte(tmp1 >> 24)
	tmp[1] = byte(tmp1 >> 16)
	tmp[2] = byte(tmp1 >> 8)
	tmp[3] = byte(tmp1 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	return
}
func (c *ClientAbilities) read(rr io.Reader) (err error) {
	var tmp [4]byte
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	c.Flags = (byte(tmp[0]) << 0)
	var tmp0 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp0 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	c.FlyingSpeed = math.Float32frombits(tmp0)
	var tmp1 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp1 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	c.WalkingSpeed = math.Float32frombits(tmp1)
	return
}

func (p *PlayerDigging) id() int { return 0x14 }
func (p *PlayerDigging) write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp[0] = byte(p.Status >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	tmp[0] = byte(p.Location >> 56)
	tmp[1] = byte(p.Location >> 48)
	tmp[2] = byte(p.Location >> 40)
	tmp[3] = byte(p.Location >> 32)
	tmp[4] = byte(p.Location >> 24)
	tmp[5] = byte(p.Location >> 16)
	tmp[6] = byte(p.Location >> 8)
	tmp[7] = byte(p.Location >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	tmp[0] = byte(p.Face >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	return
}
func (p *PlayerDigging) read(rr io.Reader) (err error) {
	var tmp [8]byte
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	p.Status = (byte(tmp[0]) << 0)
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	p.Location = (Position(tmp[7]) << 0) | (Position(tmp[6]) << 8) | (Position(tmp[5]) << 16) | (Position(tmp[4]) << 24) | (Position(tmp[3]) << 32) | (Position(tmp[2]) << 40) | (Position(tmp[1]) << 48) | (Position(tmp[0]) << 56)
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	p.Face = (byte(tmp[0]) << 0)
	return
}

func (p *PlayerAction) id() int { return 0x15 }
func (p *PlayerAction) write(ww io.Writer) (err error) {
	if err = WriteVarInt(ww, p.EntityID); err != nil {
		return
	}
	if err = WriteVarInt(ww, p.ActionID); err != nil {
		return
	}
	if err = WriteVarInt(ww, p.JumpBoost); err != nil {
		return
	}
	return
}
func (p *PlayerAction) read(rr io.Reader) (err error) {
	if p.EntityID, err = ReadVarInt(rr); err != nil {
		return
	}
	if p.ActionID, err = ReadVarInt(rr); err != nil {
		return
	}
	if p.JumpBoost, err = ReadVarInt(rr); err != nil {
		return
	}
	return
}

func (s *SteerVehicle) id() int { return 0x16 }
func (s *SteerVehicle) write(ww io.Writer) (err error) {
	var tmp [4]byte
	tmp0 := math.Float32bits(s.Sideways)
	tmp[0] = byte(tmp0 >> 24)
	tmp[1] = byte(tmp0 >> 16)
	tmp[2] = byte(tmp0 >> 8)
	tmp[3] = byte(tmp0 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	tmp1 := math.Float32bits(s.Forward)
	tmp[0] = byte(tmp1 >> 24)
	tmp[1] = byte(tmp1 >> 16)
	tmp[2] = byte(tmp1 >> 8)
	tmp[3] = byte(tmp1 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	tmp[0] = byte(s.Flags >> 0)
	if _, err = ww.Write(tmp[:1]); err != nil {
		return
	}
	return
}
func (s *SteerVehicle) read(rr io.Reader) (err error) {
	var tmp [4]byte
	var tmp0 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp0 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	s.Sideways = math.Float32frombits(tmp0)
	var tmp1 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp1 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	s.Forward = math.Float32frombits(tmp1)
	if _, err = rr.Read(tmp[:1]); err != nil {
		return
	}
	s.Flags = (byte(tmp[0]) << 0)
	return
}

func (c *CraftingBookData) id() int { return 0x17 }
func (c *CraftingBookData) write(ww io.Writer) (err error) {
	var tmp [4]byte
	if err = WriteVarInt(ww, c.Type); err != nil {
		return
	}
	if c.Type == 0 {
		tmp0 := math.Float32bits(c.RecipeID)
		tmp[0] = byte(tmp0 >> 24)
		tmp[1] = byte(tmp0 >> 16)
		tmp[2] = byte(tmp0 >> 8)
		tmp[3] = byte(tmp0 >> 0)
		if _, err = ww.Write(tmp[:4]); err != nil {
			return
		}
	}
	if c.Type == 1 {
		if err = WriteBool(ww, c.CraftingBookOpen); err != nil {
			return
		}
		if err = WriteBool(ww, c.CraftingBookFilter); err != nil {
			return
		}
	}
	return
}
func (c *CraftingBookData) read(rr io.Reader) (err error) {
	var tmp [4]byte
	if c.Type, err = ReadVarInt(rr); err != nil {
		return
	}
	if c.Type == 0 {
		var tmp0 uint32
		if _, err = rr.Read(tmp[:4]); err != nil {
			return
		}
		tmp0 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
		c.RecipeID = math.Float32frombits(tmp0)
	}
	if c.Type == 1 {
		if c.CraftingBookOpen, err = ReadBool(rr); err != nil {
			return
		}
		if c.CraftingBookFilter, err = ReadBool(rr); err != nil {
			return
		}
	}
	return
}

func (r *ResourcePackStatus) id() int { return 0x18 }
func (r *ResourcePackStatus) write(ww io.Writer) (err error) {
	if err = WriteVarInt(ww, r.Result); err != nil {
		return
	}
	return
}
func (r *ResourcePackStatus) read(rr io.Reader) (err error) {
	if r.Result, err = ReadVarInt(rr); err != nil {
		return
	}
	return
}

func (a *AdvancementTab) id() int { return 0x19 }
func (a *AdvancementTab) write(ww io.Writer) (err error) {
	if err = WriteVarInt(ww, a.Action); err != nil {
		return
	}
	if a.Action == 0 {
		if err = WriteString(ww, a.TabID); err != nil {
			return
		}
	}
	return
}
func (a *AdvancementTab) read(rr io.Reader) (err error) {
	if a.Action, err = ReadVarInt(rr); err != nil {
		return
	}
	if a.Action == 0 {
		if a.TabID, err = ReadString(rr); err != nil {
			return
		}
	}
	return
}

func (h *HeldItemChange) id() int { return 0x1A }
func (h *HeldItemChange) write(ww io.Writer) (err error) {
	var tmp [2]byte
	tmp[0] = byte(h.Slot >> 8)
	tmp[1] = byte(h.Slot >> 0)
	if _, err = ww.Write(tmp[:2]); err != nil {
		return
	}
	return
}
func (h *HeldItemChange) read(rr io.Reader) (err error) {
	var tmp [2]byte
	if _, err = rr.Read(tmp[:2]); err != nil {
		return
	}
	h.Slot = int16((uint16(tmp[1]) << 0) | (uint16(tmp[0]) << 8))
	return
}

func (c *CreativeInventoryAction) id() int { return 0x1B }
func (c *CreativeInventoryAction) write(ww io.Writer) (err error) {
	var tmp [2]byte
	tmp[0] = byte(c.Slot >> 8)
	tmp[1] = byte(c.Slot >> 0)
	if _, err = ww.Write(tmp[:2]); err != nil {
		return
	}
	if err = c.ClickedItem.Serialize(ww); err != nil {
		return
	}
	return
}
func (c *CreativeInventoryAction) read(rr io.Reader) (err error) {
	var tmp [2]byte
	if _, err = rr.Read(tmp[:2]); err != nil {
		return
	}
	c.Slot = int16((uint16(tmp[1]) << 0) | (uint16(tmp[0]) << 8))
	if err = c.ClickedItem.Deserialize(rr); err != nil {
		return
	}
	return
}

func (s *SetSign) id() int { return 0x1C }
func (s *SetSign) write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp[0] = byte(s.Location >> 56)
	tmp[1] = byte(s.Location >> 48)
	tmp[2] = byte(s.Location >> 40)
	tmp[3] = byte(s.Location >> 32)
	tmp[4] = byte(s.Location >> 24)
	tmp[5] = byte(s.Location >> 16)
	tmp[6] = byte(s.Location >> 8)
	tmp[7] = byte(s.Location >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	if err = WriteString(ww, s.Line1); err != nil {
		return
	}
	if err = WriteString(ww, s.Line2); err != nil {
		return
	}
	if err = WriteString(ww, s.Line3); err != nil {
		return
	}
	if err = WriteString(ww, s.Line4); err != nil {
		return
	}
	return
}
func (s *SetSign) read(rr io.Reader) (err error) {
	var tmp [8]byte
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	s.Location = (Position(tmp[7]) << 0) | (Position(tmp[6]) << 8) | (Position(tmp[5]) << 16) | (Position(tmp[4]) << 24) | (Position(tmp[3]) << 32) | (Position(tmp[2]) << 40) | (Position(tmp[1]) << 48) | (Position(tmp[0]) << 56)
	if s.Line1, err = ReadString(rr); err != nil {
		return
	}
	if s.Line2, err = ReadString(rr); err != nil {
		return
	}
	if s.Line3, err = ReadString(rr); err != nil {
		return
	}
	if s.Line4, err = ReadString(rr); err != nil {
		return
	}
	return
}

func (a *ArmSwing) id() int { return 0x1D }
func (a *ArmSwing) write(ww io.Writer) (err error) {
	if err = WriteVarInt(ww, a.Hand); err != nil {
		return
	}
	return
}
func (a *ArmSwing) read(rr io.Reader) (err error) {
	if a.Hand, err = ReadVarInt(rr); err != nil {
		return
	}
	return
}

func (s *SpectateTeleport) id() int { return 0x1E }
func (s *SpectateTeleport) write(ww io.Writer) (err error) {
	if err = s.Target.Serialize(ww); err != nil {
		return
	}
	return
}
func (s *SpectateTeleport) read(rr io.Reader) (err error) {
	if err = s.Target.Deserialize(rr); err != nil {
		return
	}
	return
}

func (p *PlayerBlockPlacement) id() int { return 0x1F }
func (p *PlayerBlockPlacement) write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp[0] = byte(p.Location >> 56)
	tmp[1] = byte(p.Location >> 48)
	tmp[2] = byte(p.Location >> 40)
	tmp[3] = byte(p.Location >> 32)
	tmp[4] = byte(p.Location >> 24)
	tmp[5] = byte(p.Location >> 16)
	tmp[6] = byte(p.Location >> 8)
	tmp[7] = byte(p.Location >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	if err = WriteVarInt(ww, p.Face); err != nil {
		return
	}
	if err = WriteVarInt(ww, p.Hand); err != nil {
		return
	}
	tmp0 := math.Float32bits(p.CursorX)
	tmp[0] = byte(tmp0 >> 24)
	tmp[1] = byte(tmp0 >> 16)
	tmp[2] = byte(tmp0 >> 8)
	tmp[3] = byte(tmp0 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	tmp1 := math.Float32bits(p.CursorY)
	tmp[0] = byte(tmp1 >> 24)
	tmp[1] = byte(tmp1 >> 16)
	tmp[2] = byte(tmp1 >> 8)
	tmp[3] = byte(tmp1 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	tmp2 := math.Float32bits(p.CursorZ)
	tmp[0] = byte(tmp2 >> 24)
	tmp[1] = byte(tmp2 >> 16)
	tmp[2] = byte(tmp2 >> 8)
	tmp[3] = byte(tmp2 >> 0)
	if _, err = ww.Write(tmp[:4]); err != nil {
		return
	}
	return
}
func (p *PlayerBlockPlacement) read(rr io.Reader) (err error) {
	var tmp [8]byte
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	p.Location = (Position(tmp[7]) << 0) | (Position(tmp[6]) << 8) | (Position(tmp[5]) << 16) | (Position(tmp[4]) << 24) | (Position(tmp[3]) << 32) | (Position(tmp[2]) << 40) | (Position(tmp[1]) << 48) | (Position(tmp[0]) << 56)
	if p.Face, err = ReadVarInt(rr); err != nil {
		return
	}
	if p.Hand, err = ReadVarInt(rr); err != nil {
		return
	}
	var tmp0 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp0 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.CursorX = math.Float32frombits(tmp0)
	var tmp1 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp1 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.CursorY = math.Float32frombits(tmp1)
	var tmp2 uint32
	if _, err = rr.Read(tmp[:4]); err != nil {
		return
	}
	tmp2 = (uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24)
	p.CursorZ = math.Float32frombits(tmp2)
	return
}

func (u *UseItem) id() int { return 0x20 }
func (u *UseItem) write(ww io.Writer) (err error) {
	if err = WriteVarInt(ww, u.Hand); err != nil {
		return
	}
	return
}
func (u *UseItem) read(rr io.Reader) (err error) {
	if u.Hand, err = ReadVarInt(rr); err != nil {
		return
	}
	return
}

func init() {
	packetCreator[Play][serverbound][0x00] = func() Packet { return &TeleportConfirm{} }
	packetCreator[Play][serverbound][0x01] = func() Packet { return &TabComplete{} }
	packetCreator[Play][serverbound][0x02] = func() Packet { return &ChatMessage{} }
	packetCreator[Play][serverbound][0x03] = func() Packet { return &ClientStatus{} }
	packetCreator[Play][serverbound][0x04] = func() Packet { return &ClientSettings{} }
	packetCreator[Play][serverbound][0x05] = func() Packet { return &ConfirmTransactionServerbound{} }
	packetCreator[Play][serverbound][0x06] = func() Packet { return &EnchantItem{} }
	packetCreator[Play][serverbound][0x07] = func() Packet { return &ClickWindow{} }
	packetCreator[Play][serverbound][0x08] = func() Packet { return &CloseWindow{} }
	packetCreator[Play][serverbound][0x09] = func() Packet { return &PluginMessageServerbound{} }
	packetCreator[Play][serverbound][0x0A] = func() Packet { return &UseEntity{} }
	packetCreator[Play][serverbound][0x0B] = func() Packet { return &KeepAliveServerbound{} }
	packetCreator[Play][serverbound][0x0C] = func() Packet { return &Player{} }
	packetCreator[Play][serverbound][0x0D] = func() Packet { return &PlayerPosition{} }
	packetCreator[Play][serverbound][0x0E] = func() Packet { return &PlayerPositionLook{} }
	packetCreator[Play][serverbound][0x0F] = func() Packet { return &PlayerLook{} }
	packetCreator[Play][serverbound][0x10] = func() Packet { return &PlayerVehicleMove{} }
	packetCreator[Play][serverbound][0x11] = func() Packet { return &SteerBoat{} }
	packetCreator[Play][serverbound][0x12] = func() Packet { return &CraftRecipeRequest{} }
	packetCreator[Play][serverbound][0x13] = func() Packet { return &ClientAbilities{} }
	packetCreator[Play][serverbound][0x14] = func() Packet { return &PlayerDigging{} }
	packetCreator[Play][serverbound][0x15] = func() Packet { return &PlayerAction{} }
	packetCreator[Play][serverbound][0x16] = func() Packet { return &SteerVehicle{} }
	packetCreator[Play][serverbound][0x17] = func() Packet { return &CraftingBookData{} }
	packetCreator[Play][serverbound][0x18] = func() Packet { return &ResourcePackStatus{} }
	packetCreator[Play][serverbound][0x19] = func() Packet { return &AdvancementTab{} }
	packetCreator[Play][serverbound][0x1A] = func() Packet { return &HeldItemChange{} }
	packetCreator[Play][serverbound][0x1B] = func() Packet { return &CreativeInventoryAction{} }
	packetCreator[Play][serverbound][0x1C] = func() Packet { return &SetSign{} }
	packetCreator[Play][serverbound][0x1D] = func() Packet { return &ArmSwing{} }
	packetCreator[Play][serverbound][0x1E] = func() Packet { return &SpectateTeleport{} }
	packetCreator[Play][serverbound][0x1F] = func() Packet { return &PlayerBlockPlacement{} }
	packetCreator[Play][serverbound][0x20] = func() Packet { return &UseItem{} }
}
