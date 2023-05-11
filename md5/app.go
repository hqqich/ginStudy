package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {

	fmt.Println(getMd5Hight("test"))

}

func getMd5Hight(str string) string {

	hexDigits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

	bytes := []byte(str)

	// md5 一次
	var hash = md5.Sum(bytes)

	var datas []string

	for i := range hash {
		b := hash[i]
		i1 := int64(b >> 4)
		i2 := int64(b & 0x0F)

		datas = append(datas, hexDigits[i1])
		datas = append(datas, hexDigits[i2])
	}

	join := strings.Join(datas, "")

	bytes2 := []byte(join)
	var datas2 []string
	for i := range bytes2 {

		hexStr := fmt.Sprintf("%02X", bytes2[i])
		datas2 = append(datas2, string(hexStr))

	}

	join2 := strings.Join(datas2, "")

	return join2
}
