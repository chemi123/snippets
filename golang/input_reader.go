package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
  以下の数字データをファイルから読み取って出力するスニペット。受けっとた段階ではstring扱いになるため、intとして扱うところまでやる。
  また、ファイルから入力を受け取った場合はスペースがある可能性もあるため、スペースを削除する処理も追加している。
  1
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dataSetNum, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	fmt.Println(dataSetNum)
}
