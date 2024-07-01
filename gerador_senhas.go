package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"unicode"
)

func clear() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao limpar o terminal:", err)
	}
}

func gerador(quantidade int) {
	var senha string
	var random int

	letras := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	caracteres := []string{"!", "@", "#", "$", "%", "&", "_", "?"}

	for quantidade > 0 {
		random = rand.Intn(3)

		switch random {
		case 0:
			numero := strconv.Itoa(rand.Intn(10))
			senha += numero
		case 1:
			mai_ou_min := rand.Intn(2)
			
			switch mai_ou_min {
			case 0:
				letra_case0 := letras[rand.Intn(len(letras))]
				senha += letra_case0
			case 1:
				letra_case1 := letras[rand.Intn(len(letras))]
				letra_case1_min := string(unicode.ToLower(rune(letra_case1[0])))
				senha += letra_case1_min
			}
		case 2:
			caractere_escolhido := caracteres[rand.Intn(len(caracteres))]
			senha += caractere_escolhido
		}
		quantidade--
	}
	fmt.Println("Sua senha é:", senha)
}

func main() {
	var quantidade int

	fmt.Println("-----Gerador de senhas-----")

	fmt.Println("\nDigite quantas caracteres sua senha será composta:")
	fmt.Scan(&quantidade)

	if quantidade <= 0 {
		fmt.Println("\nErro, o numero tem que ser maior que 0")
		os.Exit(1)
	}

	gerador(quantidade)
}