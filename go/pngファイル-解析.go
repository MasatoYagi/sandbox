package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// https://ascii.jp/elem/000/001/260/1260449/#fn3
// PNGファイルを分析してみる
// PNGファイルはバイナリフォーマットです。
// 先頭の8バイトがシグニチャ（固定のバイト列）となっています。
// それ以降は次のようなチャンク（データの塊）のブロックで構成されています。

// 各チャンクとその長さを列挙してみましょう。
// 以下のコードでは、 readChunks() 関数でチャンクごとに io.SectionReader を作って配列に格納して返しています。
// それをチャンクを表示する関数（dumpChunk()）で表示しています。
// サンプルとして利用しているPNG画像は、コンピュータグラフィックス業界でこれ以上の有名人はいないという、
// レナ・ソーダバーグさん 3 の画像をお借りしました。
// ・https://en.wikipedia.org/wiki/File:Lenna.png

func dumpChunk(chunk io.Reader) {
	var length int32                              // pngはデータ以外は4バイトなので32bit(4byte*8bit)を利用
	fmt.Printf("%+v\n", chunk)                    // =>&{r:0xc0000ae018 base:8 off:8 limit:33 n:25}
	binary.Read(chunk, binary.BigEndian, &length) // 先頭の4バイト（長さのデータ）を読み取り、ビッグエンディアン形式で `length` に格納
	fmt.Printf("%+v\n", chunk)                    // =>&{r:0xc0000ae018 base:8 off:12 limit:33 n:25}
	buffer := make([]byte, 4)                     // binary.Read()で読み取り位置が4バイト分進んでいる（↑のoffを参照）
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader

	// 最初の8バイト(シグネチャ)を飛ばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32                                    // pngはデータ以外は4バイトなので32bit(4byte*8bit)を利用
		err := binary.Read(file, binary.BigEndian, &length) // 先頭の4バイト（長さのデータ）を読み取り、ビッグエンディアン形式で `length` に格納
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12)) // 次のチャンクの頭から読み取り：12 = 長さ(4byte)+種類(4byte)+CRC(4byte)
		// 次のチャンクの先頭に移動
		// 現在位置は長さを読み終わった箇所なので
		// チャンク名(たぶんチャンクの種類,4バイト) + データ長 + CRC(4バイト)先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func main() {
	file, _ := os.Open("pngファイル-解析.png")
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
