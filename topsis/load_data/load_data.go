package load_data

import (
	"fmt"
	"github.com/tuyensinh/database/mysql"
	"github.com/tuyensinh/utils"
	"strconv"
	"math/rand"
	"time"
)
var (
	temp map[string]interface{}
)

func LoadDepartment() []map[string]interface{}{
	myclient := &mysql.MySQLClient{Config:utils.LoadConfigSql(mysql.PATH_CONFIG_SQL)}
	err := myclient.CreateMySqlClient(myclient.Config.MYSQL_DB_USER,myclient.Config.MYSQL_DB_PASS,myclient.Config.MYSQL_DB_HOST,myclient.Config.MYSQL_DB_NAME)
	defer myclient.Close()
	if err != nil{
		fmt.Println("FAILED TO CONNECT MYSQL!")
		fmt.Println(err)
		return nil
	}else{
		query := "select id,name,code,university_id,hobbies_id from department"
		rows := myclient.ReadTable(query)
		results := make([]map[string]interface{},0)
		for _,v  := range rows{
                  id,_ := strconv.Atoi(string(v[0]))
		  name := string(v[1])
		  code := string(v[2])
		  uni_id,_ := strconv.Atoi(string(v[3]))
                  hob_id,_ := strconv.Atoi(string(v[4]))
		  info := getArea(myclient,uni_id)
		  area,_ := strconv.Atoi(info[0])
		  rank,_ := strconv.ParseFloat(info[1],64)
		  s2016 := getScore2016(myclient,id)
			if(s2016 >= 30.0){s2016 = -1.0}
		  s2017 := getScore2017(myclient,id)
			if(s2017 >= 30.0){s2017 = -1.0}
		  nv2 := getNV2(myclient,id)
			if(nv2 >= 30.0){nv2 = -1.0}
			if(nv2 <= 0.0){
				nv2 = -1.0
			}
			temp = map[string]interface{}{
				"id" : id,
				"name": name,
				"code" : code,
				"university_id": uni_id,
				"hobbies_id": hob_id,
				"group" : getGroup(myclient,id),
				"area" : area,
				"rank" : rank,
				"s2016": s2016 ,
				"s2017" : s2017,
				"nv2" : nv2,
				"amount": getAmount(myclient,id,nv2),
			}
			results = append(results,temp)
			temp = make(map[string]interface{})
		}
		return Sort(results)
	}
      return nil
}

func getGroup(myclient *mysql.MySQLClient, id int) []string{
	query := "select group_code from department_has_group where department_id = "+strconv.Itoa(id)
	rows := myclient.ReadTable(query)
	results := make([]string,0)
	for _,v := range rows{
		results = append(results,string(v[0]))
	}
	return results
}
func getArea(myclient *mysql.MySQLClient, id int) []string{
	query := "select area,score from university where id = "+strconv.Itoa(id)
	rows := myclient.ReadTable(query)
	return rows[0]
}
func getScore2016(myclient *mysql.MySQLClient, id int) float64{
	query := "select score from first_choice where department_id = "+strconv.Itoa(id) +" and year = 2016"
	rows := myclient.ReadTable(query)
	if(len(rows)!=0){
		score,_ := strconv.ParseFloat(rows[0][0],64)
		return norScore(score)
	}else{
		query := "select score from second_choice where department_id = "+ strconv.Itoa(id)
		rows := myclient.ReadTable(query)
		if(len(rows) != 0){
			score,_ := strconv.ParseFloat(rows[0][0],64)
		        return norScore(score)
		}else { return -1.0 }
	}
}
func getScore2017(myclient *mysql.MySQLClient, id int) float64{
	query := "select score from first_choice where department_id = "+strconv.Itoa(id) +" and year = 2017"
	rows := myclient.ReadTable(query)
	if(len(rows) != 0 ){
		score,_ := strconv.ParseFloat(rows[0][0],64)
		return score
	}else{
		return -1.0
	}
}

func getAmount(myclient *mysql.MySQLClient,id int, num float64) int{
	query := "select amount from second_choice where department_id = "+ strconv.Itoa(id)
	rows := myclient.ReadTable(query)
	if(len(rows) != 0){
		if(rows[0][0]=="0"){
			return Amount(num)
		}
		amount,_ := strconv.Atoi(rows[0][0])
		return amount
	}else{
               return Amount(num)
	}

}
func norScore(num float64) float64{
	if (num > 30){ num = float64(num*3)/4.0}
	return num
}
func Amount(num float64) int{
	if num >= 0 && num <10{
		return 150
	}else{
		if(num<=20){
			return 100
		}else{
			if(num<25){
				return 50
			}else{
				return 10
			}
		}
	}

}
func getNV2(myclient *mysql.MySQLClient, id int) float64{
	query := "select score from second_choice where department_id = "+ strconv.Itoa(id)
	rows := myclient.ReadTable(query)
        if(len(rows) != 0){
		num,_ := strconv.ParseFloat(rows[0][0],64)
		return norNV2(num)
	}else{
		return -1.0
	}
}

func norNV2(num float64) float64 {
	if num == -1.0 {
		return num
	}else{
		if(num >3){
			rand.Seed(time.Now().UTC().UnixNano())
			n := rand.Intn(10)
			if(n>= 0 && n<3) {num = num - 0.75}
			if(n>=3 && n<6){ num = num - 0.5}
			if(n>=6 && n<=7){ num = num - 0.25}
			if(n==9){ num = num + 0.25}
			if (n==10){ num = num + 0.5}
		}
		return num

	}
}
func Sort(input []map[string]interface{}) []map[string]interface{}{
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