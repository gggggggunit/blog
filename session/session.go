package session

import (
	"fmt"
	"math/rand"
)

type sessionData struct {
	Username string
}

type Session struct { //структура в которой храним данные
	data map[string]*sessionData //данные (data) с ключем string по обьекту sessionData
}

func NewSession() *Session { //создание сессии
	s := new(Session)                      //создфем обьект
	s.data = make(map[string]*sessionData) //данные делаем пустым мапом чтоб не был нилом
	return s
}

func (s *Session) init(username string) string {
	sessionId := GenerateId()

	data := &sessionData{Username: username}
	s.data[sessionId] = data

	return sessionId
}

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
