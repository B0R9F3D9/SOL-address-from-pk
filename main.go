package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/gagliardetto/solana-go"
)

func readKeysFromFile(fileName string) []string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		panic(err)
	}
	lines := strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n")
	return lines
}

func addressFromKey(key string) string {
	privateKey, err := solana.PrivateKeyFromBase58(key)
	if err != nil {
		fmt.Printf("Ошибка при создании приватного ключа: %v\n", err)
		return "Ошибка: " + err.Error()
	}
	return privateKey.PublicKey().String()
}

func appendAddressToFile(address string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(address + "\n"); err != nil {
		fmt.Printf("Ошибка при записи в файл: %v\n", err)
		panic(err)
	}
}

func main() {
	// очистка addresses.txt перед запуском
	os.OpenFile("addresses.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	
	wallets := readKeysFromFile("private_keys.txt")
	for _, wallet := range wallets {
		address := addressFromKey(wallet)
		appendAddressToFile(address, "addresses.txt")
	}
}