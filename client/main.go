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

	test := example.Test{}
	test.Label = new(string)
	*test.Label = "test"
	test.Type = new(int32)
	*test.Type = 6
	test.Reps = []int64{1,2,3,4}
	test.Optionalgroup = &example.Test_OptionalGroup{}
	test.Optionalgroup.RequiredField = new(string)
	*test.Optionalgroup.RequiredField = "opt"

	data,err := proto.Marshal(&test)
	buff := make([]byte,len(data)+1,len(data)+1)

	buff[0] = byte(len(data))
	copy(buff[1:],data)
	
	err = sender.SendData(buff)
	CheckError(err)

	sender.Close()

	os.Exit(0)

}
















