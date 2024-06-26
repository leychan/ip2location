package ip2location

import (
	"fmt"
	"os"
	"path/filepath"
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
	ex, err := os.Executable() // 获取当前执行文件的路径
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	InitIp2regionSeacher(exPath + "/data/ip2location-db/ip2region.xdb")
	fmt.Println(GetIpInfo("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(GetIpInfo("1.202.118.195"))
	fmt.Println(GetIpInfo("89.185.30.208"))
	fmt.Println(GetIpInfo("103.88.176.0"))
}

func TestGetIpInfoInternational(t *testing.T) {
	ex, err := os.Executable() // 获取当前执行文件的路径
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	InitIp2locationDb(exPath + "/data/ip2location-db/lite-9.bin", exPath + "/data/ip2location-db/lite-9-ipv6.bin")
	fmt.Println(GetIpInfoInternational("2001:19f0:6001:5d0c:5400:4ff:feab:e52e"))
	fmt.Println(GetIpInfoInternational("1.202.118.195"))
	fmt.Println(GetIpInfoInternational("89.185.30.208"))
	fmt.Println(GetIpInfoInternational("103.88.176.1"))
}