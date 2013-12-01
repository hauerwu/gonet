package net_util

import(
	"errors"
	"encoding/binary"
)

const HeadLen = 4

type Netbuffer struct{
	cap uint32
	begin uint32
	end uint32
	buffer []byte
}

func Pack2Buff(in []byte) []byte{
	var l = uint32(len(in))
	var cap = l + HeadLen
	out := make([]byte,cap,cap)
	binary.LittleEndian.PutUint32(out,l)
	copy(out[HeadLen:],in)

	return out
}

func New(cap uint32) *Netbuffer {
	if cap <= HeadLen{
		return nil
	}

	b := &Netbuffer{cap,0,0,nil}
	b.buffer = make([]byte,0,b.cap)

	return b
}

func (n *Netbuffer)GetMaxCap() uint32{
	return n.cap
}

func (n *Netbuffer)GetCap() uint32{
	return n.GetMaxCap() - n.GetLen() - 1
}

func (n *Netbuffer)PutData(data []byte) error{
	l := uint32(len(data))
	if l > n.GetCap(){
		return errors.New("out of cap!")
	}

	if n.end < n.begin{
		copy(n.buffer[n.end:],data)
	}else{
		rl := n.GetMaxCap() - n.end
		if rl >= l{
			copy(n.buffer[n.end:],data)
		}else{
			copy(n.buffer[n.end:],data[:rl])
			copy(n.buffer[0:],data[rl:])
		}
	}

	return nil
}

func (n *Netbuffer)GetData(data []byte,l *uint32) error{
	return nil
}

func (n *Netbuffer)GetLen()uint32{
	if n.end >= n.begin{
		return n.end - n.begin
	}else{
		return n.cap - n.begin + n.end
	}
}





















