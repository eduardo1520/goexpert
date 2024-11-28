package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, url := range os.Args[1:] {
		var req *http.Response
		var err error
		var res []byte
		var data ViaCEP
		var file *os.File

		if req, err = http.Get("https://viacep.com.br/ws/" + url + "/json"); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer request para %s: %v\n", url, err)
		}

		defer req.Body.Close()

		if res, err = io.ReadAll(req.Body); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta da %s: %v\n", url, err)
		}

		if err = json.Unmarshal(res, &data); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta %s: %v\n", url, err)
		}

		filePath := "C:\\Users\\USER\\OneDrive\\Área de Trabalho\\cidade.txt"

		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao abrir/criar arquivo: %v\n", err)
			continue
		}

		defer file.Close()

		if _, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf)); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}

		fmt.Println(data)

	}
}
