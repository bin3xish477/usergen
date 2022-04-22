package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
)

var args struct {
	NameFile string `arg:"-n,--names" help:"file containing list of names to generate usernames for"`
	Domain   string `arg:"-d,--domain" help:"generate emails from name list containing specified domain"`
	SaveFile string `arg:"-s,--save" help:"save userlist to specified file"`
}

func writeUserList(userList []string) {
	f, err := os.Create(args.SaveFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, user := range userList {
		fmt.Fprintln(f, user)
	}
	fmt.Printf("usernames written to: %s\n", args.SaveFile)
}

func main() {
	arg.MustParse(&args)

	if args.NameFile == "" {
		fmt.Println("must provide a file containing a list of names with `-n`")
		return
	}

	var names []string
	f, err := os.Open(args.NameFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		names = append(names, line)
	}

	var userList []string
	for _, name := range names {
		lower := strings.ToLower(name)

		userList = append(userList, strings.Replace(lower, " ", ".", -1))
		userList = append(userList, strings.Replace(lower, " ", "-", -1))
		userList = append(userList, strings.Replace(lower, " ", "", -1))
		split := strings.Split(lower, " ")
		var initial, firstName, lastName, lastInitial string
		firstName = split[0]
		initial = string(firstName[0])
		lastName = split[len(split)-1]
		lastInitial = string(lastName[0])
		//firstName, initial, lastName, lastInitial = split[0], string(firstName[0]), split[len(split)-1], string(lastName[0])
		userList = append(userList, strings.Join([]string{lastName, firstName}, ""))
		userList = append(userList, strings.Join([]string{initial, lastName}, ""))
		userList = append(userList, strings.Join([]string{lastName, initial}, ""))
		userList = append(userList, strings.Join([]string{firstName, lastInitial}, ""))
		userList = append(userList, strings.Join([]string{initial, ".", lastName}, ""))
		userList = append(userList, strings.Join([]string{lastInitial, ".", firstName}, ""))
		userList = append(userList, strings.Join([]string{firstName, ".", lastInitial}, ""))
		userList = append(userList, firstName)
		userList = append(userList, lastName)
	}

	if args.Domain != "" {
		for _, user := range userList {
			userList = append(userList, fmt.Sprintf("%s@%s", user, args.Domain))
		}
	}
	if args.SaveFile != "" {
		writeUserList(userList)
	} else {
		for _, user := range userList {
			fmt.Println(user)
		}
	}

}
