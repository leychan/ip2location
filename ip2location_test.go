package ip2location

import (
	"fmt"
	"testing"
)

var ips = []string{
	"1.202.118.195",
	"111.55.11.252",
	"61.242.134.22",
	"120.197.61.20",
	"139.210.169.82",
	"117.61.86.188",
	"117.61.86.188",
	"39.144.109.62",
	"1.80.216.224",
	"117.136.74.235",
	"39.149.31.75",
	"39.149.31.75",
	"183.15.205.122",
	"183.15.205.122",
	"223.104.66.249",
	"171.81.215.125",
	"117.172.1.110",
	"14.151.33.8",
	"14.151.33.8",
	"171.43.29.150",
	"223.104.72.222",
	"223.104.72.149",
	"39.144.69.102",
	"223.104.199.42",
}

func BenchmarkGetIpInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetIpInfo(ips[i%len(ips)])
	}
}

func TestGetIpInfo(t *testing.T) {
	fmt.Println(GetIpInfo("62.210.214.190"))
	fmt.Println(GetIpInfo("1.202.118.195"))
	fmt.Println(GetIpInfo("89.185.30.208"))
}

func TestGetIpInfoInternational(t *testing.T) {
	fmt.Println(GetIpInfoInternational("62.210.214.190"))
	fmt.Println(GetIpInfoInternational("1.202.118.195"))
	fmt.Println(GetIpInfoInternational("89.185.30.208"))
}