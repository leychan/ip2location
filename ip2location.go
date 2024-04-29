package ip2location

import (
	"errors"
	"net"
	"strings"
	"sync"

	"github.com/ip2location/ip2location-go/v9"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var (
	ip2locationDB         *ip2location.DB //ipv4 db
	ip2locationIpv6DB     *ip2location.DB //ipv6 db
	initIp2locationDbOnce sync.Once // db只初始化1次 db is only initialized once
)

var (
	ip2regionSeacher         *xdb.Searcher
	initIp2regionSeacherOnce sync.Once //searcher只初始化1次 searcher is only initialized once
)

var initError error //初始化错误

// 初始化ip2location数据库 默认初始化ipv4数据库, 如果传入ipv6数据库文件地址,则同时初始化ipv6数据库
func InitIp2locationDb(path string, pathIpv6 string) {
	initIp2locationDbOnce.Do(func() {
		ip2locationDB, initError = ip2location.OpenDB(path)
		if initError != nil {
			panic(initError)
		}

		if pathIpv6 != "" {
			ip2locationIpv6DB, initError = ip2location.OpenDB(pathIpv6)
			if initError != nil {
				panic(initError)
			}
		}
	})
}

// 国内ip地址信息查询 初始化ip2region数据库, 仅支持ipv4
func InitIp2regionSeacher(path string) {
	initIp2regionSeacherOnce.Do(func() {
		cBuffer, initError = xdb.LoadContentFromFile(path)
		if initError != nil {
			panic(initError)
		}
		ip2regionSeacher, initError = xdb.NewWithBuffer(cBuffer)

		if initError != nil {
			panic(initError)
		}
	})
}

// 获取国际ip地址信息
func GetIpInfoInternational(ip string) (IPInfo, error) {
	ipParsed := net.ParseIP(ip)
	if ipParsed == nil {
		return IPInfo{}, errors.New("invalid ip")
	}

	var record ip2location.IP2Locationrecord
	var err error
	
	// ipv4,从ipv4数据库中获取
	if ipParsed.To4() != nil {
		record, err = ip2locationDB.Get_all(ip)
	} else { // ipv6 从ipv6数据库中获取
		record, err = ip2locationIpv6DB.Get_all(ip)
	}

	if err != nil {
		return IPInfo{}, err
	}

	return IPInfo{
		Country: record.Country_short,
		Region:  record.Region,
		Area:    "",
		City:    record.City,
		Isp:     "",
	}, nil
}

type IPInfo struct {
	Country string
	Region  string
	Area    string
	City    string
	Isp     string
}

// 获取国内ip地址信息
func GetIpInfo(ip string) (IPInfo, error) {
	var ipInfo IPInfo
	res, err := TransIP2RegionStrOffline(ip)
	if err != nil {
		return ipInfo, err
	}

	arr := strings.Split(res, "|")
	if len(arr) < 5 {
		return ipInfo, errors.New("res len < 5")
	}
	// 中国|0|北京|北京市|电信
	ipInfo = IPInfo{
		Country: arr[0],
		Region:  arr[2],
		Area:    "",
		City:    arr[3],
		Isp:     arr[4],
	}

	return ipInfo, nil
}

var cBuffer []byte

func TransIP2RegionStrOffline(ip string) (string, error) {
	defer ip2regionSeacher.Close()
	// 国家|区域|省份|城市|ISP
	res, err := ip2regionSeacher.SearchByStr(ip)
	if err != nil {
		return "", err
	}
	return res, nil
}
