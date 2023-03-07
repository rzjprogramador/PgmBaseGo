package linguagem

type Linguagem struct{
	Nome string
	EntradaSaidaDadosByLinguagem
	TiposDeDados
	Regras
	Arquivos
	ComandosDeSaida
	Variaveis
	Descobertas
}

type EntradaSaidaDadosByLinguagem struct{
	TemEscopoDeEntradaESaidaPrincipal bool
	QualEscopoDeEntradaESaidaPrincipal string
}

type TiposDeDados struct {
	Conceito string
	TipoPrimitivo
	TipoParaModeladores
	TipoParaVariaveis
	TipoParaAtributo
	TipoParaObjeto
	TipoParaFuncao
	TipoParaExcessoesErro
	TipoSemRetornoVazio
	TipoAceitaQualquerDadoANY
	NovoTipoPersonalizado
}

type TipoPrimitivo struct {
	Boleano ModelTipo
	Texto ModelTipoTexto
	CaractereUnico ModelTipo
	Numeros	
}

type ModelTipoTexto struct {
	Keyword string
	Representacao_PodeInstanciar string
	ValorDefault_SeNadaForPassado string
	Exemplo string
	RepresentacoesTexto RepresentacaoTexto
	
}

type RepresentacaoTexto struct {
	PularLinhas string
	Interpolar_Variavel_Em_Texto string
	Interpolar_ObjetoFuncao_Em_Texto string
	Caracteres_Especiais []string
	EscaparCaracteresEspeciaisEDeConflito string
}

type Numeros struct {
	QualquerNumero ModelTipo
	NumeroInteiro ModelTipo
	NumeroDecimal ModelTipo
}

type TipoParaModeladores struct {
	Modelador ModelTipo
}

type TipoParaVariaveis struct {
	Variavel ModelTipo
}

type TipoParaAtributo struct {
	Atributo ModelTipo
}

type TipoParaObjeto struct {
	Objeto ModelTipo 
}

type TipoParaFuncao struct {
	Funcao ModelTipo
}

type TipoParaExcessoesErro struct {
	Indefinido ModelTipo
	Nulo ModelTipo
	Erro ModelTipo
}

type TipoSemRetornoVazio struct {
	Vazio_Void ModelTipo
}

type TipoAceitaQualquerDadoANY struct {
	AnyQualquerInterface ModelTipo
}

type NovoTipoPersonalizado struct {}

type ModelTipo struct {
	Nome string
	TemNaLinguagem bool
	Keyword string
	Conceito string
	Representacao_PodeInstanciar string
	ValorDefault_SeNadaForPassado string
	Exemplo string
}

type Regras struct {
	ObrigatorioPontoeVirgulaACadaSentenca bool
}

type Arquivos struct {
	RodarArquivo string
}
type ComandosDeSaida struct {
	ViaConsole string
}

type Variaveis struct {
	PropsConceitoVariavel
	DeclaracaoVariavel string
	AspasAceitasEmString string
	AspasAceitasEmCaractereUnico string
}
type PropsConceitoVariavel struct {
	Significado string
	Sinonimos []string
}

type Descobertas struct {
	DescobrirValorViaConsole string
}
