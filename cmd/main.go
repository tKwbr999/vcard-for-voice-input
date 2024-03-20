package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	list := `GPT,じぴーてぃー
ChatGPT,ちゃっとじぴーてぃー
Chat,ちゃっと`

	lines := strings.Split(list, "\n")
	r := regexp.MustCompile(`(.*),(.*)`)

	// VCARDファイルを作成
	file, err := os.Create("voiceInput.vcf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		if len(matches) == 3 {
			name := matches[1]
			yomi := matches[2]
			vcard := fmt.Sprintf(`BEGIN:VCARD
VERSION:3.0
PRODID:-//Apple Inc.//macOS 14.2.1//EN
N:%s;;;;
FN:%s
X-PHONETIC-LAST-NAME:%s
END:VCARD`, name, name, yomi)
			// VCARDをファイルに書き込む
			_, err := file.WriteString(vcard)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
	fmt.Println("VCARDファイルが作成されました。")
}
