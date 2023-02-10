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
	result += fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?><VanDyke version="3.0"><key name="Sessions"><key name="%s">`, time.Now().Format("2006-01-02"))

	obj := make(map[string][]string)
	for _, line := range inputLineList {
		data := strings.Split(line, ",")
		folderName := strings.TrimSpace(data[0])
		serverName := strings.TrimSpace(data[1])
		obj[folderName] = append(obj[folderName], serverName)
	}

	for key, serverList := range obj {
		result += fmt.Sprintf(`<key name="%s">`, key)
		for _, server := range serverList {
			result += fmt.Sprintf(`<key name="%s"><string name="Hostname">%s</string><dword name="[SSH2] Port">%d</dword><string name="Username">%s</string></key>`, server+" "+key, server, PORT, USER)
		}
		result += `</key>`
	}
	result += `</key></key></VanDyke>`
	err = ioutil.WriteFile("output.xml", []byte(result), 0644)
	if err == nil {
		fmt.Println("make done")
	}
}
