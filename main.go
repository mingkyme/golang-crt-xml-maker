package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var USER = "root"
var PORT = 22

func main() {

	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("input.txt 파일 읽기 실패")
	}
	inputLineList := strings.Split(string(inputBytes), "\n")
	result := ""
	result += fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?><VanDyke version="3.0"><key name="Sessions"><key name="%s">`, time.Now().Format("2006-01-02 15:04:05"))
	for _, l := range inputLineList {
		result += fmt.Sprintf(`<key name="%s"><string name="Hostname">%s</string><dword name="[SSH2] Port">%d</dword><string name="Username">%s</string></key>`, l, l, PORT, USER)
	}
	result += `</key></key></VanDyke>`
	err = ioutil.WriteFile("output.xml", []byte(result), 0644)
	if err == nil {
		fmt.Println("make done")
	}
}
