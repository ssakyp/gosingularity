package cli

import (
	"bufio"
	"fmt"
	"github.com/ssakyp/gosingularity/lesson1/greet"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	name, err := Reader("Your name: ")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	tableStr, err := Reader("Your table: ")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	table, err := strconv.Atoi(tableStr)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	greeting := greet.Greet(name, table)
	fmt.Println(greeting)

}

func Reader(text string) (string, error) {
	temp := ""
	errMes := ""
	reader := bufio.NewReader(os.Stdin)
	for temp == "" {
		fmt.Printf("%s%s", errMes, text)
		var err error
		temp, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		temp = strings.TrimSpace(temp)
		if temp == "" {
			errMes = "Invalid input: empty input\n"
		}
	}
	return temp, nil
}
