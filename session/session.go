package session

import (
	"blogg/utils"
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

func (s *Session) Init(username string) string {
	sessionId := utils.GenerateId()

	data := &sessionData{Username: username}
	s.data[sessionId] = data

	return sessionId
}

func (s *Session) Get(sessionId string) string {
	data := s.data[sessionId]

	if data == nil {
		return ""
	}

	return data.Username
}
