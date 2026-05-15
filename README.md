# Webhook Router em Go 🚀

Um middleware escalável desenvolvido em Go para receber, padronizar e rotear payloads de diferentes fontes de mensageria (como WhatsApp e Instagram). O projeto aplica conceitos de ETL (Extract, Transform, Load) para preparar dados desestruturados em um formato JSON padronizado antes de enviá-los para fluxos de automação avançados, como n8n ou CRMs locais.

## Conceitos Técnicos Aplicados

* **Arquitetura Baseada em Interfaces:** Uso de polimorfismo para processar diferentes tipos de mensagens em uma fila de execução única.
* **Composição de Structs:** Aplicação do modelo de Orientação a Objetos nativo do Go para reaproveitamento de código e organização de dados em vez de herança tradicional.
* **Roteamento e Validação:** Utilização de Type Switches (`switch type`) para aplicar regras de negócios específicas e gerar métricas de uso segmentadas por canal.


## ⚠️ Nota de Arquitetura (Projeto Educacional)
Este repositório foi construído com o objetivo estrito de treinar os conceitos de **Interfaces, Polimorfismo e Composição** em Go. O output de dados (o texto que simula o JSON) foi gerado via formatação de strings (`fmt.Sprintf`). Em um ambiente de produção real, a serialização de dados seria delegada ao pacote nativo `encoding/json` da linguagem. 

## Como rodar o projeto

1. Clone o repositório na sua máquina.
2. Abra o terminal na pasta do projeto.
3. Execute o comando
`go run main.go`