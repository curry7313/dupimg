package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const InputDir = ""

func main() {
	cmd := exec.Command("findimagedupes", "-t", "10", InputDir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	fmt.Println()
	fmt.Println()
	rows := strings.Split(string(out), "\n")
	for idx, row := range rows {
		r := strings.TrimSpace(row)
		if r == "" {
			continue
		}
		if idx == 0 {
			rowStr := strings.TrimSpace(TrimHiddenCharacter(row))
			temp := strings.Split(rowStr, "  ")
			r = strings.TrimSpace(temp[len(temp)-1])
		}
		tmp := strings.Split(r, " ")
		for i := 1; i < len(tmp); i++ {
			filePath := strings.TrimSpace(tmp[i])
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println("remove err", err)
				continue
			}
		}
	}
}

func TrimHiddenCharacter(originStr string) string {
	srcRunes := []rune(originStr)
	dstRunes := make([]rune, 0, len(srcRunes))
	for _, c := range srcRunes {
		if c >= 0 && c <= 31 {
			continue
		}
		if c == 127 {
			continue
		}
		dstRunes = append(dstRunes, c)
	}
	return string(dstRunes)
}
