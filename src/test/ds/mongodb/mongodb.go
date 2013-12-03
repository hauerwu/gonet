package mongodb

import (
	"errors"
	"strings"
	_ "fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type DAO struct{
	host string
	port string
	db string
	col string
	key string
	session *mgo.Session
}

func New() *DAO{
	return &DAO{}
}

func (d *DAO)connect() error{
	var mgo_url string
	var err error

	if d.port == ""{
		mgo_url = d.host
	}else{
		mgo_url = d.host + ":" + d.port
	}
	d.session,err = mgo.Dial(mgo_url)
	if err != nil{
		return err
	}

	d.session.SetMode(mgo.Monotonic, true)

	return nil
}

func (d DAO)Initialize(url string) error{
	strs := strings.Split(url,":")
	if len(strs) < 5{
		return errors.New("can't parse ulr,try like this: 127.0.0.1:80:test:peple:key")
	}

	d.host = strs[0]
	d.port = strs[1]
	d.db = strs[2]
	d.col = strs[3]
	d.key = strs[4]

	err := d.connect()

	return err
}

func (d DAO)Finalize(){
	d.session.Close()
}

func (d DAO)Select(key interface{},result interface{}) (error){
	if d.session == nil{
		return errors.New("session invalid,please reconnect!")
	}

	c := d.session.DB(d.db).C(d.col)
	err := c.Find(bson.M{d.key:key}).All(result)
	return err
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











