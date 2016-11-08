
package demomgo

import (
        "fmt"
	"log"
        "time"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Operation struct {
        Description string
        Timestamp time.Time
}

type Blathering struct {
        Blather string
        Timestamp time.Time
}


func BlatherMe(b string) {

        fmt.Println("hi neil ", b)

        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("blatherings")
        err = c.Insert(&Blathering{b, time.Now()})
        if err != nil {
                log.Fatal(err)
        }


}

func BlahInsert() {

        fmt.Println("hi neil")
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("operations")
        err = c.Insert(&Operation{"I am starting up the application", time.Now()},
	               &Operation{"And I am ready to go to work", time.Now()})
        if err != nil {
                log.Fatal(err)
        }

        result := Operation{}
        err = c.Find(bson.M{"description": "I am starting up the application"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Operation:", result.Description)
}