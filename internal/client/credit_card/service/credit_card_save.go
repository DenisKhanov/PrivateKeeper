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

func (p *CreditCardProvider) Save() {
	red := color.New(color.FgRed).SprintFunc()

	if !p.state.IsAuthorized() {
		fmt.Println(red("You are not authorized, please use 'login' or 'register'"))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	cyanBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	cardData := models.CardData{}
	fmt.Println(cyanBold("Input credit card data 'number, owner, expires, cvv, info':"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("Input number in format %s: ", yellow("'dddd dddd dddd dddd'"))
	scanner.Scan()
	data := scanner.Text()
	cardData.Number = data

	fmt.Printf("Input owner in format %s: ", yellow("'name surname'"))
	scanner.Scan()
	data = scanner.Text()
	cardData.HolderName = data

	fmt.Printf("Input expiry date in format %s: ", yellow("'dd-mm-yyyy'"))
	scanner.Scan()
	data = scanner.Text()
	cardData.ExpDate = data

	fmt.Printf("Input cvv in format %s: ", yellow("'ddd'"))
	scanner.Scan()
	data = scanner.Text()
	cardData.CVV = data

	fmt.Printf("Input card info as %s: ", yellow("'text'"))
	scanner.Scan()
	data = scanner.Text()
	cardData.Info = data

	err := p.creditCardService.SaveCreditCard(p.state.GetToken(), cardData)
	if err != nil {
		logrus.WithError(err).Error("Card save failed")
		fmt.Println(red("Card save failed"), "please try again")
		lib.UnpackGRPCError(err)
		return
	}

	fmt.Println(color.New(color.FgGreen).SprintFunc()("Card successfully saved"))
}
