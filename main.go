package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

var args struct {
	NameFile string `arg:"-n,--names" help:"file containing list of names to generate usernames for"`
	Domain   string `arg:"-e,--email" help:"generate emails from name list containing specified domain"`
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
	userGen := UsernameGenerator{
		names: names,
	}

	fmt.Println(userGen)
}
