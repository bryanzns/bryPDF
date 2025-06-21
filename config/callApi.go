package config

import (
	"context"
	"fmt"
	"log"
	"sync"

	"google.golang.org/genai"
)
func CallApi(messenger  chan Pdf, message  []string, APIkey string, wg *sync.WaitGroup){
	counter  := 0
    ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  APIkey,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
    log.Printf("erro: %v", err)
    }
  defer client.ClientConfig().HTTPClient.CloseIdleConnections()
	for _, v := range message {
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
		counter += 1
		messenger <- Pdf{
			Indice: counter,
			Text: result.Text(),
		}
	}
	close(messenger)
	wg.Done()
}
