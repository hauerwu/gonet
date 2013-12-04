package mysql

import(
	"strings"
	"fmt"
	"errors"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DAO struct{
	host string
	port string
	db string
	table string
	key string
}

func New() *DAO{
	return new(DAO)
}

func (d *DAO)connect() error{
	return nil
}

func (d *DAO)Initialize(url string) error{
	strs := strings.Split(url,":")
	if len(strs) < 5{
		fmt.Printf("invalid url [%s]\n",url)
		return errors.New("can't parse ulr,try like this: 127.0.0.1:80:test:peple:key")
	}

	d.host = strs[0]
	d.port = strs[1]
	d.db = strs[2]
	d.table = strs[3]
	d.key = strs[4]

	err := d.connect()

	return err
}

func (d *DAO)Finalize(){
}

func (d DAO)Select(key interface{},result interface{}) (error){
	return nil
}

func (d DAO)Update(key interface{},value interface{}) (error){
	return nil
}

func (d DAO)Delete(key interface{}) (error){
	return nil
}

func (d DAO)Insert(value interface{}) (error){
	return nil
}

