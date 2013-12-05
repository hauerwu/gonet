package mongodb

import (
	"errors"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

type DAO struct {
	host    string
	port    string
	db      string
	col     string
	key     string
	session *mgo.Session
}

func New() *DAO {
	return new(DAO)
}

func (d *DAO) connect() error {
	var mgo_url string
	var err error

	if d.port == "" {
		mgo_url = d.host
	} else {
		mgo_url = d.host + ":" + d.port
	}
	d.session, err = mgo.Dial(mgo_url)
	if err != nil {
		fmt.Printf("[%s]: %s\n", mgo_url, err)
		return err
	}

	d.session.SetMode(mgo.Monotonic, true)
	fmt.Printf("connected to %s\n", mgo_url)

	return nil
}

func (d *DAO) Initialize(url string) error {
	strs := strings.Split(url, ":")
	if len(strs) < 5 {
		fmt.Printf("invalid url [%s]\n", url)
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

func (d *DAO) Finalize() {
	d.session.Close()
}

func (d DAO) Select(key interface{}, result interface{}) error {
	if d.session == nil {
		return errors.New("session invalid,please reconnect!")
	}

	c := d.session.DB(d.db).C(d.col)
	err := c.Find(bson.M{d.key: key}).All(result)
	return err
}

func (d DAO) Update(key interface{}, value interface{}) error {
	if d.session == nil {
		return errors.New("session invalid,please reconnect!")
	}

	c := d.session.DB(d.db).C(d.col)
	err := c.Update(bson.M{d.key: key}, value)
	return err
}

func (d DAO) Delete(key interface{}) error {
	if d.session == nil {
		return errors.New("session invalid,please reconnect!")
	}

	c := d.session.DB(d.db).C(d.col)
	err := c.Remove(bson.M{d.key: key})
	return err
}

func (d DAO) Insert(value interface{}) error {
	if d.session == nil {
		return errors.New("session invalid,please reconnect!")
	}

	c := d.session.DB(d.db).C(d.col)
	err := c.Insert(value)
	return err
}
