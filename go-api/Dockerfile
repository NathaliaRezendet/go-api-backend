FROM golang:1.22.4

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o código-fonte do projeto para o contêiner
COPY . .

# Expor a porta que a aplicação irá rodar
EXPOSE 8000

# Baixe o wait-for-it.sh
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

ENV GIN_MODE=release

# Compilar o projeto Go com detalhes de saída adicionais
RUN go build -o main main.go

# Comando para rodar as migrações e iniciar a aplicação
CMD ["./wait-for-it.sh", "mysql_db:3306", "--", "./main"]
