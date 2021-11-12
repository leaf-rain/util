package tool

import (
	"net"
	"net/http"
	"strings"
)

func GetClientIP(req *http.Request) string {
	var ips []string
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		ips = strings.Split(ip, ",")
	}
	if len(ips) > 0 && ips[0] != "" {
		rip, _, err := net.SplitHostPort(ips[0])
		if err != nil {
			rip = ips[0]
		}
		return rip
	}
	if ip, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		return ip
	}
	return req.RemoteAddr
}
