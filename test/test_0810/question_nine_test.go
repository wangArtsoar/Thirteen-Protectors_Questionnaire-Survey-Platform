package test_0810

import (
	"fmt"
	"strings"
	"testing"
)

func TestAnswerByClaude(t *testing.T) {
	var ips []string
	for i := 1; i <= 255; i++ {
		for j := 1; j <= 255; j++ {
			ips = append(ips, fmt.Sprintf("10.0.%d.%d", i, j))
		}
	}
	fmt.Println(strings.Join(ips, "\n"))
}

func TestAnswerByFirst(t *testing.T) {
	// 定义一个包含 IP 地址的切片
	ips := []string{
		"192.168.0.1",
		"192.168.0.2",
		"192.168.0.3",
		"192.168.0.4",
	}
	// 遍历切片并打印每个 IP 地址
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func TestAnswerBySecond(t *testing.T) {
	// 首先定义一个字符串数组，用于存储 IP 地址
	ipList := []string{"192.168.1.1", "8.8.8.8", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "1.1.1.1"}

	// 遍历数组，将每个 IP 地址转换为一个字符串并输出
	for _, ip := range ipList {
		fmt.Println(ip)
	}
}
