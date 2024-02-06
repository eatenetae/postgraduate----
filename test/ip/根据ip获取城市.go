package main

import (
	"fmt"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"net"
)

func main() {
	ip := net.ParseIP("10.102.162.58")
	ip4 := ip.To4()
	fmt.Println(ip4)
	db, _ := geoip2db.NewGeoipDbByStatik()
	defer db.Close()
	record, _ := db.City(ip)
	fmt.Println(record)
}

func IsIntranetIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}
	return false
}
