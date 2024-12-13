package mindnum

import (
	"embed"
)

var (
	//go:embed descriptions/*
	Embedded        embed.FS
	DescriptionPath = "descriptions"
)

type MindnumRepository interface {
	FindByNumber(number int) (*Mindnum, error)
}
