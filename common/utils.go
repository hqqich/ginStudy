// Package common @Author hqqich 2022-11-22 16:23:00
package common

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"unicode"
)

// OpenBrowser @Title 打开浏览器
// @Description 根据系统命令，打开浏览器并访问地址
// @Author hqqich 2022-11-22 16:23:00
// @Param name
// @Return string
func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		log.Println(err)
	}
}

func GetIp() (ip string) {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return ip
	}

	for _, a := range ips {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				if strings.HasPrefix(ip, "10") {
					return
				}
				if strings.HasPrefix(ip, "172") {
					return
				}
				if strings.HasPrefix(ip, "192.168") {
					return
				}
				ip = ""
			}
		}
	}
	return
}

var sizeKB = 1024
var sizeMB = sizeKB * 1024
var sizeGB = sizeMB * 1024
var sizeTB = sizeGB * 1024

func Bytes2Size(num int64) string {
	numStr := ""
	unit := "B"
	if num/int64(sizeTB) > 1 {
		numStr = fmt.Sprintf("%f", float64(num)/float64(sizeTB))
		unit = "TB"
	} else if num/int64(sizeGB) > 1 {
		numStr = fmt.Sprintf("%f", float64(num)/float64(sizeGB))
		unit = "GB"
	} else if num/int64(sizeMB) > 1 {
		numStr = fmt.Sprintf("%f", float64(num)/float64(sizeMB))
		unit = "MB"
	} else if num/int64(sizeKB) > 1 {
		numStr = fmt.Sprintf("%f", float64(num)/float64(sizeKB))
		unit = "KB"
	} else {
		numStr = fmt.Sprintf("%d", num)
	}
	numStr = strings.Split(numStr, ".")[0]
	return numStr + " " + unit
}

func Interface2String(inter interface{}) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	case int:
		return fmt.Sprintf("%d", inter.(int))
	case float64:
		return fmt.Sprintf("%f", inter.(float64))
	}
	return "Not Implemented"
}

// SpecialLetters 如果存在特殊字符，直接在特殊字符前添加
/**
判断是否为字母： unicode.IsLetter(v)
判断是否为十进制数字： unicode.IsDigit(v)
判断是否为数字： unicode.IsNumber(v)
判断是否为空白符号： unicode.IsSpace(v)
判断是否为Unicode标点字符 :unicode.IsPunct(v)
判断是否为中文：unicode.Han(v)
*/
func SpecialLetters(letter rune) (bool, []rune) {
	if unicode.IsPunct(letter) || unicode.IsSymbol(letter) || unicode.Is(unicode.Han, letter) {
		var chars []rune
		chars = append(chars, '\\', letter)
		return true, chars
	}
	return false, nil
}

// getMd5Hight 获取奇葩的MD5加密。
//
// str string 被加密字符串
//
// string 加密结果
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
