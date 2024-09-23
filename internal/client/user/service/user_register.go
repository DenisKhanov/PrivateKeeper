package service //nolint:dupl

import (
	"bufio"
	"fmt"
	"github.com/DenisKhanov/PrivateKeeper/internal/domain"
	"github.com/DenisKhanov/PrivateKeeper/pkg/validate"
	"os"
	"strings"

	"github.com/fatih/color"
)

func (u *UserProvider) RegisterUser() {
	scanner := bufio.NewScanner(os.Stdin)
	red := color.New(color.FgRed).SprintFunc()

	var (
		name,
		email,
		login,
		password string
	)

	yellowBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(yellowBold("Input 'login password' to register:"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("Input your name as %s: ", yellow("'text'"))
	scanner.Scan()
	data := scanner.Text()
	name = strings.TrimSpace(data)
	if len(name) == 0 {
		fmt.Println(red("Name must not be empty please try again"))
		return
	}

	fmt.Printf("Input email as %s: ", yellow("'text'"))
	scanner.Scan()
	data = scanner.Text()
	email = strings.TrimSpace(data)
	if validate.CheckEmail(email) != nil {
		fmt.Println(red("Invalid email format please try again"))
		return
	}

	fmt.Printf("Input login as %s: ", yellow("'valid email'"))
	scanner.Scan()
	data = scanner.Text()
	login = strings.TrimSpace(data)
	err := validate.CheckLogin(login, domain.MinLoginLength, domain.MaxLoginLength)
	if err != nil {
		fmt.Println(red(err.Error()), "please try again")
		return
	}

	fmt.Printf("Input password as %s: ", yellow("'text'"))
	scanner.Scan()
	data = scanner.Text()
	password = strings.TrimSpace(data)
	err = validate.CheckPassword(password, domain.MinPasswordLength, domain.MaxPasswordLength)
	if err != nil {
		fmt.Println(red(err.Error()), "please try again")
		return
	}

	token, err := u.userService.RegisterUser(name, email, login, password)
	if err != nil {
		fmt.Println(red(err.Error()), "please try again")
	} else {
		u.state.SetToken(token)
		u.state.SetIsAuthorized(true)
		u.state.SetLogin(login)
	}
}
