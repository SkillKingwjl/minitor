package main

import (
	"encoding/json"
	"fmt"
	"log"
	"minitor/mod"
	"minitor/os"
	_ "minitor/os"
	"minitor/vpn"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	go Run()
	for {

	}
}

func MakeData() string {
	//fmt.Println(os.GetCpuPercent())
	//fmt.Println(os.GetMemPercent())
	//fmt.Println(os.GetDiskPercent())
	//fmt.Println(os.GetCpuInfo())
	//fmt.Println(os.GetIOInfo())
	//fmt.Println(os.GetNetInfo("upd4"))
	//fmt.Println(vpn.StatusHandler("31.13.213.236:17562"))
	conn, err := net.Dial("udp", "8.8.8.8:80")
	var data mod.UploadData
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.IP.String())
	data.CPU = os.GetCpuPercent()
	data.MEM = os.GetMemPercent()
	data.IP = localAddr.IP.String()
	ip := localAddr.IP.String() + ":" + "17562"
	loadStateHandler := vpn.LoadStateHandler(ip)
	data.ServerInfo = vpn.StatusHandler(ip)
	data.ServerInfo.Client = loadStateHandler.Client
	data.ServerInfo.BytesIn = loadStateHandler.BytesIn
	data.ServerInfo.BytesOut = loadStateHandler.BytesOut
	marshal, err := json.Marshal(data)
	dataStr := string(marshal)
	log.Println(dataStr)
	return dataStr
}

func Run() {
	ticker := time.NewTicker(time.Minute * 1)
	for {
		select {
		case <-ticker.C:
			data := MakeData()
			resp, err := http.Post("http://127.0.0.1:7788/servers/monitor", "application/json;charset=utf-8", strings.NewReader(data))
			if err != nil {
				log.Printf("post请求失败 error: %+v", err)
				return
			}
			defer resp.Body.Close()
		}
	}

}
