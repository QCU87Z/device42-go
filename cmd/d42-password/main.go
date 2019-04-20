package main

import (
	"bufio"
	"fmt"
	"github.com/qcu87z/device42-go/pkg/device42"
	"log"
	"os"
	"strconv"
	"strings"
)

func mainMenu() {
	//todo Rewrite main menu using wmenu by dixonwille
	fmt.Println("Menu")
	fmt.Println("\t1: Password by Device")
	//todo Add func in library to support printing all devices
	fmt.Println("\t2: Print all Devices")
	fmt.Println("\t0: Exit")
	fmt.Print("Select a menu option: ")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(char)
	client := device42.NewBasicAuthClient("admin", "adm!nd42")
	switch char {
	case '1':
		fmt.Println("Running Password by Device")
		reader1 := bufio.NewReader(os.Stdin)
		text, _ := reader1.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if len(text) != 0 {
			p2, err := client.GetPasswordByDevice(text)
			if err != nil {
				fmt.Println(err.Error())
			}
			if p2.TotalCount > 0 {
				fmt.Println(p2.TotalCount)
				for index, element := range p2.Passwords {
					fmt.Printf("%d: ", index)
					fmt.Println(element.Username)
				}
			} else {
				fmt.Println("p2 is empty")
				break
			}
			fmt.Println("\tChoose a username: ")
			reader2 := bufio.NewReader(os.Stdin)
			username, _ := reader2.ReadString('\n')
			username = strings.TrimSuffix(username, "\n")
			usernameID, _ := strconv.Atoi(username)

			for _, a := range p2.Passwords {
				if a.Username == p2.Passwords[usernameID].Username {
					fmt.Printf("Password for username %s is %s", username, a.Password)
				}
			}
		}
	case '2':
		//errors.New("Not implemented yet")
		log.Fatal("Not implemented yet")
	case '0':
		fmt.Println("Exiting")
		os.Exit(0)
	default:
		fmt.Println("Invalid Choice")
		os.Exit(1)
	}
}

func main() {
	//client := device42_pass.NewBasicAuthClient("thor", "applepie")
	client := device42.NewBasicAuthClient("admin", "adm!nd42")

	p1, err := client.GetPasswordById(2)
	if err != nil {
		fmt.Println(err.Error())
	}
	if p1 != nil {

		output := fmt.Sprintf("%d, %s: %s", p1.Passwords[0].ID, p1.Passwords[0].Username, p1.Passwords[0].Password)
		fmt.Println(output)
	} else {
		fmt.Println("p1 is empty")
	}

	mainMenu()
}
