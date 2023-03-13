package tecnica

import "github.com/rzjprogramador/PgmBase/utils"

type Tecnica struct {
	Nome string
	Tecnica string
	Exemplo string
}

func newTecnica(t Tecnica) Tecnica{
	return t
}

func ExecuteTecnica() {
	utils.ShowObject(newTecnica(fooTecnica))
}

var fooTecnica = Tecnica{
	Nome: "nome da tecnica",
	Tecnica: "tecnica em questao",
	Exemplo: `exemplo da tecnica`,
}