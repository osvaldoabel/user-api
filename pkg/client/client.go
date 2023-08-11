package client

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/faelp22/browser"
)

var Browser browser.BrowserCli = browser.NewBrowser(browser.BrowserConfig{
	BaseURL:   "https://api.github.com",
	SSLVerify: false,
	Header: http.Header{
		"Content-Type": []string{"application/json; charset=utf-8"},
	},
	Timeout: 3, // espera 3 segundos
})

func Exemplo() {

	fmt.Println(Browser.GetHeader())

	resp, err := Browser.Get("/users/faelp22")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler Body da request")
		fmt.Println(err.Error())
	}

	if resp.StatusCode > 200 {
		fmt.Println("Erro na requisição")
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
		fmt.Println(string(body))
		os.Exit(1)
	}

	fmt.Println(string(body))
}
