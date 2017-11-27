package process

import (
	"reflect"
	"strings"
	"math"
	"fmt"
)

var(
	temp1,temp3,temp4 map[string]interface{}
	x16,x17,x2,amount float64
)
//[]map[string]interface{}
func GetMax(arr []map[string]interface{}, att string) float64{
	max := 0.0
	for i:= 0;i<len(arr);i++{
		next,_ := (arr[i][att]).(float64)
		if(next >= max){
			max = next
		}
	}
	return max
}

func GetMin(arr []map[string]interface{}, att string) float64{
	min := 30.0
	for i:= 0;i<len(arr);i++{
		next,_ := (arr[i][att]).(float64)
		if(next <= min && next!= -1){
			min = next
		}
	}
	return min
}

func Normalize(x float64, max float64, min float64) float64{
	return float64(x - min)/float64(max-min)
}
func NormalizeData(arr []map[string]interface{}) []map[string]interface{}{
        max2016 := GetMax(arr,"s2016")
	min2016 := GetMin(arr,"s2016")
	max2017 := GetMax(arr,"s2017")
	min2017 := GetMin(arr,"s2017")
	maxnv2 := GetMax(arr,"nv2")
	minnv2 := GetMin(arr,"nv2")
	maxamount := GetMaxAmount(arr,"amount")
	minamount := GetMinAmount(arr,"amount")
	maxrank := GetMax(arr,"rank")
	minrank := GetMin(arr,"rank")
	fmt.Println(minrank)
	results := make([]map[string]interface{},0)
	for i:=0;i<len(arr);i++{
		if(arr[i]["s2016"].(float64) != -1.0){
                     x16 = Normalize(arr[i]["s2016"].(float64),max2016,min2016)*0.15
		}else{
			x16 = arr[i]["s2016"].(float64)
		}
		if(arr[i]["s2017"].(float64) != -1.0){
			x17 = Normalize(arr[i]["s2017"].(float64),max2017,min2017)*0.15
		}else{
			x17 = arr[i]["s2017"].(float64)
		}
		if(arr[i]["nv2"].(float64) != -1.0){
			x2 = Normalize(arr[i]["nv2"].(float64),maxnv2,minnv2)*0.05
		}else{
			x2 = arr[i]["nv2"].(float64)
		}
		amount = Normalize(float64(arr[i]["amount"].(int)),maxamount,minamount)*0.05
		area := Normalize((arr[i]["area"].(float64)),0,-2)*0.2
		rank := Normalize(arr[i]["rank"].(float64),maxrank,minrank)*0.1
		temp1 = map[string]interface{}{
			"id" : arr[i]["id"],
			"s2016": x16,
			"s2017": x17,
			"nv2": x2,
			"amount": amount,
			"area": area,
			"rank" : rank,
			"hobbies": arr[i]["hobbies"].(float64)*0.3,
		}
		results = append(results,temp1)
		temp1 = make(map[string]interface{})
	}
	return results

}

func ConvertData(group string,score float64,area int, hobbies int,arr []map[string]interface{}) []map[string]interface{}{
	temp2 := make([]map[string]interface{},0)
	results := make([]map[string]interface{},0)
        for i:= 0;i<len(arr);i++{
           if(checkGroup(group,arr[i]["group"])){
		   temp2 = append(temp2,arr[i])
	   }
	}
	for i:=0;i<len(temp2);i++{
		temp3 = map[string]interface{}{
			"id" : temp2[i]["id"],
			"name": temp2[i]["name"],
			"code" : temp2[i]["code"],
			"university_id": temp2[i]["university_id"],
			"group" : temp2[i]["group"],
			"area" :-math.Abs(float64(area - temp2[i]["area"].(int))),
			"rank" : temp2[i]["rank"],
			"s2016": score - temp2[i]["s2016"].(float64) ,
			"s2017" : score - temp2[i]["s2017"].(float64),
			"nv2" : score - temp2[i]["nv2"].(float64),
			"amount": temp2[i]["amount"].(int),
			"hobbies" : GetHobbies(hobbies,temp2[i]["hobbies_id"].(int)),
		}
		results = append(results,temp3)
		temp3 = make(map[string]interface{})
	}
       return results
}

func checkGroup(group string, arr interface{}) bool{
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i:=0;i<s.Len();i++{
			if (strings.Compare(group,s.Index(i).Interface().(string))==0){
				return true
			}
		}
	}
	return false
}
//Due to amount is different type
func GetMaxAmount(arr []map[string]interface{}, att string) float64{
	max := 0
	for i:= 0;i<len(arr);i++{
		next,_ := (arr[i][att]).(int)
		if(next >= max){
			max = next
		}
	}
	return float64(max)
}

func GetMinAmount(arr []map[string]interface{}, att string) float64{
	min := 30
	for i:= 0;i<len(arr);i++{
		next,_ := (arr[i][att]).(int)
		if(next <= min && next!= -1){
			min = next
		}
	}
	return float64(min)
}

//Get Hobbies
func GetHobbies(hob_user int, hob_dep int) float64{
	arr := [14][14]float64{
		{1,0.2,0.5,0.57,0.22,0.6,0.8,0.3,0.45,0.28,0.52,0.2,0.35,0.1},
		{0.2,1,0.4,0.25,0.3,0.3,0.4,0.3,0.2,0.45,0.4,0.2,0.2,0.1},
		{0.5,0.4,1,0.7,0.2,0.56,0.6,0.2,0.5,0.2,0.85,0.4,0.35,0.2},
		{0.57,0.25,0.7,1,0.3,0.6,0.55,0.3,0.5,0.2,0.7,0.2,0.21,0.1},
		{0.22,0.3,0.2,0.3,1,0.35,0.35,0.5,0.5,0.2,0.3,0.7,0.57,0.1},
		{0.6,0.3,0.56,0.6,0.35,1,0.45,0.4,0.3,0.25,0.55,0.3,0.43,0.12},
		{0.8,0.4,0.6,0.55,0.35,0.45,1,0.47,0.45,0.4,0.5,0.43,0.25,0.5},
		{0.3,0.3,0.2,0.3,0.5,0.4,0.47,1,0.3,0.15,0.3,0.6,0.4,0.3},
		{0.45,0.2,0.5,0.5,0.5,0.3,0.45,0.3,1,0.15,0.7,0.45,0.47,0.2},
		{0.28,0.45,0.2,0.2,0.2,0.25,0.4,0.15,0.15,1,0.15,0.2,0.15,0.1},
		{0.52,0.4,0.85,0.7,0.3,0.55,0.5,0.3,0.7,0.15,1,0.35,0.4,0.25},
		{0.2,0.2,0.4,0.2,0.7,0.3,0.43,0.6,0.45,0.2,0.35,1,0.65,0.2},
		{0.35,0.2,0.35,0.21,0.57,0.43,0.25,0.4,0.47,0.15,0.4,0.65,1,0.18},
		{0.1,0.1,0.2,0.1,0.1,0.12,0.5,0.3,0.2,0.1,0.25,0.2,0.18,1},
	}
	return arr[hob_user-1][hob_dep-1]
}