package main
var (
	DB *sql.DB
	err error
)

func main(){

}

func InitDB(){
	DB, err = sql.Open("sqlite3", "./db/healthDB.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
		return
	}
}


func CreateTable() error {
	for _, table := range Tables{
		_, err := DB.Exec(table)
		if err !=  nil {
			return err
		}
	}
	return nil
}


