package service

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"github.com/DenisKhanov/PrivateKeeper/internal/client/lib"
)

func (p *TextDataProvider) Load() {
	red := color.New(color.FgRed).SprintFunc()

	if !p.state.IsAuthorized() {
		fmt.Println(red("You are not authorized, please use 'login' or 'register'"))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	yellowBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(yellowBold("Input text ID to load text data:"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("Input text ID as %s: ", yellow("'dd'"))
	scanner.Scan()
	data := scanner.Text()
	metadataID, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		fmt.Println("Error: Invalid number format:", err)
		return
	}

	textData, err := p.textDataService.LoadTextData(p.state.GetToken(), metadataID)
	if err != nil {
		logrus.WithError(err).Error("Text data load failed")
		fmt.Println(red("Text data load failed"), "please try again")
		lib.UnpackGRPCError(err)
		return
	}

	fmt.Println("-------------------------------------")

	green := color.New(color.FgGreen).SprintFunc()

	var sb strings.Builder
	sb.WriteString("Text data: " + textData.Content + "\n")
	sb.WriteString("Text metadata: " + textData.Info + "\n")
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
