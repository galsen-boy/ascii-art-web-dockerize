package main

import (
	"fmt"
	"os"
	"strings"
)

func ConvertStr(str, banner string) string {

	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	dataSplit := strings.Split(string(data), "\n")
	strTab := strings.Split(str, "\r\n")
	var newStr string
	for j, args := range strTab {
		if args != "" {
			var result []string
			for _, char := range args {
				for i := 1; i <= 8; i++ {
					result = append(result, dataSplit[((char-32)*9)+rune(i)])
				}
			}
			var tab [8][]string
			for i, val := range result {
				tab[i%8] = append(tab[i%8], val)
			}
			for _, ligne := range tab {
				for _, part := range ligne {
					newStr += part
				}
				newStr += "\n"

			}
		} else if j != len(str)-1 && args != "\n" {
			newStr += "\n"
		}
	}
	// fmt.Println(newStr)
	return newStr
}
func isValid(s string) bool {
	for _, ch := range s {
		if ch < ' ' && ch != 10 && ch != '\r' || ch > '~' {
			return false
		}
	}
	return true
}
