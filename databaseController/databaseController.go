package databaseController

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"strings"
	"os/exec"
)


func InitDB(filepath string) *sql.DB{
	db, err := sql.Open("sqlite3", filepath)
	if err != nil{
		fmt.Println("database is broken")
	}
	if db == nil{
		fmt.Println("could not load database")
	}
	return db
}

func command_exists(commandName string) bool{
	output, _ := exec.Command("which", commandName).Output()
	if string(output) == ""{
		return false
	}else{
		return true
	}
}

func GetAllColumns(db *sql.DB) (string, string){
	sql_readall := `
	SELECT name, command FROM commands;
	`
	output, _ := db.Query(sql_readall) 
	defer output.Close() 
	var name string 
	var command string 
	for output.Next(){
		output.Scan(&name, &command)

	}
	return name, command
}

func GetCommandByName(itemName string, db *sql.DB) ([]string, string){

	sql_readall := fmt.Sprintf(`
	SELECT command, message FROM commands WHERE name == "%v" AND os == "linux";
	`, itemName)
	output, _ := db.Query(sql_readall)
	defer output.Close()
	var name [10]string
	var finalCommand string
	var message string
	count := 0
	for output.Next(){
		output.Scan(&name[count], &message)
		placeHolder := strings.Split(name[count], " ")
		if command_exists(placeHolder[0]){
			finalCommand = name[count]
			break
		}
		count++

	}
	convertCommandToArray := strings.Split(finalCommand, " ")
	return convertCommandToArray, message
}

