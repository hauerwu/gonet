/* ResolveIP
 */

package main

import (
	"test/net_util"
	"os"
	"fmt"
	"strconv"
	"proto/example"
	"code.google.com/p/goprotobuf/proto"
)

func CheckError(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func ParseBuff(buff []byte) (int32,[]byte){
	l := int32(buff[0])
	if l != int32(len(buff) - 1){
		return -1,buff
	}
	
	if l == 0{
		return 0,buff
	}else{
		return l,buff[1:]
	}
}

func Handle(buff []byte) error{
	l,b := ParseBuff(buff)
	if l <= 0{
		return nil
	}

	fmt.Println(b)
	test := example.Test{}
	proto.Unmarshal(b,&test)

	fmt.Println(test)
	//fmt.Printf("%s %d\n",test.GetLabel(),test.GetType())
	return nil
}

func main() {
	if len(os.Args) <  3  {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port\n", os.Args[0])
		os.Exit(1)
	}
	ip,port_str := os.Args[1],os.Args[2]

	port,err := strconv.Atoi(port_str)
	
	l := net_util.NewListener(ip,port,"tcp")

	err = l.Listen()
	CheckError(err)

	l.Run(Handle)

	os.Exit(0)
}









