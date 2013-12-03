package main

import (
        "fmt"
        "labix.org/v2/mgo"
        _ "labix.org/v2/mgo/bson"
)

type Person struct {
        Name string
        Phone string
	Id []int
}

func main() {
        session, err := mgo.Dial("127.0.0.1")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("mgos")
        err = c.Insert(&Person{"Ale", " +55 53 8116 9639",[]int{1001,1002}},
	               &Person{"Cla", " +55 53 8402 8510",[]int{1003,1004}})
        if err != nil {
                panic(err)
        }

	result := Person{}
	it := c.Find(nil).Iter()
	for {
	    if !it.Next(&result){
	       break
	    }

	    fmt.Println(result)
	}

/*
        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
                panic(err)
        }
*/
}