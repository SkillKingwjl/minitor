package vpn

import (
	"context"
	"minitor/vpn/openvpn"
	_ "minitor/vpn/openvpn"
	"time"
)

func StatusHandler(serverAddr string) openvpn.ServerStatus {
	var s openvpn.ServerStatus
	server := openvpn.Server{
		Address: serverAddr,
	}
	if err := server.Connect(); err != nil {
		return s
	}
	defer server.Close()

	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	status, err := server.RequestStatus(c)
	if err != nil {
		return status
	}
	return status
}
