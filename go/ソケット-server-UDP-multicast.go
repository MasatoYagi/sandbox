package main

import (
	"fmt"
	"net"
	"time"
)

// https://ascii.jp/elem/000/001/411/1411547/
// UDPソケットのマルチキャスト

// マルチキャストでは使える宛先IPアドレスがあらかじめ決められていて、
// ある送信元から同じマルチキャストアドレスに属するコンピュータに対してデータを配信できます。
// 送信元とマルチキャストアドレスの組み合わせをグループといい、同じグループであれば、受信するコンピュータが100台でも送信側の負担は1台ぶんです。
// IPv4については、先頭4ビットが1110のアドレス（224.0.0.0 ～ 239.255.255.255）がマルチキャスト用として予約されています2。
// IPv6については、先頭8ビットが11111111のアドレスがマルチキャスト用アドレスです。
// IPv4では224.0.0.0 ～ 224.0.0.255の範囲がローカル用として予約されているので、このアドレスはテストなどで使えます。

// 1対多通信を想定した例
func main() {
	fmt.Println("Start tick server at 224.0.0.1:9999")
	conn, err := net.Dial("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	start := time.Now()
	wait := 10*time.Second - time.Nanosecond*time.Duration(start.UnixNano()%(10*1000*1000*1000))
	time.Sleep(wait)
	ticker := time.Tick(10 * time.Second)
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}
