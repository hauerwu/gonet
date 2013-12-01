package main

import(
	"fmt"
	"os"
	"strconv"
	"test/net_util"
	"proto/example"
	"code.google.com/p/goprotobuf/proto"
)


func CheckError(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func main(){

	if len(os.Args) < 4{
		fmt.Printf("usage:%s ip port cmd\n",os.Args[0])
		os.Exit(1)
	}

	ip,port_str,_ := os.Args[1],os.Args[2],os.Args[3]

	port,err := strconv.Atoi(port_str)
	CheckError(err)


	sender := net_util.NewSender("tcp",ip,port)

	err = sender.Connect()
	CheckError(err)

	
	test := &example.Test {
			Label: proto.String("hello"),
			Type:  proto.Int32(17),
			Optionalgroup: &example.Test_OptionalGroup {
				RequiredField: proto.String("good bye"),
			},
	}

	data,err := proto.Marshal(test)
	buff := net_util.Pack2Buff(data)
	fmt.Println(buff)
	err = sender.SendData(buff)
	CheckError(err)

	sender.Close()

	os.Exit(0)
}
















