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
)

func main() {
	fmt.Println(MakeData())
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
	data.ServerInfo = vpn.StatusHandler("31.13.213.236:17562")
	marshal, err := json.Marshal(data)
	dataStr := string(marshal)
	return dataStr
}
