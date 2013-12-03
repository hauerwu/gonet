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

	if len(os.Args) < 5{
		fmt.Printf("usage:%s ip port cmd id [name]\n",os.Args[0])
		os.Exit(1)
	}

	ip,port_str,cmd,id_str := os.Args[1],os.Args[2],os.Args[3],os.Args[4]
	var name string = ""

	if len(os.Args) >5{
		name = os.Args[5]
	}

	var port,id int
	var err error
	
	id,err = strconv.Atoi(id_str)
	CheckError(err)
	port,err = strconv.Atoi(port_str)
	CheckError(err)

	fmt.Printf("cmd:%s id:%d name:%s\n",cmd,id,name)

	sender := net_util.NewSender("tcp",ip,port)

	err = sender.Connect()
	CheckError(err)


	test := &example.Test {
		Head: &example.Head{
			Cmd: proto.String(cmd),
		},
		Para: &example.Para{
			Id:  proto.Int32(int32(id)),
			Name: proto.String(name),
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
















