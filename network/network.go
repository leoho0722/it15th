package network

import (
	"fmt"
	"log"
	"net"
)

// GetLocalNetworkIPAddress 取得電腦位於區網的 IP
//
//	Returns:
//	 ip: 電腦位於區網的 IP
func GetLocalNetworkIPAddress() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println("IP：", localAddr.IP)
	return localAddr.IP.String()
}
