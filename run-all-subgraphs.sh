#!/bin/bash

# Lista de pastas dos subgraphs
subgraph_folders=("carts" "items" "pets" "reviews" "users")

# Função para encerrar todos os processos
cleanup() {
    echo "Encerrando servidores..."
    for pid in "${server_pids[@]}"; do
        kill "$pid" 2>/dev/null
    done
    exit
}

# Captura o sinal Ctrl+C (SIGINT) e chama a função de cleanup
trap cleanup SIGINT

# Array para armazenar os PIDs dos processos
declare -a server_pids

# Loop sobre cada pasta
for folder in "${subgraph_folders[@]}"
do
    echo "Entrando na pasta $folder"
    cd "$folder"

    echo "Executando 'go run server.go' em $folder"
    # Executa o comando em segundo plano e armazena o PID
    go run server.go &

    # Salva o PID do processo
    server_pids+=($!)

    # Voltar para a pasta pai
    cd ..
done

# Bloqueia o terminal e espera por uma interrupção
echo "Todos os servidores estão rodando. Pressione Ctrl+C para encerrar."

# Espera todos os processos em segundo plano terminarem
wait