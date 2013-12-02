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


func Handle(buff []byte) error{

	fmt.Println(buff)
	test := example.Test{}
	e := proto.Unmarshal(buff,&test)
	if e != nil{
		fmt.Printf("decode failed,due to:%s\n",e.Error())
		return e
	}

	fmt.Printf("%s %d\n",test.GetLabel(),test.GetType())
	fmt.Println(test)
	return nil
}

func main() {
	if len(os.Args) <  3  {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port\n", os.Args[0])
		os.Exit(1)
	}
	ip,port_str := os.Args[1],os.Args[2]

	port,err := strconv.Atoi(port_str)
	
	l := net_util.NewListener(ip,port,"tcp",1024)

	err = l.Listen()
	CheckError(err)

	l.Run(Handle)

	os.Exit(0)
}









