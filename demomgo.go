
package demomgo

import (
        "fmt"
	"log"
        "time"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
//        "text/template"
//        "os"
//        "bytes"

)

type Operation struct {
        Description string
        Timestamp time.Time
}

type Blathering struct {
        Blather string
        Timestamp time.Time
}



func FullReport(anchor string) []Blathering {

    var results []Blathering
    session, err := mgo.Dial("localhost,localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()
    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)
    cst := session.DB("test").C("blatherings")
    err = cst.Find(bson.M{"blather": bson.RegEx{anchor,"i"}}).Sort().All(&results)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Results For: ", anchor)
    fmt.Println("Results All: ", results)


//    tmpl, err := template.New("test").Parse("{{ . }}")
  //  if err != nil { panic(err) }
    //err = tmpl.Execute(os.Stdout, results)
    //if err != nil { panic(err) }



    // buf := new(bytes.Buffer)
    // tmpl.Execute(buf, results)
    // fmt.Println("Results From the Crazy Template Stuff: ", buf.String())
    // return buf.String()

    return results


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