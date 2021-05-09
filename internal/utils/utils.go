package utils

import (
	"net"
	"time"
	"math"
	"math/rand"
	"strings"
)

func RandStr(n int) string {
	rand.Seed(time.Now().UnixNano())

	letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func IsDirPath(s string) (bool) {
	if string(s[len(s)-1]) == "/" {
		return true
	}

	return false
}

func Basename(s string) (string){
	splits := strings.Split(s, "/")

	if IsDirPath(s) {
		return splits[len(splits)-2]
	}
	
	return splits[len(splits)-1]
}

func ParseSize(s int64) (float64, string){
	suffix := []string{"B","KB","MB","GB","TB"}
	var i int

	val := float64(s)
	for i=0;i<len(suffix);i++ {
		if val < 1024 {
			break
		}
		val = val / 1024
	}

	val = math.Round(val*100)/100

	return val, suffix[i]
}

func GetIP(host string) (string) {
	if host != "0.0.0.0" {
		return host
	}

    netInterfaceAddresses, _ := net.InterfaceAddrs()

    for _, netInterfaceAddress := range netInterfaceAddresses {
        networkIp, ok := netInterfaceAddress.(*net.IPNet)
        if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
            return networkIp.IP.String()
        }
    }

    return "IP"
}