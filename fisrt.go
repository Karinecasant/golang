// 1) Estrutura básica de um arquivo executável: package main, lista de imports e função main.
// 2) Para executar arquivo no terminal: go run <nome arquivo>

package main // Referência do arquivo dentro da main. Necessário para tornar o arquivo executável.

import ( // Lista de imports. Referência dos pacotes.
	"fmt"
)

func main() {
	fmt.Println("Hello World") // Função main. Necessário para tornar o arquivo executável.
}
