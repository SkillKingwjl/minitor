package openvpn

type ConnectedClient struct {
	CommonName         string `json:"name"`
	RealAddress        string `json:"realAddress"`
	VirtualAddress     string `json:"virtualAddress"`
	VirtualIPv6Address string `json:"virtualIpv6Address"`
	BytesRX            int    `json:"bytesRx"`
	BytesTX            int    `json:"bytesTx"`
	ConnectedSince     int    `json:"connectedSince"`
	Username           string `json:"username"`
	ClientID           int    `json:"clientId"`
	PeerID             int    `json:"peerId"`
}

type ServerStatus struct {
	Version  string            `json:"version"`
	Time     int               `json:"time"`
	Client   string            `json:"client"`
	BytesIn  string            `json:"bytesIn"`
	BytesOut string            `json:"bytesOut"`
	Clients  []ConnectedClient `json:"clients"`
}
