package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"github.com/DenisKhanov/PrivateKeeper/internal/client/lib"
)

func (p *LoginPasswordProvider) Load() {
	red := color.New(color.FgRed).SprintFunc()

	if !p.state.IsAuthorized() {
		fmt.Println(red("You are not authorized, please use 'login' or 'register'"))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	cyanBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(cyanBold("Input login/password ID to load login_password 'login password metadata':"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("Input login/password ID in format %s: ", yellow("'dd'"))
	scanner.Scan()
	data := scanner.Text()
	metadataID, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		fmt.Println("Error: Invalid number format:", err)
		return
	}
	loginData, err := p.loginPasswordService.LoadLoginPassword(p.state.GetToken(), metadataID)
	if err != nil {
		lib.UnpackGRPCError(err)
		return
	}

	fmt.Println("-------------------------------------")

	green := color.New(color.FgGreen).SprintFunc()

	var sb strings.Builder
	sb.WriteString("Credential login: " + loginData.Login + "\n")
	sb.WriteString("Credential password: " + loginData.Password + "\n")
	sb.WriteString("Credential metadata: " + loginData.Info + "\n")
	sb.WriteString("-------------------------------------" + "\n")

	fmt.Print(green("Write info to file or print (leave empty or write to file): "))
	scanner.Scan()
	path := scanner.Text()

	if len(path) == 0 {
		fmt.Print(sb.String())
		return
	}

	if p.state.GetDirPath() != "" {
		path = p.state.GetDirPath() + "/" + path
	}

	err = lib.SaveToFile(path, sb.String())
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Error writing to file with path %s, please try again\n", red(path))
		return
	}

	fmt.Printf("Data successfully written to file %s\n", green(path))
}
