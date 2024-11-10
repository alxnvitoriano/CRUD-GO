package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Estrutura 'struct' que define os campos das informações da pessoa
type Person struct {
    Nome      string
    Sobrenome string
    Idade     int
    Endereco  string
    Email     string
    Telefone  string
    Time      string
}

// Variável global que armazena todos os registros
var people = make([]Person, 0)

// Função main que mostra o menu e entrada de usuários
func main() {
    reader := bufio.NewReader(os.Stdin)
    for { //laço de repeticao para mostrar o menu
        fmt.Println("Opções:")
        fmt.Println("[1] - Registrar pessoas")
        fmt.Println("[2] - Procurar pessoas")
        fmt.Println("[3] - Listar pessoas")
        fmt.Println("[e] - Sair do programa")
        fmt.Println("[R] - Limpar banco de dados")
        fmt.Print("Escolha uma opção: ")

        option, _ := reader.ReadString('\n')
        option = strings.TrimSpace(option)

        switch strings.ToLower(option) { //switch é um 'se', ele deixa todos os comandos prontos, porém só executa aquele que for solicitado pelo usuario
        case "1":
            registerPerson(reader)
        case "2":
            searchPerson(reader)
        case "3":
            listPeople()
        case "e":
            exitProgram()
        case "r":
            resetDatabase()
        default:
            fmt.Println("Opção inválida. Por favor, escolha outra opção.")
        }
    }
}

// Função para coletar dados de uma nova pessoa e adicioná-la à lista 'people'
func registerPerson(reader *bufio.Reader) {
    var person Person

    fmt.Print("Digite o nome: ")
    person.Nome, _ = reader.ReadString('\n')
    person.Nome = strings.TrimSpace(person.Nome)
    if person.Nome == "" { //condicional por conta do nome ser obrigatorio, logo, se o usuario nao digitar nada, aparecerá a mensagem abaixo
        fmt.Println("Nome é obrigatório.")
        return
    }

    fmt.Print("Digite o sobrenome: ")
    person.Sobrenome, _ = reader.ReadString('\n')
    person.Sobrenome = strings.TrimSpace(person.Sobrenome)

    fmt.Print("Digite a idade: ")
    ageInput, _ := reader.ReadString('\n')
    ageInput = strings.TrimSpace(ageInput)
    age, err := strconv.Atoi(ageInput)
    if err != nil || age <= 0 { //condicional para que a idade nao seja um numero negativo
        fmt.Println("Idade inválida. Por favor, digite um número válido.")
        return
    }
    person.Idade = age

    fmt.Print("Digite o endereço: ")
    person.Endereco, _ = reader.ReadString('\n')
    person.Endereco = strings.TrimSpace(person.Endereco)
    if person.Endereco == "" { //condicional por conta do endereço ser obrigatorio, logo, se o usuario nao digitar nada, aparecerá a mensag
        fmt.Println("Endereço é obrigatório.")
        return
    }

    fmt.Print("Digite o email: ")
    person.Email, _ = reader.ReadString('\n')
    person.Email = strings.TrimSpace(person.Email)

    fmt.Print("Digite o telefone: ")
    person.Telefone, _ = reader.ReadString('\n')
    person.Telefone = strings.TrimSpace(person.Telefone)

    fmt.Print("Digite o time de futebol: ") 
    person.Time, _ = reader.ReadString('\n')
    person.Time = strings.TrimSpace(person.Time)
    if person.Time == "" { //condicional por conta do time de futebol ser obrigatorio, logo, se o usuario nao digitar nada, aparecerá a mensag
        fmt.Println("Time de futebol é obrigatório.")
        return
    }

    people = append(people, person)
    fmt.Println("Pessoa registrada com sucesso.")
}

// Função para buscar uma pessoa por campos diferentes (nome, endereço, sobrenome, email, telefone ou time)
func searchPerson(reader *bufio.Reader) {
    if len(people) == 0 { //condicional, caso ainda nao tenha sido cadastrado nenhum nome a pesquisa termina sem precisar gastar memoria pesquisando 'sem ter o que pesquisar'
        fmt.Println("Nenhuma pessoa cadastrada ainda.")
        return
    }

    fmt.Println("Campos para pesquisar:")
    fmt.Println("[1] - Nome")
    fmt.Println("[2] - Sobrenome")
    fmt.Println("[3] - Endereço")
    fmt.Println("[4] - Email")
    fmt.Println("[5] - Telefone")
    fmt.Println("[6] - Time de futebol")
    fmt.Print("Escolha uma opção: ")

    fieldInput, _ := reader.ReadString('\n')
    fieldInput = strings.TrimSpace(fieldInput)
    field, err := strconv.Atoi(fieldInput)
    if err != nil || field < 1 || field > 6 { //condicional para que caso o numero escolhido fosse maior que o numero de opcoes do menu, apareça a mensagem abaixo
        fmt.Println("Campo inválido escolhido.")
        return
    }

    fmt.Print("Digite o valor para pesquisar: ")
    value, _ := reader.ReadString('\n')
    value = strings.TrimSpace(value)

    var found bool
    for _, person := range people {
        switch field { //switch para escolha do campo de pesquisa
        case 1:
            if person.Nome == value {
                found = true
                printPerson(person)
            }
        case 2:
            if person.Sobrenome == value {
                found = true
                printPerson(person)
            }
        case 3:
            if person.Endereco == value {
                found = true
                printPerson(person)
            }
        case 4:
            if person.Email == value {
                found = true
                printPerson(person)
            }
        case 5:
            if person.Telefone == value {
                found = true
                printPerson(person)
            }
        case 6:
            if person.Time == value {
                found = true
                printPerson(person)
            }
        }
    }

    if !found { //condicional para que a pessoa seja encontrada, caso nao bata informações, a mesnagem abaixo sera exibida
        fmt.Println("Nenhuma pessoa encontrada com o valor informado.")
    }
}

// Função para listar as pessoas registradas
func listPeople() {
    if len(people) == 0 { //condicional para que caso nao haja pessoa cadastrada o programa finalize ali
        fmt.Println("Nenhuma pessoa cadastrada ainda.")
        return
    }

    for _, person := range people {
        printPerson(person)
        fmt.Println()
    }
}

func printPerson(person Person) {
    fmt.Printf("Nome: %s\n", person.Nome)
    fmt.Printf("Sobrenome: %s\n", person.Sobrenome)
    fmt.Printf("Idade: %d\n", person.Idade)
    fmt.Printf("Endereço: %s\n", person.Endereco)
    fmt.Printf("Email: %s\n", person.Email)
    fmt.Printf("Telefone: %s\n", person.Telefone)
    fmt.Printf("Time de futebol: %s\n", person.Time)
}

// Função para finalizar a execução do programa
func exitProgram() {
    fmt.Println("Fechando o programa.")
    os.Exit(0)
}

// Função para limpar todos os dados do banco de dados
func resetDatabase() {
    people = make([]Person, 0)
    fmt.Println("Banco de dados limpo com sucesso.")
}
