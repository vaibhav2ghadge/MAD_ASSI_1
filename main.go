package main

import (
	"fmt"
	"os"
	"bufio"
    "log"
	"strings"
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
	// break
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
fmt.Println("Number Of Record Inserted : ",cnt)
/*
//second assign
var id domain.ID;
id = "5c45653b79492b6216397592"
x,y:=repoaccess.Get(id)
fmt.Println("get",x,y)
xx,yy:=repoaccess.GetAll()
fmt.Println("getAll",xx,yy)
for _,obj:=range xx {
			fmt.Println(obj,yy)
	}
//third assignment
*/
cmdArgument := os.Args[1:]
if len(cmdArgument)>0 {
	if strings.Contains(cmdArgument[0],"find") {
			cmdArgument = strings.SplitAfter(cmdArgument[1],"=")
			if strings.Contains(cmdArgument[0],"--type_of_food") {
				fmt.Println("Types Of Food")
				rest,err := repoaccess.FindByTypeOfFood(cmdArgument[1])
				for _,obj:=range rest {
					fmt.Println(obj,err)
				}
				fmt.Println("invalid1")
			} else if strings.Contains(cmdArgument[0],"--postcode") && len(cmdArgument)==2{
				fmt.Println("argument",cmdArgument,cmdArgument[1])
				rest,err := repoaccess.FindByTypeOfPostCode(cmdArgument[1])
				for _,obj:=range rest {
					fmt.Println(obj)
				}
				fmt.Println("invalid3",err)
			} else {
				fmt.Println("invalid")
			}//count number of restuarnt in give food type
	} else if strings.Contains(cmdArgument[0],"count") {
			cmdArgument = strings.SplitAfter(cmdArgument[1],"=")
			if strings.Contains(cmdArgument[0],"--type_of_food") {
				count,err := repoaccess.CountByTypeOfFood(cmdArgument[1])
				fmt.Println("Total Restarant : ",count,err)
	}
}
}}
