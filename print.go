package utils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	bold_print_placeholder = "="
)

func BoldPrint(format string, s ...interface{}) {
	msg := fmt.Sprintf(format, s...)
	lines := strings.Split(msg, "\n")

	var lineMaxLen int
	for i, line := range lines {
		if l := utf8.RuneCountInString(line); l > lineMaxLen {
			lineMaxLen = l
		}
		lines[i] = "| " + line + "\n"
	}
	var placeholder string
	for i := 0; i < lineMaxLen; i++ {
		placeholder += bold_print_placeholder
	}
	fmt.Print(
		"==" + placeholder + "==\n" +
			strings.Join(lines, "") +
			"==" + placeholder + "==\n",
	)
}
