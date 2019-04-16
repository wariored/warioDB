package process

import (
	"fmt"
	"os"
)

type DbInstance struct {
	DbName      string   // DB name
	IsConnected bool     // if true the db exists
	Tables      [] Table // db tables located in json file format {DBName}.json
}

// check if the file exist in the "source" folder
func dbExits(dbName string) bool {
	dbFileRoot := DbPath + dbName + ".json" // DbPath is locate in "statics.go" file
	if _, err := os.Stat(dbFileRoot); os.IsNotExist(err) {
		fmt.Println(err)
		return false
	}
	return true
}

// global variable, as a singleton for Connect func
var instance DbInstance

// connect to the database
func Connect(dbName string) DbInstance {
	if dbExits(dbName) {
		if !instance.IsConnected {
			instance.IsConnected = true
			instance.DbName = dbName
			GetTables(&instance) // get database's tables
		}
		// return instance
	}
	// panic("The database doesn't exists")
	return instance
}

// disconnect to the DB
func Disconnect() {
	// reinitialize the singleton
	instance = DbInstance{}

}
