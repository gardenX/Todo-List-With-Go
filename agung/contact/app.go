package contact

import (
	"bufio"
	"fmt"
	"net/textproto"
	"os"
	"strings"
)

func Add(contacts *map[string]string, name string, phone string) {
	contactsRef := *contacts

	if phone == "" {
		fmt.Println("Unable to store empty phone")
		return
	}

	contactsRef[name] = phone
	fmt.Println("New contact saved")
}

func Delete(contacts *map[string]string, name string) {
	contactsRef := *contacts
	target := contactsRef[name]

	defer func() {
		if target != "" {
			fmt.Println("Contact removed")
		}
	}()

	delete(contactsRef, name)
}

func Search(contacts map[string]string, name string) {
	if contacts[name] == "" {
		fmt.Println("Contact not found")
	} else {
		fmt.Println(showContact(name, contacts[name]))
	}
}

func Main() {
	var command string
	contacts := make(map[string]string)

	fmt.Println("Simple Contact")
	fmt.Println("-------------------")

	showHelp()

	for {
		fmt.Println()
		command = strings.ToUpper(readString("Command : "))

		switch command {
		case "L":
			list(contacts)
		case "A":
			Add(&contacts, readString("Contact name : "), readString("Phone number : "))
		case "R":
			Delete(&contacts, readString("Contact name : "))
		case "S":
			Search(contacts, readString("Contact name : "))
		case "H":
			showHelp()
		case "E":
			fmt.Println("Good bye")
			os.Exit(0)
		default:
			fmt.Println("Unknown command given")
		}
	}
}

func readString(prompt string) string {
	var target string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)
	target, _ = reader.ReadString('\n')

	return textproto.TrimString(target)
}

func list(contacts map[string]string) {
	if len(contacts) == 0 {
		fmt.Println("Contact empty")
		return
	}

	for key, phone := range contacts {
		fmt.Println(showContact(key, phone))
	}
}

func showContact(name string, phone string) string {
	return fmt.Sprintf("%s (%s)", name, phone)
}

func showHelp() {
	commands := []string{
		"Supported commands :",
		"(H) Help",
		"(L) List",
		"(A) Add",
		"(R) Remove",
		"(S) Search",
		"(E) Exit",
	}

	for _, command := range commands {
		fmt.Println(command)
	}
}
