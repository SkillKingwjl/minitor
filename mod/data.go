package mod

import "minitor/vpn/openvpn"

type UploadData struct {
	CPU        float64
	MEM        float64
	IP         string
	ServerInfo openvpn.ServerStatus
}
