package main

import (
	"../process"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	var conn process.DbInstance

	fmt.Println("------- Welcome to warioDB command line ------")

	for {
		fmt.Print("Enter the database name: ")
		// dbName, _ := reader.ReadString('\n')
		var dbName string
		_, _ = fmt.Scanf("%s", &dbName)
		conn = process.Connect(dbName)
		if conn.IsConnected {
			break
		}
		process.Disconnect()
		fmt.Println("The database does not exist")
		fmt.Println("-------")
	}
	fmt.Println("-------")
	fmt.Printf("Connected to %s \n", conn.DbName)
	for {
		fmt.Printf(">> ")
		var str1, str2, str3, str4 string
		_, _ = fmt.Scanf("%s %s %s %s", &str1, &str2, &str3, &str4)
		if str1 == "Q" || str1 == "q" {
			break
		}

		switch str1 {
		case "select":
			for _, table := range conn.Tables {
				if table.Name == str4 && str3 == "from" {
					if str2 == "*" {
						var dataOutput [][]string
						var headerOutput []string
						tableOutput := tablewriter.NewWriter(os.Stdout)

						for _, column := range table.Columns {
							//fmt.Printf("%s %s %s %s", str1, str2, str3, str4)
							headerOutput = append(headerOutput, column.Name)
							//fmt.Println(headerOutput)
							//var columnsVerticalValues []string
							//dataOutput =  append(dataOutput, columnsVerticalValues)
						}
						dataOutput = process.SelectAll(conn.DbName, table.Name)
						tableOutput.SetHeader(headerOutput)
						tableOutput.SetBorder(false)
						tableOutput.AppendBulk(dataOutput)
						tableOutput.Render()
						/*data := [][]string{
							{"1/1/2014", "Domain name", "2233", "$10.98"},
							{"1/1/2014", "January Hosting", "2233", "$54.95"},
							{"1/4/2014", "February Hosting", "2233", "$51.00"},
							{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
						}
						table := tablewriter.NewWriter(os.Stdout)
						table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
						table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
						table.SetBorder(false)                                // Set Border to false
						table.AppendBulk(data)                                // Add Bulk Data
						table.Render()*/
					}
				}
			}
		case "udtape":
			fmt.Println("In process")
		case "delete":
			fmt.Println("In process")
		case "create":
			fmt.Println("In process")

		}
	}

}
