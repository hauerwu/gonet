/* ResolveIP
 */

package main

import (
	"test/net_util"
	"os"
	"fmt"
)

func CheckError(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func Handle(buff []byte) error{
	str := string(buff)

	fmt.Println(str)
	
	return nil
}

func main() {
	if len(os.Args) !=  1  {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	
	l := net_util.NewListener("127.0.0.1",8000,"tcp")

	err := l.Listen()
	CheckError(err)

	l.Run(Handle)

	os.Exit(0)
}
