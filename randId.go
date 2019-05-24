package main

import (
	"crypto/rand"
	"fmt"
)

func GenerateId() string {
	b := make([]byte, 16)       //16 байт
	rand.Read(b)                //рандомно их заполняем
	return fmt.Sprintf("%x", b) //выводим их в строку
}
