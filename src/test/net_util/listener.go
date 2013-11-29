package net_util

import(
	"net"
	"fmt"
	"io/ioutil"
)


type Listener struct{
	IP string
	Port int
	Net string
	Ls net.Listener
}

func NewListener(ip string,port int,net string) *Listener{
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
			//buff := make([]byte,1024)
			res,readerr := ioutil.ReadAll(conn)
			_ = readerr
			conn.Close()

			handle(res)
		}(conn)
	}
}



















