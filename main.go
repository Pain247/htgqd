package main

import (
	"github.com/tuyensinh/server"
)
func NewServer() *server.Server{
	return &server.Server{}
}
func main(){
	server := NewServer()
	server.InitServer()
	server.StartServer()

}