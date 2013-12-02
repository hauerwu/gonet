package net_util

import(
	"errors"
	"encoding/binary"
	"net"
	"fmt"
	"io"
)

const HeadLen = 4

type Netbuffer struct{
	cap uint32
	begin uint32
	end uint32
	buffer []byte
}

type Listener struct{
	IP string
	Port int
	Net string
	Ls net.Listener
}

func Pack2Buff(in []byte) []byte{
	var l = uint32(len(in))
	var cap = l + HeadLen
	out := make([]byte,cap,cap)
	binary.LittleEndian.PutUint32(out,l)
	copy(out[HeadLen:],in)

	return out
}

func NewBuffer(cap uint32) *Netbuffer {
	if cap <= HeadLen{
		return nil
	}

	b := &Netbuffer{cap,0,0,nil}
	b.buffer = make([]byte,cap,cap)

	return b
}

func NewListener(ip string,port int,net string,buffcap uint32) *Listener{
	l := Listener{ip,port,net,nil}
	
	return &l
}

func (l *Listener)Listen() error {
	addr := fmt.Sprintf("%s:%d",l.IP,l.Port)
	ls,err := net.Listen(l.Net,addr)

	l.Ls = ls
	
	if err != nil{
		return err
	}
	return nil
}

func (l *Listener)Run(handle func([]byte) error) {
	for{
		conn,err := l.Ls.Accept()
		if err != nil{
			continue
		}

		fmt.Printf("%s connected\n",conn.RemoteAddr())

		go func(net.Conn){
			buff := NewBuffer(1024)
			temp := make([]byte,1024)
			for{
				m, e := conn.Read(temp)
				if e == io.EOF {
					break
				}
				if e != nil {
					break
				}

				e = buff.PutData(temp[:m])
				if e != nil{
					break
				}
			}
			
			for{
				l := uint32(1024)
				e := buff.GetData(temp,&l)
				if e != nil{
					break
				}

				handle(temp[0:l])
			}
			conn.Close()

		}(conn)
	}
}

func (n *Netbuffer)Reset(){
	n.begin = 0
	n.end = 0
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
		n.end += l
	}else{
		rl := n.GetMaxCap() - n.end
		if rl >= l{
			copy(n.buffer[n.end:],data)
			n.end += l
		}else{
			copy(n.buffer[n.end:],data[:rl])
			copy(n.buffer[0:],data[rl:])
			n.end = l - rl
		}
	}

	return nil
}

func (n *Netbuffer)GetDataLen(data []byte,l uint32) error{
	if n.GetLen() < l{
		return errors.New("buffer not ready")
	}

	if n.end > n.begin{
		copy(data[:l],n.buffer[n.begin:n.begin+l])
		n.begin += l
	}else{
		rl := n.GetMaxCap() - n.begin
		if rl >= l{
			copy(data,n.buffer[n.begin:n.begin+l])
			n.begin += l
		}else{
			copy(data,n.buffer[n.begin:n.begin+rl])
			copy(data[rl:],n.buffer[:l-rl])
			n.begin = l - rl
		}
	}

	return nil
}

func (n *Netbuffer)GetData(data []byte,l *uint32) error{
	if n.GetLen() < HeadLen{
		return errors.New("buffer not ready")
	}

	len_buffer := make([]byte,HeadLen,HeadLen)
	err := n.GetDataLen(len_buffer,HeadLen)
	if err != nil {
		return err
	}

	*l = binary.LittleEndian.Uint32(len_buffer)

	err = n.GetDataLen(data,*l)
	if err != nil {
		return err
	}

	return nil
}

func (n *Netbuffer)GetLen()uint32{
	if n.end >= n.begin{
		return n.end - n.begin
	}else{
		return n.cap - n.begin + n.end
	}
}





















