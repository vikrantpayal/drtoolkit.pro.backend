package main
import(
    "net/http"
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
)
/*import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)*/

const (
	host     = "localhost"
	port     = 5432
	user     = "drdev_backend"
	password = ""
	dbname   = "drdev"
)



func main() {
    router :=gin.Default()
    router.Use(static.Serve("/", static.LocalFile("./views", true)))
	api := router.Group("/api")
	{
	  api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
		  "message": "pong",
		})
	  })
	}
	api.GET("/patients", GetPatients)
	router.Run(":3000")  

	//GetPatients()
}



func GetPatients(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
	  "message":"Patients handler not implemented yet",
	})
	/*
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open db
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	rows, err := db.Query("select * from patients")
	CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var patient_id int
		var patient_fname string
		var patient_lname string

		err = rows.Scan(&patient_id, &patient_fname, &patient_lname)
		CheckError(err)
		fmt.Println(patient_id, patient_fname, patient_lname)
	}

	CheckError(err)

	// close db
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
	*/
}
