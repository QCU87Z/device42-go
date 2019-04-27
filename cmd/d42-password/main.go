package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/QCU87Z/device42-go/pkg/device42"
	"github.com/dixonwille/wmenu"
	"os"
	"strings"
)

const apiURL = "https://10.11.12.239/api/1.0"

func main() {
	mm := mainMenu()
	for {
		err := mm.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print("\n\t----------\n")
	}
}

func mainMenu() *wmenu.Menu {
	menu := wmenu.NewMenu("Select an option: ")
	menu.Option("Get password by Device name", "name", true, nil)
	//menu.Option("Get password by Device ID", "id", false, nil)
	menu.Option("Testing", "test", false, nil)
	menu.Option("Exit", nil, false, func(opt wmenu.Opt) error {
		os.Exit(0)
		return nil
	})
	menu.Action(func(opts []wmenu.Opt) error {
		if len(opts) != 1 {
			return errors.New("wrong number of options were chosen")
		}

		//fmt.Println(opts[0].Value)
		fmt.Printf("\tChoose a %s: ", opts[0].Value)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		switch opts[0].Value {
		case "name":
			client := device42.NewBasicAuthClient(apiURL, "admin", "adm!nd42")
			p, err := client.GetNewPasswordsByName(input)
			if err != nil {
				fmt.Println(err)
			}
			for _, pass := range p {
				fmt.Println(pass.Username)
				fmt.Println(pass)

			}
		case "id":
			fmt.Println("Not implemented")

		case "test":
			fmt.Println("Not implemented")
		}
		return nil
	})
	return menu
}