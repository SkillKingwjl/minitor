package mod

import "minitor/vpn/openvpn"

type UploadData struct {
	CPU        float64              `json:"cpu""`
	MEM        float64              `json:"mem""`
	IP         string               `json:"ip""`
	ServerInfo openvpn.ServerStatus `json:"serverInfo"`
}
