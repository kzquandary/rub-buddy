package service

import (
	rubbuddychat "rub_buddy/features/chat"
)

type ChatServiceInterface struct {
	d rubbuddychat.ChatDataInterface
}

func New(d rubbuddychat.ChatDataInterface) rubbuddychat.ChatServiceInterface {
	return &ChatServiceInterface{
		d: d,
	}
}

func (s *ChatServiceInterface) GetChat(id uint, role string) ([]rubbuddychat.ChatInfo, error) {
	return s.d.GetChat(id, role)
}
