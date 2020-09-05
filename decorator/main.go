package main

import (
	"fmt"

	"github.com/mytkdals93/goweb/decorator/cipher"
	"github.com/mytkdals93/goweb/decorator/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var recvData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	//Send data
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(d string) {
	//Zip data
	data, err := lzw.Write([]byte(d))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(data))
}

type EncryptCompnent struct {
	key string
	com Component
}

func (self *EncryptCompnent) Operator(d string) {
	//Encrypt
	data, err := cipher.Encrypt([]byte(d), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(data))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(d string) {
	//Decerypt
	data, err := cipher.Decrypt([]byte(d), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(data))
}

type UnzipComponent struct {
	key string
	com Component
}

func (self *UnzipComponent) Operator(d string) {
	//Decerypt
	data, err := lzw.Read([]byte(d))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(data))
}

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	//Send data
	recvData = data
}
func main() {
	sender := &EncryptCompnent{key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		}}
	sender.Operator("Hello World")
	fmt.Println(sentData)
	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		}}
	receiver.Operator(sentData)
	fmt.Println(recvData)
}
