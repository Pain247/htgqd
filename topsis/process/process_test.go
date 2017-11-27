package process

import (
	"testing"
	"fmt"
	"github.com/tuyensinh/topsis/load_data"
)

func TestGetAplus(t *testing.T) {
	arr := NormalizeData(ConvertData("A1",20.0,1,5,load_data.LoadDepartment()))
	mapsub := GetAsub(NormalizeData(ConvertData("A1",20.0,1,5,load_data.LoadDepartment())))
	mapstar := GetAstar(NormalizeData(ConvertData("A1",20.0,1,5,load_data.LoadDepartment())))
        sstar := GetSstar(arr,mapstar)
	ssub := GetSstar(arr,mapsub)
	agg := Aggregate(sstar,ssub)
	fmt.Println(GetMaxC(agg))
}
