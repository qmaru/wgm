package common

import (
	"fmt"
	"net"
)

func isDomain(domain string) bool {
	_, err := net.LookupHost(domain)
	return err == nil
}

func PublicValidator(publicAddr string) (string, error) {
	if isDomain(publicAddr) {
		_, err := net.LookupHost(publicAddr)
		if err != nil {
			return "", fmt.Errorf("domain error")
		}
		return publicAddr, nil
	}
	ipNet := net.ParseIP(publicAddr)
	if ipNet == nil {
		return "", fmt.Errorf("ip error")
	}
	return publicAddr, nil
}

func PrivateValidator(privateAddr string) (string, error) {
	if isDomain(privateAddr) {
		addrs, err := net.LookupHost(privateAddr)
		if err != nil {
			return "", fmt.Errorf("domain error")
		}
		for _, addr := range addrs {
			return addr, nil
		}
	}
	ipNet := net.ParseIP(privateAddr)
	if ipNet == nil {
		return "", fmt.Errorf("ip error")
	}
	return privateAddr, nil
}

// CIDRValidator 检查IP块
func CIDRValidator(cidr string) (string, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}
	return ipNet.String(), nil
}

// IPCheck 检查 IP 是否在划分的范围内
func IPcheck(ip string, cidr string) bool {
	_, subnet, _ := net.ParseCIDR(cidr)
	nip := net.ParseIP(ip)
	return subnet.Contains(nip)
}
