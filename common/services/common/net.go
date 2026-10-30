package common

import (
	"errors"
	"net"
)

func PublicValidator(publicAddr string) (string, error) {
	if net.ParseIP(publicAddr) != nil {
		return publicAddr, nil
	}
	if _, err := net.LookupHost(publicAddr); err != nil {
		return "", errors.New("ip error")
	}
	return publicAddr, nil
}

func PrivateValidator(privateAddr string) (string, error) {
	if net.ParseIP(privateAddr) != nil {
		return privateAddr, nil
	}
	addrs, err := net.LookupHost(privateAddr)
	if err != nil || len(addrs) == 0 {
		return "", errors.New("ip error")
	}
	return addrs[0], nil
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
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	nip := net.ParseIP(ip)
	return nip != nil && subnet.Contains(nip)
}
