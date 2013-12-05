package main

import (
	"fmt"
	"reflect"
	_ "database/sql"
)

type Test struct {
	Id int
	Name string
}

func GetSqlResult(d interface{}) {
	t := reflect.ValueOf(d)
	e := t.Elem()

	e = e.Slice(0,e.Len())
	
	i := 0
	for i < e.Len(){
		it := e.Index(i)
		
		j := 0 
		for j < it.NumField(){
			switch it.Field(j).Kind(){
			case reflect.Int:
				it.Field(j).SetInt(int64(i))
			case reflect.String:
				it.Field(j).SetString("test")
			default:
			}
			j++
		}
		i++
	}
}

func main() {
	s := make([]Test,5,10)
	GetSqlResult(&s)

	for _,e := range s{
		fmt.Println(e)
	}
}

















