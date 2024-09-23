package service

import (
	"bufio"
	"fmt"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/fatih/color"

	"github.com/DenisKhanov/PrivateKeeper/internal/client/lib"
)

func (p *LoginPasswordProvider) Save() {
	red := color.New(color.FgRed).SprintFunc()

	if !p.state.IsAuthorized() {
		fmt.Println(red("You are not authorized, please use 'login' or 'register'"))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	cyanBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	loginData := models.LoginData{}
	fmt.Println(cyanBold("Input login_password data 'login, password, info':"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("Input login as %s: ", yellow("'text'"))
	scanner.Scan()
	loginData.Login = scanner.Text()

	fmt.Printf("Input password as %s: ", yellow("'text'"))
	scanner.Scan()
	loginData.Password = scanner.Text()

	fmt.Printf("Input info as %s: ", yellow("'text'"))
	scanner.Scan()
	loginData.Info = scanner.Text()

	err := p.loginPasswordService.SaveLoginPassword(p.state.GetToken(), loginData)
	if err != nil {
		logrus.WithError(err).Error("Login/password save failed")
		fmt.Println(red("Login/password save failed"), "please try again")
		lib.UnpackGRPCError(err)
		return
	}

	fmt.Println(color.New(color.FgGreen).SprintFunc()("Login/password successfully saved"))
}
