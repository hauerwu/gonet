/* ResolveIP
 */

package main

import (
	"test/net_util"
	"test/dispatcher"
	"test/ds"
	"os"
	"fmt"
	"strconv"
	"errors"
	"proto/example"
	"code.google.com/p/goprotobuf/proto"
)

func Select(msg interface{}) (interface{},error){
	m,ok := msg.(example.Para)

	if !ok{
		return nil,errors.New("system error")
	}
	
	fmt.Println(m)

	dao := ds.New("mongodb")
	dao.Initialize("127.0.0.1::test:test:id")
	
	result := make([]example.Para,1,1)
	err := dao.Select(m.GetId(),result)
	if err != nil{
		return nil,err
	}

	fmt.Println(result)

	return result,nil
}

func CheckError(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func Handle(buff []byte) error{

	//fmt.Println(buff)
	test := example.Test{}
	e := proto.Unmarshal(buff,&test)
	if e != nil{
		fmt.Printf("decode failed,due to:%s\n",e.Error())
		return e
	}

	d := dispatcher.Instance()
	f := d.Select(test.GetHead().GetCmd())

	if f == nil{
		fmt.Printf("invalid cmd %s\n",test.GetHead().GetCmd())
	}else{
		f(*test.GetPara())
	}
	

	//fmt.Printf("%s %d %s\n",test.GetCmd(),test.GetId(),test.GetName())
	//fmt.Println(test)
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

	var gDispatcher = dispatcher.Instance()
	gDispatcher.Add("select",Select)

	l.Run(Handle)

	os.Exit(0)
}









