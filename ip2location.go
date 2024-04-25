package ip2location

import (
	"errors"
	"strings"
	"sync"

	"github.com/ip2location/ip2location-go/v9"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var ip2locationDB *ip2location.DB
var initIp2locationDbOnce sync.Once

var ip2regionSeacher *xdb.Searcher
var initIp2regionSeacherOnce sync.Once

var initError error

func init() {
    initIp2locationDbOnce.Do(func() {
        ip2locationDB, initError = ip2location.OpenDB("db/lite-9.bin")
    })

    if initError != nil {
        panic(initError)
    }
    initIp2regionSeacherOnce.Do(func() {
        cBuffer, _ = xdb.LoadContentFromFile("db/ip2region.xdb")
	    ip2regionSeacher, initError = xdb.NewWithBuffer(cBuffer)
    })

    if initError != nil {
        panic(initError)
    }

}

func GetIpInfoInternational(ip string) (IPInfo, error) {
    record, err := ip2locationDB.Get_all(ip)
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
	//中国|0|北京|北京市|电信
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
	//国家|区域|省份|城市|ISP
	res, err := ip2regionSeacher.SearchByStr(ip)
	if err != nil {
		return "", err
	}
	return res, nil
}
