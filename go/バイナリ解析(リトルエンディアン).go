package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// https://ascii.jp/elem/000/001/260/1260449/

// 現在主流のCPU 2 はリトルエンディアンです
// （サーバや組み込み機器で使用されるCPUにはビッグエンディアンのものもあります）。
// リトルエンディアンでは、10000という数値（16進表示で0x2710）があったときに、
// 小さい桁からメモリに格納されます（Go言語で書けば []byte{0x10, 0x27, 0x0, 0x0} と表現されます）。

func main() {
	// 32bit
	b4 := []byte{0x10, 0x27, 0x0, 0x0} // 4バイト（8bit×4=32ビット）で構成。0xは16進数を表す
	var i4 int32
	buf4 := bytes.NewReader(b4)
	// 0x2710 = 2*16*16*16 + 7*16*16 + 1*16 + 0
	_ = binary.Read(buf4, binary.LittleEndian, &i4)
	fmt.Println(i4)
	// 64bit
	b8 := []byte{0x10, 0x27, 0x0, 0x0, 0x10, 0x27, 0x0, 0x0} // 8バイト（8bit×8=64ビット）で構成。0xは16進数を表す
	var i8 int64
	buf8 := bytes.NewReader(b8)
	// 0x271000002710
	_ = binary.Read(buf8, binary.LittleEndian, &i8)
	fmt.Println(i8)
}
