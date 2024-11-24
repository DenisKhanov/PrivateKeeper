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

func (p *TextDataProvider) Save() {
	red := color.New(color.FgRed).SprintFunc()

	if !p.state.IsAuthorized() {
		fmt.Println(red("You are not authorized, please use 'login' or 'register'"))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	cyanBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	req := models.TextData{}
	fmt.Println(cyanBold("Input text data 'text metadata':"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("Input text data as %s: ", yellow("'text'"))
	scanner.Scan()
	data := scanner.Text()
	req.Content = data

	fmt.Printf("Input info as %s: ", yellow("'text'"))
	scanner.Scan()
	data = scanner.Text()
	req.Info = data

	err := p.textDataService.SaveTextData(p.state.GetToken(), req)
	if err != nil {
		logrus.WithError(err).Error("Text save failed")
		fmt.Println(red("Text save failed"), "please try again")
		lib.UnpackGRPCError(err)
		return
	}

	fmt.Println(color.New(color.FgGreen).SprintFunc()("Text data successfully saved"))
}
