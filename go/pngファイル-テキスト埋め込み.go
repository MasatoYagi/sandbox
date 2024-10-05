package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

// https://ascii.jp/elem/000/001/260/1260449/#fn3
// ・https://en.wikipedia.org/wiki/File:Lenna.png

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData))) // データ長をバッファに書き込む(チャンク内の1要素目)
	buffer.WriteString("tEXt")                                    // 種類(チャンク名)を書き込む(チャンク内の2要素目)
	buffer.Write(byteData)                                        // データを書き込む(チャンク内の3要素目)
	// CRCを計算して追加
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32()) // CRCを書き込む((チャンク内の4要素目))
	return &buffer
}

// チャンク毎のio.Readerのスライスを返す
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
	defer file.Close()
	newFile, _ := os.Create("pngファイル-解析-new.png")
	defer newFile.Close()
	chunks := readChunks(file)
	// シグニチャ書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// 先頭に必要なIHDRチャンクを書き込み
	io.Copy(newFile, chunks[0])
	// テキストチャンクを追加
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	// pngのtEXtに書き込まれたデータを文字列にして出力
	fileNew, _ := os.Open("pngファイル-解析-new.png")
	chunks = readChunks(fileNew)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length) // データ長(チャンク内の1要素目)をlengthに入れる
	buffer := make([]byte, 4)
	chunk.Read(buffer) // 種類(チャンク名)をバッファにいれる(チャンク内の2要素目)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}
