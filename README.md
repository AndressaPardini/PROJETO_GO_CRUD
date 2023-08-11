# Projeto de Back-end com Foco em Testes Unitários usando Golang

Este é um projeto desenvolvido para o time de SMB no projeto NUIMP da empresa Stone. O objetivo principal é criar um sistema de CRUD (Create, Read, Update, Delete) com foco em testes unitários usando a linguagem de programação Golang.

## Funcionalidades

- Criação, leitura, atualização e exclusão de registros de usuários em um banco de dados.
- Implementação de testes unitários usando o padrão Triple A (Arrange, Act, Assert) para garantir a qualidade do código.

## Pacotes Utilizados

- **Banco-de-dados.go**: Um pacote que realiza a conexão com o banco de dados MySQL e executa consultas.
- **Banco.go**: Um pacote que define interfaces e implementações para interagir com o banco de dados, incluindo preparação de consultas e manipulação de resultados.
- **Servidor.go**: Um pacote que lida com as rotas HTTP e as operações CRUD, incluindo a criação de usuários.

## Instalação

1. Certifique-se de ter o Golang instalado. Caso contrário, [instale-o](https://golang.org/doc/install).
2. Clone este repositório para sua máquina local:

```bash
git clone git@github.com:AndressaPardini/PROJETO_GO_STONE.git
```

3. Execute os seguintes comandos para instalar as dependências e compilar o projeto:

```bash
cd PROJETO_GO_STONE
go mod tidy
go build
```

## Uso

1. Configure as informações de conexão do banco de dados no arquivo `Banco.go`.
2. Execute o binário compilado:

```bash
./projeto_go_stone
```

3. Acesse as rotas HTTP para interagir com o CRUD, por exemplo:

- `POST /usuarios`: Cria um novo usuário.
- `GET /usuarios`: Obtém a lista de usuários.
- `GET /usuarios/{id}`: Obtém os detalhes de um usuário específico.
- `PUT /usuarios/{id}`: Atualiza os dados de um usuário.
- `DELETE /usuarios/{id}`: Exclui um usuário.

## Testes Unitários

O projeto inclui testes unitários para garantir a robustez e a qualidade do código. Os testes estão localizados no arquivo `servidor_test.go`. Para executar os testes, execute o seguinte comando:

```bash
go test
```

## Contribuição

Contribuições são bem-vindas! Se você deseja contribuir com melhorias, correções de bugs ou novos recursos, siga estas etapas:

1. Fork este repositório.
2. Crie um branch com a sua feature (`git checkout -b minha-feature`).
3. Faça commit das suas alterações (`git commit -m 'Adicionando nova feature'`).
4. Faça push para o branch (`git push origin minha-feature`).
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).

Copyright (c) 2020 OtavioGallego

Permissão é concedida, gratuitamente, a qualquer pessoa que obtenha uma cópia
deste software e dos arquivos de documentação associados (o "Software"), para lidar
no Software sem restrições, incluindo, sem limitação, os direitos
para usar, copiar, modificar, mesclar, publicar, distribuir, sublicenciar e/ou vender
cópias do Software, e para permitir que pessoas a quem o Software é
fornecido façam isso, sujeitas às seguintes condições:

O aviso de direitos autorais acima e este aviso de permissão devem ser incluídos em todas
cópias ou partes substanciais do Software.

O SOFTWARE É FORNECIDO "COMO ESTÁ", SEM GARANTIA DE QUALQUER TIPO, EXPRESSA OU
IMPLÍCITA, INCLUINDO, MAS NÃO SE LIMITANDO ÀS GARANTIAS DE COMERCIALIZAÇÃO,
ADEQUAÇÃO A UM DETERMINADO FIM E NÃO VIOLAÇÃO. EM NENHUMA HIPÓTESE OS AUTORES OU DETENTORES
DE DIREITOS AUTORAIS SERÃO RESPONSÁVEIS POR QUALQUER REIVINDICAÇÃO, DANOS OU OUTRA
RESPONSABILIDADE, SEJA EM AÇÃO DE CONTRATO, DELITO OU DE OUTRA FORMA, DECORRENTE DE,
FORA OU EM CONEXÃO COM O SOFTWARE OU O USO OU OUTRAS NEGOCIAÇÕES NO
SOFTWARE.

## Contato

Para perguntas ou mais informações, entre em contato pelo email [andressa.pardini@stone.com.br](mailto:andressa.pardini@stone.com.br) ou visite meu perfil no GitHub: [Andressa Pardini](https://github.com/AndressaPardini).
