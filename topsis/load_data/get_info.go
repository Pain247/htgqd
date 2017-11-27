package load_data

import (
	"fmt"
	"github.com/tuyensinh/database/mysql"
	"github.com/tuyensinh/utils"
	"strconv"
)
var(
	temp_info map[string]interface{}
)
func LoadHobbies() []map[string]interface{} {
	myclient := &mysql.MySQLClient{Config:utils.LoadConfigSql(mysql.PATH_CONFIG_SQL)}
	err := myclient.CreateMySqlClient(myclient.Config.MYSQL_DB_USER, myclient.Config.MYSQL_DB_PASS, myclient.Config.MYSQL_DB_HOST, myclient.Config.MYSQL_DB_NAME)
	defer myclient.Close()
	if err != nil {
		fmt.Println("FAILED TO CONNECT MYSQL!")
		fmt.Println(err)
		return nil
	} else {
           query := "select id,name from hobbies"
	   rows := myclient.ReadTable(query)
		results := make([]map[string]interface{},0)
		for _,v  := range rows{
			id,_ := strconv.Atoi(v[0])
			name := v[1]
			temp_info = map[string]interface{}{
				" id" : id,
				 "name" : name,
			}
			results = append(results,temp_info)
			temp_info = make(map[string]interface{})
		}
		return SortAtt(results)
	}
}
func LoadGroup() []string{
	myclient := &mysql.MySQLClient{Config:utils.LoadConfigSql(mysql.PATH_CONFIG_SQL)}
	err := myclient.CreateMySqlClient(myclient.Config.MYSQL_DB_USER, myclient.Config.MYSQL_DB_PASS, myclient.Config.MYSQL_DB_HOST, myclient.Config.MYSQL_DB_NAME)
	defer myclient.Close()
	if err != nil {
		fmt.Println("FAILED TO CONNECT MYSQL!")
		fmt.Println(err)
		return nil
	} else {
		query := "select code from group"
		rows := myclient.ReadTable(query)
		results := make([]string,0)
		for _,v  := range rows{
			results = append(results,v[0])
		}
		return results
	}
}
func GetInfoDep(id int) map[string]interface{}{
	myclient := &mysql.MySQLClient{Config:utils.LoadConfigSql(mysql.PATH_CONFIG_SQL)}
	err := myclient.CreateMySqlClient(myclient.Config.MYSQL_DB_USER, myclient.Config.MYSQL_DB_PASS, myclient.Config.MYSQL_DB_HOST, myclient.Config.MYSQL_DB_NAME)
	defer myclient.Close()
	if err != nil {
		fmt.Println("FAILED TO CONNECT MYSQL!")
		fmt.Println(err)
		return nil
	} else {
		query := "select name,code,university_id from department where id = "+strconv.Itoa(id)
		rows := myclient.ReadTable(query)
		uni,_ := strconv.Atoi(rows[0][2])
		arr := GetInfoUni(uni)
		return  map[string]interface{}{
			"id" : id,
			"dep_name" : rows[0][0],
			"code" : rows[0][1],
			"uni_name": arr[0],
			"area" : arr[1],
		}
	}
}
func GetInfoUni(id int) []string{
	myclient := &mysql.MySQLClient{Config:utils.LoadConfigSql(mysql.PATH_CONFIG_SQL)}
	err := myclient.CreateMySqlClient(myclient.Config.MYSQL_DB_USER, myclient.Config.MYSQL_DB_PASS, myclient.Config.MYSQL_DB_HOST, myclient.Config.MYSQL_DB_NAME)
	defer myclient.Close()
	if err != nil {
		fmt.Println("FAILED TO CONNECT MYSQL!")
		fmt.Println(err)
		return nil
	} else {
		query := "select name,area from university where id ="+ strconv.Itoa(id)
		rows := myclient.ReadTable(query)
		return rows[0][0]
	}
}
func SortAtt(input []map[string]interface{}) []map[string]interface{}{
	n := len(input)
	if n == 1 || n==0 {
		return input
	}
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if input[i-1]["id"].(int) > input[i]["id"].(int) {
				// swap values using Go's tuple assignment
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}