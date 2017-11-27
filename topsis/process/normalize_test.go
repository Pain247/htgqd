package process

import (
	"testing"
	"github.com/tuyensinh/topsis/load_data"
	"fmt"
)

func TestGetMax(t *testing.T) {
	arr := NormalizeData(ConvertData("A1",16.0,1,7,load_data.LoadDepartment()))
	for i:=0;i<len(arr);i++{
		if(arr[i]["rank"]==0.0){
			fmt.Println(arr[i])
		}
	}
}
func TestGetHobbies(t *testing.T) {
	fmt.Println(GetHobbies(7,13))
}