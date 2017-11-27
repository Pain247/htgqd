package process

import (
	"math"
)
var(
	pro_temp,pro_temp1 map[string]interface{}
	area,s2016,s2017,nv2,rank,amount1 float64
)

func GetAstar(arr []map[string]interface{}) map[string]interface{}{
	results := map[string]interface{}{
		"area": GetMax(arr,"area"),
		"s2016": GetMax(arr,"s2016"),
		"s2017": GetMax(arr,"s2017"),
		"nv2" : GetMax(arr,"nv2"),
		"amount": GetMax(arr,"amount"),
		"rank" : GetMax(arr,"rank"),
		"hobbies": GetMax(arr,"hobbies"),
	}
	return results
}
func GetAsub(arr []map[string]interface{}) map[string]interface{}{
	results := map[string]interface{}{
		"area": GetMin(arr,"area"),
		"s2016": GetMin(arr,"s2016"),
		"s2017": GetMin(arr,"s2017"),
		"nv2" : GetMin(arr,"nv2"),
		"amount": GetMin(arr,"amount"),
		"rank" : GetMin(arr,"rank"),
		"hobbies": GetMin(arr,"hobbies"),
	}
	return results
}
func GetSstar(arr []map[string]interface{}, maps map[string]interface{})[]map[string]interface{}{
	S := 0.0
	results := make([]map[string]interface{},0)
	for i:= 0; i<len(arr);i++{
		if(arr[i]["area"].(float64) != -1.0){
			area = math.Pow(arr[i]["area"].(float64)-maps["area"].(float64),2)
		}else{
			area = -1.0
		}
		if(arr[i]["s2016"].(float64) != -1.0){
			s2016 = math.Pow(arr[i]["s2016"].(float64)-maps["s2016"].(float64),2)
		}else{
			s2016 = -1.0
		}
		if(arr[i]["s2017"].(float64) != -1.0){
			s2017 = math.Pow(arr[i]["s2017"].(float64)-maps["s2017"].(float64),2)
		}else{
			s2017 = -1.0
		}
		if(arr[i]["nv2"].(float64) != -1.0){
			nv2 = math.Pow(arr[i]["nv2"].(float64)-maps["nv2"].(float64),2)
		}else{
			nv2 = -1.0
		}
		if(arr[i]["rank"].(float64) != 0.0){
			rank = math.Pow(arr[i]["rank"].(float64)-maps["rank"].(float64),2)
		}else{
			rank = -1.0
		}
		if(arr[i]["amount"].(float64) != 0.0 && arr[i]["amount"].(float64)!=-1.0){
			amount1 = math.Pow(arr[i]["amount"].(float64)-maps["amount"].(float64),2)
		}else{
			amount1 = -1.0
		}
		S =  math.Pow(arr[i]["hobbies"].(float64)-maps["hobbies"].(float64),2)
		if(area != -1.0){
			S+=area
		}
		if(s2016 != -1.0){
			S+=s2016
		}
		if(s2017 != -1.0){
			S+=s2017
		}
		if(nv2 != -1.0){
			S+=nv2
		}
		if(amount1 != -1.0){
			S+=amount1
		}
		if(rank != -1.0){
			S+=rank
		}
		pro_temp = map[string]interface{}{
			"id" : arr[i]["id"],
			"s": math.Sqrt(S),
		}
		results = append(results,pro_temp)
		pro_temp = make(map[string]interface{})
	        S = 0.0
	}
	return results
}

func Aggregate(mapstar []map[string]interface{}, mapsub []map[string]interface{}) []map[string]interface{}{
        results := make([]map[string]interface{},0)
	for i:=0;i<len(mapstar);i++{
		for j:=0;j<len(mapsub);j++{
			if (mapstar[i]["id"].(int) == mapsub[j]["id"].(int)){
				sstar := mapstar[i]["s"].(float64)
				ssub := mapsub[j]["s"].(float64)
				pro_temp1 = map[string]interface{}{
					"id" : mapstar[i]["id"],
					"s" : float64(ssub)/float64(ssub+sstar),
				}
				results = append(results,pro_temp1)
				pro_temp1 = make(map[string]interface{})
			}
		}
	}
	return results

}
func GetMaxC(arr []map[string]interface{}) (float64,int){
	max := 0.0
	id := 0

	for i:= 0;i<len(arr);i++{
		next,_ := (arr[i]["s"]).(float64)
		if(next >= max){
			max = next
			id = arr[i]["id"].(int)
		}
	}
	return max,id
}
func GetList(arr []map[string]interface{}) []map[string]interface{}{
	max := 0.0
	max1 := 0.0
	max2 := 0.0
	id := 0
	id1 := 0
	id2 := 0
        if len(arr) <= 3 {
		return arr
	}else{
		for i:= 0;i<len(arr);i++{
			next,_ := (arr[i]["s"]).(float64)
			if(next >= max){
				max2 = max1
				max1 = max
				max = next
				id2 = id1
				id1 = id
				id = arr[i]["id"].(int)
			}
		}
		results := []map[string]interface{}{{"stt": 1, "c": max, "id" :id},{"stt": 2, "c": max1, "id" :id1},{"stt": 3, "c": max2, "id" :id2}}
		return results
	}

}