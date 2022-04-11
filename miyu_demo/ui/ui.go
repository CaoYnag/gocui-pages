package ui

type CUI interface {
	Init() error
	Run() error
	Release()
}
