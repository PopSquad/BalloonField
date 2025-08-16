package util

import (
	"encoding/binary"
	"fmt"
	"net"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Str2GBK(s string) ([]byte, error) {
	utf8Bytes := []byte(s)
	encoder := simplifiedchinese.GBK.NewEncoder()
	gbkBytes, _, err := transform.Bytes(encoder, utf8Bytes)
	return gbkBytes, err
}

func IPv4ToUint32(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, fmt.Errorf("invalid ip: %s", ipStr)
	}

	ip = ip.To4()
	if ip == nil {
		return 0, fmt.Errorf("not an IPv4 address: %s", ipStr)
	}

	return binary.LittleEndian.Uint32(ip), nil
}
