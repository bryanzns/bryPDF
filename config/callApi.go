package config

import (
	"context"
	"fmt"
	"log"
	"sync"

	"google.golang.org/genai"
)
func CallApi(mensageiro chan Pdf, mensagem []string, chaveapi string, wg *sync.WaitGroup){
	contador := 0
    ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  chaveapi,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
    log.Printf("erro: %v", err)
    }
  defer client.ClientConfig().HTTPClient.CloseIdleConnections()
	for _, v := range mensagem {
    // Monta o prompt para Gemini
    prompt := fmt.Sprintf("Traduza o seguinte texto para o português brasileiro, sem comentários extras, apenas o texto traduzido, Responda exclusivamente com a tradução, sem introdução, explicação ou comentários., traduzido:\n\n%s", v) 
    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.0-flash",
        genai.Text(prompt),
        nil,
    )
    if err != nil {
 log.Printf("erro: %v", err)
    }
		contador += 1
		mensageiro <- Pdf{
			Indice: contador,
			Texto: result.Text(),
		}
	}
	close(mensageiro)
	wg.Done()
}
