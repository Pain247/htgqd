package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
	"github.com/tuyensinh/utils"
)
type MySQLClient struct {
	db *sql.DB
	Config utils.ConfigSql

}

func (myclient *MySQLClient)CreateMySqlClient(user string, pass string,host string, dbname string) error{
	//connect to db----------------------------------------
	dsn := user + ":" + pass + "@" + host + "/" + dbname + "?charset=utf8"
	var err error
	myclient.db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Faild to connect mysql server!")
		return err
	}else{
		return nil
	}
}

func (myclient *MySQLClient)ReadTable(query string) (map[int][]string){//return data row can read

	//query := "SELECT id,url,plus_bid,number_view_calculating FROM sspdb.dsp_info where isactive = 1;"
	// Execute the query
	rows, err := myclient.db.Query(query)
	if err != nil {
		return nil
	}
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	// Fetch rows
	numRow := 0
	resultMap := make(map[int][]string)
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil
		}
		arow := make([]string, len(columns))
		for k,v := range values{
			arow[k] = string(v)
		}
		resultMap[numRow] = arow
		numRow ++
		//fmt.Println("---------------------------------------------------------------------------------")
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	rows.Close()
	return resultMap
}

func (myclient *MySQLClient)CountRow(query string) int{//return number row can read

	//query := "SELECT id,url,plus_bid,number_view_calculating FROM sspdb.dsp_info where isactive = 1;"
	// Execute the query
	rows, err := myclient.db.Query(query)
	if err != nil {
		return 0
	}
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return 0
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	// Fetch rows
	numRow := 0
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return 0
		}
		// Now do something with the data.
		// Here we just print each column as a string.

		i, err2 := strconv.Atoi(string(values[0]))
		if err2 != nil {
			numRow = 0
		}
		numRow = i
		//fmt.Println("---------------------------------------------------------------------------------")
	}
	if err = rows.Err(); err != nil {
		return 0
	}
	rows.Close()
	return numRow
}

func (myclient *MySQLClient) InsertRowToMySQL(query string,args map[int][]interface{}) {
	// Prepare statement for inserting data
	stmtIns, err := myclient.db.Prepare(query) // ? = placeholder
	if err != nil {
		fmt.Println("Error in prepare inserting db mysql!")
		fmt.Println(err)
		return
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	// Insert to the database
	for _, values := range args{
		_, err = stmtIns.Exec(values ...)
		if err != nil {
			fmt.Println("Error in inserting db mysql!")
			fmt.Println(err)
			continue
		}
	}
}

func (myclient *MySQLClient) Close(){
	myclient.db.Close()
}