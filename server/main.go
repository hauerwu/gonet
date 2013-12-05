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

type temp struct {
		Id int32
		Name string
}

var dao ds.DAO

func InitDB(dao *ds.DAO) error {
	/*
	*dao = ds.New("mongodb")
	if dao == nil{
		fmt.Printf("can't new a %s dao\n","mongodb")
		os.Exit(1)
	}
	err := (*dao).Initialize("127.0.0.1::test:test:id")
	CheckError(err)
        */

	*dao = ds.New("mysql")
	if dao == nil{
		fmt.Printf("can't new a %s dao\n","mongodb")
		os.Exit(1)
	}
	err := (*dao).Initialize("127.0.0.1:3306:hauerwu:821010:test:test:id")
	CheckError(err)
	return err
}

func Select(msg interface{}) (interface{},error) {
	m,ok := msg.(example.Para)

	if !ok{
		return nil,errors.New("system error")
	}
	
	fmt.Println(m)
	
	result := make([]example.Para,3,3)
	err := dao.Select(m.GetId(),&result)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}

	for _,r := range result{
		fmt.Printf("%d %s\n",r.GetId(),r.GetName())
	}

	//dao.Finalize()

	return result,nil
}

func Update(msg interface{}) (interface{},error) {
	m,ok := msg.(example.Para)

	if !ok{
		return nil,errors.New("system error")
	}
	
	fmt.Println(m)

	err := dao.Update(m.GetId(),temp{Id:m.GetId(),Name:m.GetName()})
	CheckError(err)
	//dao.Finalize()

	return nil,err
}

func Insert(msg interface{}) (interface{},error) {
	m,ok := msg.(example.Para)

	if !ok{
		return nil,errors.New("system error")
	}
	
	fmt.Println(m)
	err := dao.Insert(temp{Id:m.GetId(),Name:m.GetName()})
	CheckError(err)
	//dao.Finalize()

	return nil,err
}

func Delete(msg interface{}) (interface{},error) {
	m,ok := msg.(example.Para)

	if !ok{
		return nil,errors.New("system error")
	}
	
	fmt.Println(m)

	err := dao.Delete(m.GetId())
	CheckError(err)

	//dao.Finalize()

	return nil,err
}

func CheckError(err error) {
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func Handle(buff []byte) error {

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
	gDispatcher.Add("update",Update)
	gDispatcher.Add("insert",Insert)
	gDispatcher.Add("delete",Delete)

	err = InitDB(&dao)
	CheckError(err)

	fmt.Println(dao)
	

	l.Run(Handle)
	dao.Finalize()

	os.Exit(0)
}









