package linguagem

import (
	"github.com/rzjprogramador/PgmBase/utils"
)

func NewLinguagem(l Linguagem) *Linguagem{
	return &l
}

func ExecuteLinguagem(l Linguagem) {
	utils.ShowObject(l)
}

func MainLinguagem() {
	ExecuteLinguagem(LinguagemDart)
}

