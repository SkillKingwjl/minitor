package main

import (
	"encoding/json"
	"flag"
	"log"
	"minitor/mod"
	"minitor/os/info"
	_ "minitor/os/info"
	"minitor/vpn"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	url string
	ti  int64
)

func init() {
	flag.StringVar(&url, "url", "https://api.transwarpv.xyz/servers/monitor", "data reporting address")
	flag.Int64Var(&ti, "t", 10, "upload interval")
}

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	go Run(url, ti)
	wg.Wait()
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
	data.CPU = info.GetCpuPercent()
	data.MEM = info.GetMemPercent()
	ip := localAddr.IP.String()
	//ip := "31.13.213.52"
	port := "17562"
	ipStr := ip + ":" + port
	//data.IP = localAddr.IP.String()
	data.IP = ip
	loadStateHandler := vpn.LoadStateHandler(ipStr)
	data.ServerInfo = vpn.StatusHandler(ipStr)
	data.ServerInfo.Client = loadStateHandler.Client
	data.ServerInfo.BytesIn = loadStateHandler.BytesIn
	data.ServerInfo.BytesOut = loadStateHandler.BytesOut
	marshal, err := json.Marshal(data)
	dataStr := string(marshal)
	log.Println(dataStr)
	return dataStr
}

func Run(url string, ti int64) {
	duration := time.Duration(60 * 1000 * 1000 * 1000 * ti)
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			data := MakeData()
			resp, err := http.Post(url, "application/json;charset=utf-8", strings.NewReader(data))
			if err != nil {
				log.Printf("post请求失败 error: %+v", err)
				continue
			}
			defer resp.Body.Close()
		}
	}

}
