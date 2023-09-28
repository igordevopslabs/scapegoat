# Scapegoat HTTP Server

Um simples servidor HTTP que responde a algumas rotas padrão e exibe informações sobre o estado do serviço.

## Rotas

- `/_ready`: Confirma que o servidor está pronto para receber requisições.
- `/_healthy`: Indica que o servidor está saudável.
- `/`: Exibe uma mensagem padrão indicando que o servidor está ativo e a variável de ambiente `ENV_NAME`.

## Uso

Para rodar o servidor:

```bash
go run <nome-do-arquivo>.go
```

Por padrão, o servidor será iniciado na porta 3000 e a variável de ambiente ENV_NAME será definida como development.

## Configuração

Você pode configurar o servidor usando as seguintes variáveis de ambiente:

`PORT`: A porta em que o servidor HTTP deve escutar. Padrão: 3000.
`ENV_NAME`: Nome do ambiente para ser exibido na rota raiz /. Padrão: development.


Para executar com configurações personalizadas:

```bash
PORT=8080 ENV_NAME=production go run <nome-do-arquivo>.go
```

## Detalhes de Implementação

* O servidor usa o pacote http padrão do Go para lidar com requisições.
* As respostas são fixas para cada rota.
* A função getEnv é usada para obter variáveis de ambiente com um valor padrão, caso a variável não esteja definida.
