package test_0804

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"testing"
	"time"
)

// TestAnswerClaude ip list for test from claude answer
func TestAnswerClaude(t *testing.T) {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	var ipList string
	// Generate 50 random IP addresses
	for i := 0; i < 50; i++ {
		octet1 := rand.Intn(255)
		octet2 := rand.Intn(255)
		octet3 := rand.Intn(255)
		octet4 := rand.Intn(255)

		ip := fmt.Sprintf("%d.%d.%d.%d\n", octet1, octet2, octet3, octet4)
		ipList += ip
	}
	fmt.Print(ipList)
}

// TestAnswerOne IP list for test from the first answer
func TestAnswerOne(t *testing.T) {
	ipList := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
	}

	formattedIpList := joinIpList(ipList)
	fmt.Println("Formatted IPv4 list:", formattedIpList)
}

func joinIpList(ipList []string) string {
	var formattedIpList []string
	for _, ip := range ipList {
		// formattedIp := ip + strings.Repeat(".") err
		formattedIp := ip + strings.Repeat(".", 0)
		formattedIpList = append(formattedIpList, formattedIp)
	}

	return strings.Join(formattedIpList, " ")
}

// TestAnswerTwo IP list for test from the second answer
func TestAnswerTwo(t *testing.T) {
	ipList := []net.IP{net.ParseIP("192.168.0.1"), net.ParseIP("10.0.0.1"), net.ParseIP("172.16.0.1")}
	ipStr := joinIPList(ipList)
	fmt.Println(ipStr)
}

func joinIPList(ipList []net.IP) string {
	ipStr := ""
	for _, ip := range ipList {
		ipStr += ip.String()
		ipStr += ","
	}
	if len(ipStr) > 0 {
		ipStr = ipStr[:len(ipStr)-1]
	}
	return ipStr
}

// TestAnswerThree IP list for test from the third answer
func TestAnswerThree(t *testing.T) {
	ipList := []string{"192.168.1.1", "10.0.0.1", "8.8.8.8"}

	// 使用net.IPNet结构体将IP地址转换为net.IP类型
	// var ipNets []net.IPNet error
	var ipNets []net.IP
	for _, ip := range ipList {
		ipNets = append(ipNets, net.ParseIP(ip))
	}

	// 使用for循环遍历IP网络并打印每个IP地址
	for _, ipNet := range ipNets {
		fmt.Printf("%s\n", ipNet.String())
	}
}

// TestAnswerFour IP list for test from the fourth answer
func TestAnswerFour(t *testing.T) {
	// 创建一个切片来存储 IP 地址
	ipAddresses := []string{}

	// 循环遍历一个 IP 地址范围
	for i := 1; i <= 255; i++ {
		// 创建一个 net.IP 对象，并将其转换为字符串
		ip := net.ParseIP(fmt.Sprintf("%d.%d.%d.%d", i, i, i, i))

		// 如果 IP 地址有效，则将其添加到切片中
		if ip != nil {
			ipAddresses = append(ipAddresses, ip.String())
		}
	}
	// 使用 Join 函数将切片中的 IP 地址连接成一个字符串并打印
	fmt.Println(strings.Join(ipAddresses, ", "))
}

// TestAnswerFive IP list for test from the fifth answer
func TestAnswerFive(t *testing.T) {
	ipList := []string{
		"192.168.0.1",
		"10.0.0.1",
		"8.8.8.8",
		"127.0.0.1",
	}
	// 使用 Join 函数将 IP 列表连接成一个字符串
	ipString := JoinIPList(ipList)
	// 输出结果
	fmt.Println(ipString)
}

// JoinIPList 函数接收一个字符串切片，并使用 net.JoinIPList 函数将它们连接成一个 IP 地址列表
func JoinIPList(ipList []string) string {
	ipString := ""
	for _, ip := range ipList {
		ipString += ip + ","
	}
	// 去掉最后一个逗号并返回结果
	return ipString[:len(ipString)-1]
}
