package playstring

import (
	"fmt"
	"regexp"
	"strings"
)

func PrintString(name string) {
	for _, v := range name {
		fmt.Printf("%c \n", v)
	}
}

func ChangeString(nameByte []rune, index int) string {
	nameByte[index] = 'a'
	return string(nameByte)
}

func CheckforJira(key string) {
	KEY := strings.ToUpper(key)
	fmt.Printf("KEY is %s  and fist is %c \n", KEY, KEY[0])
	match, _ := regexp.MatchString("^[0-9]", KEY)
	if match {
		fmt.Printf("TURE match is %t\n", match)
	} else {
		fmt.Printf("FALSE match is %t\n", match)
	}
}
