package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	device      = "en0"                      // 替换为你的网络接口
	targetIP    = "192.168.1.8"              // 目标 IP
	port        = "3000"                     // 目标端口
	filter      = "tcp and dst port " + port // 抓包过滤器
	snapshotLen = int32(1600)                // 抓包的最大长度
	promiscuous = false                      // 是否启用混杂模式
)

// 维护 TCP 流数据缓存
var tcpStreamData = make(map[string]*bytes.Buffer)

func main() {
	// 打开网络设备
	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// 设置过滤器
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("正在抓取 192.168.1.8:3000 上包含 'save' 的 POST 表单请求...")

	// 开始抓取包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		processPacket(packet)
	}
}

// 处理每个包
func processPacket(packet gopacket.Packet) {
	// 提取 IPv4 层
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return // 跳过非 IPv4 包
	}
	ip, _ := ipLayer.(*layers.IPv4)
	if ip.DstIP.String() != targetIP {
		return // 跳过不符合目标 IP 的包
	}

	// 提取 TCP 层
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return // 跳过非 TCP 包
	}
	tcp, _ := tcpLayer.(*layers.TCP)

	// 构建流的唯一标识符
	streamID := fmt.Sprintf("%s:%d -> %s:%d", ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort)

	// 初始化或更新流的缓冲区
	if _, found := tcpStreamData[streamID]; !found {
		tcpStreamData[streamID] = &bytes.Buffer{}
	}
	tcpStreamData[streamID].Write(tcp.Payload)

	// 检查是否包含完整的 HTTP 请求
	if bytes.Contains(tcpStreamData[streamID].Bytes(), []byte("\r\n\r\n")) {
		// 提取表单数据
		formData := extractFormData(tcpStreamData[streamID].String())
		if len(formData) > 0 {
			fmt.Println("抓到符合条件的请求，表单数据如下:")
			for key, value := range formData {
				fmt.Printf("%s: %s\n", key, value)
			}
		}

		// 处理完成后清除该流的缓存
		delete(tcpStreamData, streamID)
	}
}

// 提取表单数据
func extractFormData(payload string) map[string]string {
	formData := make(map[string]string)
	// 查找 Content-Type 和 form data 的位置
	if strings.Contains(payload, "Content-Type: application/x-www-form-urlencoded") {
		// HTTP 请求内容在 \r\n\r\n 后出现
		parts := strings.Split(payload, "\r\n\r\n")
		if len(parts) > 1 {
			formBody := parts[1]
			// 解析 URL 编码的表单数据
			values, err := url.ParseQuery(formBody)
			if err == nil {
				for key, val := range values {
					formData[key] = val[0]
				}
			}
		}
	}
	return formData
}
