Para evitar problemas de reconhecimento de pacotes pelo Go, separem cada arquivo em uma pasta própria, tendo um diretório composto pelas pastas onde cada arquivo está.

Caso ainda tenham problemas, executem o seguinte comando em um terminal: go env -w GO111MODULE=off



Utilização da Calculadora:
Serão necessário dois terminais para testar o programa da calculadora.

Em um terminal, inicializem o servidor utilizando: go run server_v2.go

Inicializem o cliente em outro terminal utilizando: go run <cliente de sua escolha>

O client_v4.go irá reconhecer comandos no formato <Operação> <Número> <Número>. Exemplo: Mult 2 4

As funções de Load e Store possuem a seguinte sintaxe: <Operação> <Posição> <Número>.
