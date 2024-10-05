package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// https://ascii.jp/elem/000/001/260/1260449/2/

// バイナリ解析の次はテキスト解析です。
// バイナリ解析の場合は、読み込むバイト数が固定であったり、
// 可変長データの場合も読み込むバイト数や個数などが事前に明示されていることがほとんどです。
// 一方、テキスト解析ではデータ長が決まらず、スキャンしながら区切りを探すしかありません。
// そのため、探索しながら読み込んでいく必要があります。

// 改行／単語で区切る
// テキスト解析の基本は改行区切りです。
// 全部読み込んでしまってから文字列処理で改行に分割する、という方法もありますが、
// io.Reader による入力では bufio.Reader を使うという手があり、そちらのほうが比較的シンプルです。
// ReadString()、ReadBytes() を使うと、任意の文字で分割することもできます。

var source = `1行目
2行目
3行目`

func main() {
	// 任意の文字で分割
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%#v\n", line)
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", line)
	}
	// 出力
	// "1行目\n"
	// "2行目\n"
	// "3行目"

	// 終端を気にせずにもっと短く書きたいのでれば、 bufio.Scanner を使う方法もあります。
	// これを使うと、上記のコードの main() 関数がこんなに短く書けます。
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
	// 出力
	// "1行目"
	// "2行目"
	// "3行目"
}
