package main
import (
	"fmt"
	"os"
	"bufio"
    "log"
	dbrepo "./dbrepository"
	mongoutils "./utils"
	domain "./domain"
	"encoding/json"
)

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)
	//Run sample commands
//first assign

file, err := os.Open("./restaurant.json")
 if err != nil {
     log.Fatal(err)
 }
 defer file.Close()
 var cnt int
 var data domain.Restaurant
 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
	 p := []byte(scanner.Text())
     json.Unmarshal(p, &data)
	 //data.DBID = domain.String(domain.NewID())
	 data.DBID =domain.NewID()
	 did,_ := repoaccess.Store(&data)
	 if did== domain.ID(0){
		fmt.Println("Error in Insert")
		break
	} else {
		cnt = cnt+1
	}
}
fmt.Println("record inserted ",cnt)
}
