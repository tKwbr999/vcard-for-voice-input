package main

import (
	"fmt"
	"os"
)

func main() {
	list := map[string]string{
		"GPT":     "じぴーてぃー",
		"ChatGPT": "ちゃっとじぴーてぃー",
		"Chat":    "ちゃっと",
		"ruby":    "るびー",
		"Go言語":    "ごーげんご",
	}

	// VCARDファイルを作成
	file, err := os.Create("voiceInput.vcf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for k, v := range list {
		vcard := fmt.Sprintf(`BEGIN:VCARD
VERSION:3.0
PRODID:-//Apple Inc.//macOS 14.2.1//EN
N:%s;;;;
FN:%s
X-PHONETIC-LAST-NAME:%s
END:VCARD`, k, k, v)
		// VCARDをファイルに書き込む
		_, err := file.WriteString(vcard)
		if err != nil {
			panic(err)
		}
	}
}
