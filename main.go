package main

import (
	"fmt"
)

func main() {
	zap := PayloadWhatsApp{MensagemRecebida: MensagemRecebida{
		Telefone: "55359999-9999",
		Texto:    "Olá, bom dia! Gostaria de agendar uma avaliação para implante. Vocês têm horário para sexta?",
	},
		AudioLink: "https://storage.aws.com/api-media/audio_lead_776655.ogg",
	}

	insta := PayloadInstagram{MensagemRecebida: MensagemRecebida{
		Telefone: "55358888-8888",
		Texto:    "Gostei muito desse apartamento! Qual seria o valor mínimo de entrada?"},
		PostRespondido: "https://www.instagram.com/p/C2xY_abc123/ (Campanha_Lançamento_Centro)",
	}

	padronizar := []Padronizador{zap, insta}

	ProcessarFila(padronizar)
}

type MensagemRecebida struct {
	Telefone string
	Texto    string
}

type PayloadWhatsApp struct {
	MensagemRecebida
	AudioLink string
}

type PayloadInstagram struct {
	MensagemRecebida
	PostRespondido string
}

// interface
type Padronizador interface {
	GerarJSON() string
}

func (p PayloadWhatsApp) GerarJSON() string {
	text := fmt.Sprintf(`{"origem": "WhatsApp", "telefone": "%s", "texto": "%s", "audio": "%s"}`, p.Telefone, p.Texto, p.AudioLink)
	return text
}

func (i PayloadInstagram) GerarJSON() string {
	text := fmt.Sprintf(`{"origem": "Instagram" , "telefone": "%s", "texto:", "%s", "post:","%s"}`, i.Telefone, i.Texto, i.PostRespondido)
	return text

}

func ProcessarFila(fila []Padronizador) {
	totalWhats := 0
	totalInsta := 0
	for _, v := range fila {

		pacotePronto := v.GerarJSON()
		fmt.Println("[SISTEMA] Disparando para o fluxo:", pacotePronto)
		switch v.(type) {
		case PayloadWhatsApp:
			totalWhats += 1
		case PayloadInstagram:
			totalInsta += 1

		}

	}

	fmt.Println("RELATÓRIO FINAL: ", totalWhats, "mensagens de WhatsApp e ", totalInsta, " mensagens de instagram roteadas")

}
