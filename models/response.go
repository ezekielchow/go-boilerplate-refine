package models

type Error struct {
	Message string
}

type ListSuccess struct {
	Data    interface{}
	Headers interface{}
}
