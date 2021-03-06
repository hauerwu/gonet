package net_util

import(
	"fmt"
	"net"
)

type senderError struct{}

type Sender struct{
	Net string
	RemoteIP string
	RemotePort int
	Conn net.Conn
}

func (e *senderError) Error() string {return "sender error"}

func NewSender(net,ip string,port int) *Sender {
	s := Sender{net,ip,port,nil}

	return &s
}

func (s *Sender)Connect() error {
	addr := fmt.Sprintf("%s:%d",s.RemoteIP,s.RemotePort)
	conn,err := net.Dial(s.Net,addr)

	s.Conn = conn

	return err
}

func (s *Sender)SendData(data []byte) error {
	if s.Conn == nil{
		return &senderError{}
	}
	
	len,err := s.Conn.Write(data)

	fmt.Printf("send %d bytes\n",len)

	return err
}

func (s *Sender)Close() {
	s.Conn.Close()
}



















