package main

import (
	"time"

	"os"

	"github.com/helloticket/ffparser"
)

// HeaderBancoBrasil Registro Header de Arquivo-Remessa
type HeaderBancoBrasil struct {
	IdentificacaoRegistroHeader         int64     `record:"start=1,end=1"`
	TipoOperacao                        int64     `record:"start=2,end=2"`
	IdentificacaoPorExtensoTipoOperacao string    `record:"start=3,end=9"`
	TipoServico                         int64     `record:"start=10,end=11"`
	IdentificacaoPorExtensoTipoServico  string    `record:"start=12,end=19"`
	ComplementoRegistro1                string    `record:"start=20,end=26"`
	Agencia                             int64     `record:"start=27,end=30"`
	DigitoAgencia                       string    `record:"start=31,end=31"`
	ContaCorrente                       int64     `record:"start=32,end=39"`
	DigitoContaCorrente                 string    `record:"start=40,end=40"`
	ComplementoRegistro2                string    `record:"start=41,end=46,padchar=0"`
	NomeCedente                         string    `record:"start=47,end=76"`
	UsoExclusivoBancoBrasil             string    `record:"start=77,end=94"`
	DataGeracao                         time.Time `record:"start=95,end=100,decorator=BrazilSmallDateDecorator"`
	SequencialRemessa                   int64     `record:"start=101,end=107"`
	ComplementoRegistro3                string    `record:"start=108,end=129"`
	NumeroConvenioLider                 int64     `record:"start=130,end=136"`
	ComplementoRegistro4                string    `record:"start=137,end=394"`
	SequencialRegistro                  int64     `record:"start=395,end=400"`
}

// TraillerBancoBrasil Registro Trailler de Arquivo-Remessa
type TraillerBancoBrasil struct {
	IdentificacaoRegistro int64  `record:"start=1,end=1"`
	ComplementoRegistro1  string `record:"start=2,end=394"`
	SequencialRegistro    int64  `record:"start=395,end=400"`
}

// SegmentoDetalheSete Registro Detalhe de Arquivo-Remessa
type SegmentoDetalheSete struct {
	IdentificacaoRegistro                   int64     `record:"start=1,end=1"`
	TipoInscricaoCedente                    int       `record:"start=2,end=3"`
	NumeroCpfCnpjCedente                    int64     `record:"start=4,end=17"`
	Agencia                                 int       `record:"start=18,end=21"`
	DigitoAgencia                           string    `record:"start=22,end=22"`
	ContaCorrenteCedente                    int       `record:"start=23,end=30"`
	DigitoContaCorrenteCedente              string    `record:"start=31,end=31"`
	NumeroConvenio                          int       `record:"start=32,end=38"`
	CodigoControleEmpresa                   string    `record:"start=39,end=63"`
	NossoNumero                             int64     `record:"start=64,end=80"`
	NumeroPrestacao                         int       `record:"start=81,end=82"`
	GrupoValor                              int       `record:"start=83,end=84"`
	ComplementoRegistro                     string    `record:"start=85,end=87"`
	IndicativoMensagemSacadorOuAvalista     string    `record:"start=88,end=88"`
	PrefixoTitulo                           string    `record:"start=89,end=91"`
	VariacaoCarteira                        int       `record:"start=92,end=94"`
	ContaCaucao                             int       `record:"start=95,end=95"`
	NumeroBordero                           int       `record:"start=96,end=101"`
	TipoCobranca                            string    `record:"start=102,end=106"`
	CarteiraCobraca                         int       `record:"start=107,end=108"`
	Comando                                 int       `record:"start=109,end=110"`
	SeuNumeroOuNumeroTituloAtribuidoCedente string    `record:"start=111,end=120"`
	DataVencimento                          time.Time `record:"start=121,end=126,decorator=BrazilSmallDateDecorator"`
	ValorTitulo                             float64   `record:"start=127,end=139,decorator=BrazilMoneyDecorator"`
	NumeroBanco                             int       `record:"start=140,end=142"`
	AgenciaCobradora                        int       `record:"start=143,end=146"`
	DigitoAgenciaCobradora                  string    `record:"start=147,end=147"`
	EspecieTitulo                           int       `record:"start=148,end=149"`
	AceiteTitulo                            string    `record:"start=150,end=150"`
	DataEmissao                             time.Time `record:"start=151,end=156,decorator=BrazilSmallDateDecorator"`
	InstrucaoCodificada1                    int       `record:"start=157,end=158"`
	InstrucaoCodificada2                    int       `record:"start=159,end=160"`
	ValorMora                               float64   `record:"start=161,end=173,decorator=BrazilMoneyDecorator"`
	DataLimite                              time.Time `record:"start=174,end=179,decorator=BrazilMoneyDecorator"`
	ValorDesconto                           float64   `record:"start=180,end=192,decorator=BrazilSmallDateDecorator"`
	ValorIOF                                float64   `record:"start=193,end=205,decorator=BrazilMoneyDecorator"`
	ValorAbatimento                         float64   `record:"start=206,end=218,decorator=BrazilMoneyDecorator"`
	TipoInscricaoSacado                     int       `record:"start=219,end=220"`
	NumeroCnpjOuCpfSacado                   int64     `record:"start=221,end=234"`
	NomeSacado                              string    `record:"start=235,end=271"`
	ComplementoRegistro1                    string    `record:"start=272,end=274"`
	EnderecoSacado                          string    `record:"start=275,end=314"`
	BairroSacado                            string    `record:"start=315,end=326"`
	CepSacado                               int       `record:"start=327,end=334"`
	CidadeSacado                            string    `record:"start=335,end=349"`
	UfSacado                                string    `record:"start=350,end=351"`
	ObservacoesMensagemSacadorAvalista      string    `record:"start=352,end=391"`
	NumeroDiasProtesto                      string    `record:"start=392,end=393"`
	ComplementoRegistro2                    string    `record:"start=394,end=394"`
	SequencialRegistro                      int       `record:"start=395,end=400"`
}

func main() {
	ffp := ffparser.NewSimpleParser()

	file, err := os.Create("BB_REMESSA.rem")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := HeaderBancoBrasil{
		IdentificacaoRegistroHeader:         0,
		TipoOperacao:                        1,
		IdentificacaoPorExtensoTipoOperacao: "REMESSA",
		TipoServico:                         1,
		IdentificacaoPorExtensoTipoServico:  "COBRANCA",
		Agencia:                             5140,
		DigitoAgencia:                       "1",
		ContaCorrente:                       51445,
		DigitoContaCorrente:                 "6",
		NomeCedente:                         "EMPRESA",
		UsoExclusivoBancoBrasil:             "001BANCODOBRASIL",
		DataGeracao:                         time.Now(),
		SequencialRemessa:                   1,
		NumeroConvenioLider:                 1115292,
		SequencialRegistro:                  1,
	}

	detalhe7 := SegmentoDetalheSete{
		IdentificacaoRegistro:      7,
		TipoInscricaoCedente:       2,
		NumeroCpfCnpjCedente:       60637469000168,
		Agencia:                    5140,
		DigitoAgencia:              "1",
		ContaCorrenteCedente:       51445,
		DigitoContaCorrenteCedente: "5",
		NumeroConvenio:             1115292,
		CodigoControleEmpresa:      "29872920000000001",
		NossoNumero:                11152920000000001,
		NumeroPrestacao:            0,
		GrupoValor:                 0,
		VariacaoCarteira:           19,
		ContaCaucao:                0,
		NumeroBordero:              0,
		CarteiraCobraca:            17,
		Comando:                    1,
		DataVencimento:             time.Date(2017, 4, 28, 0, 0, 0, 0, time.UTC),
		ValorTitulo:                10.20,
		NumeroBanco:                1,
		AgenciaCobradora:           0,
		EspecieTitulo:              12,
		AceiteTitulo:               "N",
		ValorDesconto:              0,
		ValorIOF:                   0,
		ValorAbatimento:            0,
		TipoInscricaoSacado:        1,
		NumeroCnpjOuCpfSacado:      41276007213,
		NomeSacado:                 "JOSE DA SILVA",
		EnderecoSacado:             "RUA SEM SAIDA",
		BairroSacado:               "BAIRRO",
		CepSacado:                  580405430,
		CidadeSacado:               "JOAO PESSOA",
		UfSacado:                   "PB",
		SequencialRegistro:         2,
	}

	trailler := TraillerBancoBrasil{
		IdentificacaoRegistro: 9,
		SequencialRegistro:    3,
	}

	result, _ := ffp.ParseToText(&header)
	file.WriteString(result)
	file.WriteString("\n")

	result, _ = ffp.ParseToText(&detalhe7)
	file.WriteString(result)
	file.WriteString("\n")

	result, _ = ffp.ParseToText(&trailler)
	file.WriteString(result)
	file.WriteString("\n")
}
