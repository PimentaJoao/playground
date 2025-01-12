# Signal Filters

Coleção de filtros estudados no meu trabalho de conclusão de curso, responsáveis por desenvolver os conceitos iniciais
nas diferentes técnicas de filtragem, culminando na introdução à família de filtros gaussianos por meio do estudo do 
filtro de Kalman.

## Uso

### Pré-requisitos

- Python versão 3.13 ou superior;
- Clonar este repositório e entrar na pasta ./signal-filters;
- Criar um ambiente virtual para a instalação das dependências:

```sh
    cd ./playground/signal-filters

    # Criando o ambiente virtual do projeto.
    python<version> -m venv <virtual-environment-name>

    # Por exemplo:

    python3 -m venv env
```
- Ativar o ambiente virtual:

```sh
    source ./<virtual-environment-name>/bin/activate

    # Por exemplo:

    source ./env/bin/activate
```

- Instalação das dependências utilizadas através do comando à seguir:

```sh
    pip install -r ./requirements.txt
```

- Ao final do uso, para desativar o ambiente virtual:

```sh
    deactivate
```

### Execução

```sh
    python3 <nome-do-arquivo>.py
```
