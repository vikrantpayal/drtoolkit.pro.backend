package main

import (
    "database/sql"
    "fmt"
    _  "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "drdev_backend"
    password = "bhejaFry#3"
    dbname = "drdev"
)

func main(){
    // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

   // open db
   db, err := sql.Open("postgres", psqlconn)
   CheckError(err)


    rows, err := db.Query("select * from patients")
    CheckError(err)

    defer rows.Close()

    for rows.Next(){
        var dr int
        var patient_id int
        var patient_fname string
        var patient_lname string

        err = rows.Scan(&dr, &patient_id, &patient_fname, &patient_lname)
        CheckError(err)
        fmt.Println(dr, patient_id, patient_fname, patient_lname)
    }

   CheckError(err)

   // close db
   defer db.Close()

   // check db
   err = db.Ping()
   CheckError(err)

   fmt.Println("Connected")

}


func CheckError(err error){
    if err != nil {
        panic(err)
    }
}
