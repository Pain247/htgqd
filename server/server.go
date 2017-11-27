package server

import (
	"github.com/NYTimes/gziphandler"
	"net/http"
	"time"
	"fmt"
	"github.com/tuyensinh/utils"
	"strconv"
	"github.com/tuyensinh/topsis/load_data"
	"github.com/tuyensinh/topsis/process"
	"encoding/json"
)
func (server *Server) InitMux() {
	server.mux = make(map[string]func(http.ResponseWriter, *http.Request))
	server.mux["/api/result"] = server.Topsis
	server.mux["/api/hobby"] = server.Hobbies
	server.mux["/api/group"] = server.Group
}
func (server *Server) InitServer() {
	server.config = utils.LoadConfigServer(PATH_SERVER_CONFIG)
	server.InitMux()
	withoutGz := server
	withGz := gziphandler.GzipHandler(withoutGz)
	server.server = &http.Server{
		Addr:         server.config.ADDR,
		Handler:      withGz,
		ReadTimeout:  500 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (server *Server) StartServer() {
	fmt.Println("SSP server is listening ...")
	server.server.ListenAndServe()
}
func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := server.mux[r.URL.Path]; ok {
		h(w, r)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("<b><font size=\"6\">Bad request</font></b>"))
}

func (server *Server) Topsis(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	score,_ := strconv.ParseFloat(r.Form["score"][0],64)
	hobbies,_ := strconv.Atoi(r.Form["hobby"][0])
	area,_ := strconv.Atoi(r.Form["area"][0])
	group := string(r.Form["group"][0])
	input := load_data.LoadDepartment()
	convert := process.ConvertData(group,score,area,hobbies,input)
	arr := process.NormalizeData(convert)
	mapsub := process.GetAsub(arr)
	mapstar := process.GetAstar(arr)
	sstar := process.GetSstar(arr,mapstar)
	ssub := process.GetSstar(arr,mapsub)
	agg := process.Aggregate(sstar,ssub)
	list := process.GetList(agg)
	results := make([]map[string]interface{},0)
	for i:=0;i<len(list);i++{
		results = append(results,map[string]interface{}{"sst": list[i]["stt"], "data": load_data.GetInfoDep(list[i]["id"].(int))})
	}
	res, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (server *Server) Hobbies(w http.ResponseWriter, r *http.Request){
        results := load_data.LoadHobbies()
	res,_ := json.Marshal(results)
	w.Write(res)
}
func (server *Server) Group(w http.ResponseWriter, r *http.Request){
	results := load_data.LoadGroup()
	res,_ := json.Marshal(results)
	w.Write(res)
}