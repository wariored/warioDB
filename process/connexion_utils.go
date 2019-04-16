package process

import (
	"encoding/json"
	"io/ioutil"
)

type Db struct {
	Tables [] Table
}
type Table struct {
	Name    string `json:"name"`
	Columns [] Column
}
type Column struct {
	Name       string   `json:"name"`
	ColumnType string   `json:"type"`
	Length     int   `json:"length"`
	Values     []string `json:"values"`
}

func GetTables(dbInstance *DbInstance) {
	dbFilePath := DbPath + dbInstance.DbName + ".json"

	file, _ := ioutil.ReadFile(dbFilePath)
	data := Db{}
	_ = json.Unmarshal([]byte (file), &data)
	//fmt.Println(err)
	//fmt.Println(string(file))
	//fmt.Println(data)
	dbInstance.Tables = data.Tables
	for i:= 0; i <  len( data.Tables); i++  {
		data.Tables[i].Columns = getColumns(dbInstance.DbName,  data.Tables[i].Name)
	}

}

func getColumns(DbName, tableName string) []Column {
	tableFilePath := DbPath + DbName + "." + tableName + ".json"
	file, _ := ioutil.ReadFile(tableFilePath)
	data := Table{}
	_ = json.Unmarshal([]byte (file), &data)
	//fmt.Println(data.Columns)
	return data.Columns

}
