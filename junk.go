package main
import ( "fmt"
		 "database/sql"
		 _"github.com/go-sql-driver/mysql"
		 )
func main() {
	fmt.Println("BEGINNNG DB CONNECTION!")

	db,err := sql.Open("mysql","root:**@tcp(127.0.0.1)/mydb") //define database and error
	if err!=nil{
		panic(err.Error()) //panic prints the error.
	}
	fmt.Println("Connected")
	rows,err := db.Query("SELECT * FROM names")
	checkErr(err)
	for rows.Next(){

		var LastName string //define values
		var FirstName string
		var RegNo int
		var BRANCH string
		err = rows.Scan(&LastName, &FirstName, &RegNo, &BRANCH) //read values
		checkErr(err)
		fmt.Println("LAST NAME:",LastName)
		fmt.Println("FIRST NAME",FirstName)
		fmt.Println("REG NO:",RegNo)
		fmt.Println("BRANCH:",BRANCH)
	}
	defer db.Close()
}
func checkErr(err error) {

	if err!=nil { //check errors
		panic(err)
	}
}