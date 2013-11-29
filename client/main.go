package main

import(
	"fmt"
	"os"
	"strconv"
	"test/net_util"
	"proto/example"
)


func CheckError(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func main(){
	test := example.Test{}

	_ = test

	if len(os.Args) < 4{
		fmt.Printf("usage:%s ip port cmd\n",os.Args[0])
		os.Exit(1)
	}

	ip,port_str,cmd := os.Args[1],os.Args[2],os.Args[3]

	port,err := strconv.Atoi(port_str)
	CheckError(err)


	sender := net_util.NewSender("tcp",ip,port)

	err = sender.Connect()
	CheckError(err)

	err = sender.SendCmd(cmd)
	CheckError(err)

	sender.Close()

	os.Exit(0)

}
















