package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// https://ascii.jp/elem/000/001/260/1260449/

// ネットワーク上で転送されるデータの多くは、大きい桁からメモリに格納されるビッグエンディアン（ネットワークバイトオーダーとも呼ばれます）です。 そのため多くの環境では、
// ネットワークで受け取ったデータをリトルエンディアンに修正する必要があるのです。
// 任意のエンディアンの数値を、現在の実行環境のエンディアンの数値に修正するには、
// encoding/binaryパッケージを使います。 このパッケージの binary.Read() メソッドに、
// io.Reader とデータのエンディアン、それに変換結果を格納する変数のポインタを渡せば、
// エンディアンが修正されたデータが得られます。

func main() {
	// 32ビットのビッグエンディアンのデータ（10000）
	data := []byte{0x0, 0x0, 0x27, 0x10} // 0x2710
	var i int32
	// エンディアンの変換(0x1027)
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i) // Intel CPU上ではリトルエンディアンのため、ビッグ→リトルに変換してデータをiに格納
	fmt.Println(data)
	fmt.Printf("data: %d\n", i)
}
