package mysql

import (
	"errors"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"reflect"
	"test/common/struct_helper"
)

type DAO struct {
	host  string
	port  string
	user  string
	pass  string
	db    string
	table string
	key   string

	conn *sql.DB
}

func ToUpdate(t reflect.StructField,v reflect.Value,sql *string,ex string,isfirst bool){
	name := t.Name
	if name == ex{
		return
	}

	switch {
	case t.Type.Kind() == reflect.String:
			if isfirst{
				*sql = fmt.Sprintf("%s %s='%s'",*sql,name,v.String())
			}else{
				*sql = fmt.Sprintf("%s ,%s='%s'",*sql,name,v.String())
			}
	case t.Type.Kind() >= reflect.Int && t.Type.Kind() <= reflect.Uint64: 
		if isfirst{
			*sql = fmt.Sprintf("%s %s=%d",*sql,name,v.Int())
		}else{
			*sql = fmt.Sprintf("%s ,%s=%d",*sql,name,v.Int())
		}
	default:
	}
}

func ToInsert(t reflect.StructField,v reflect.Value,sql *string,ex string,isfirst bool){
	
	switch {
	case t.Type.Kind() ==  reflect.String:
			if isfirst{
				*sql = fmt.Sprintf("%s'%s'",*sql,v.String())
			}else{
				*sql = fmt.Sprintf("%s,'%s'",*sql,v.String())
			}
	case t.Type.Kind() >= reflect.Int && t.Type.Kind() <= reflect.Uint64: 	
			if isfirst{
				*sql = fmt.Sprintf("%s%d",*sql,v.Int())
			}else{
				*sql = fmt.Sprintf("%s,%d",*sql,v.Int())
			}
	default:
	}
}

func New() *DAO {
	return new(DAO)
}

func (d *DAO) connect() error {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",d.user,d.pass,d.host,d.port,d.db)
	var err error
	
	d.conn,err = sql.Open("mysql",url)

	return err
}

func (d *DAO) Initialize(url string) error {
	strs := strings.Split(url, ":")
	if len(strs) < 7 {
		fmt.Printf("invalid url [%s]\n", url)
		return errors.New("can't parse ulr,try like this: 127.0.0.1:80:user:pass:test:peple:key")
	}

	d.host = strs[0]
	d.port = strs[1]
	d.user = strs[2]
	d.pass = strs[3]
	d.db = strs[4]
	d.table = strs[5]
	d.key = strs[6]

	err := d.connect()

	return err
}

func (d *DAO) Finalize() {
}

func (d DAO) Select(key interface{}, result interface{}) error {
	if d.conn == nil{
		return errors.New("no connection to mysql!")
	}
	
	var err error
	var rows *sql.Rows
	sql := fmt.Sprintf("select * from %s where id = %d",d.table,key)
	rows,err = d.conn.Query(sql)
	if err != nil{
		fmt.Println(err)
		return err
	}

	t := reflect.ValueOf(result)
	e := t.Elem()

	e = e.Slice(0,e.Len())
	i := 0
	for rows.Next() && i < e.Len(){
		j := 0
		s := make([]interface{},0,10)
		for j < e.Index(i).NumField(){
			s = append(s,e.Index(i).Field(j).Addr().Interface())
			j++
			//fmt.Println(j)
		}
		err = rows.Scan(s...)
		i++
	}
	
	return err
}

func (d DAO) Update(key interface{}, value interface{}) error {
	if d.conn == nil{
		return errors.New("no connection to mysql!")
	}

	var sql string = fmt.Sprintf("update %s set ",d.table)

	struct_helper.ParseSimpleStruct(value,ToUpdate,&sql,d.key)
	
	sql = fmt.Sprintf("%s where %s = %d",sql,d.key,key)

	fmt.Println(sql)

	sql = fmt.Sprintf(sql)

	_,err := d.conn.Exec(sql)

	return err
}

func (d DAO) Delete(key interface{}) error {
	if d.conn == nil{
		return errors.New("no connection to mysql!")
	}

	sql := fmt.Sprintf("delete from %s where %s = %d",d.table,d.key,key)
	fmt.Println(sql)
	_,err := d.conn.Exec(sql)
	
	return err
}

func (d DAO) Insert(value interface{}) error {
	if d.conn == nil{
		return errors.New("no connection to mysql!")
	}

	sql := fmt.Sprintf("insert into %s values (",d.table)
	
	struct_helper.ParseSimpleStruct(value,ToInsert,&sql,d.key)
	sql += ")"

	fmt.Println(sql)
	_,err := d.conn.Exec(sql)

	return err
}





