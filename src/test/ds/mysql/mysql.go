package mysql

import (
	"errors"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
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

	for rows.Next(){
		var id int
		var name string
		err = rows.Scan(&id,&name)
		fmt.Printf("id:%d name:%s\n",id,name)
	}
	
	return err
}

func (d DAO) Update(key interface{}, value interface{}) error {
	return nil
}

func (d DAO) Delete(key interface{}) error {
	return nil
}

func (d DAO) Insert(value interface{}) error {
	return nil
}
