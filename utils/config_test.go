package utils

import (
	"testing"
)

func TestLoadConfigServer (t *testing.T){

	var para string = "../conf/server.conf"
	config := LoadConfigServer(para)
	t.Log(config)

}