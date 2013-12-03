package main

import(
	"fmt"
	"test/part"
)

type RedPart struct{ 
	part.Part
//	Id int
}

type OptionCommon struct {
    ShortName string "short option name"
    LongName  string "long option name"
}

func (obj RedPart)String() string{
	return obj.Part.String()
}

func main(){
	var p = RedPart{part.Part{1002,"test2"}}
	var q = OptionCommon{}

	_ = p
	fmt.Printf("%d\n",p.Id)
	fmt.Printf("LongName:%s\n",q.LongName)
}



















