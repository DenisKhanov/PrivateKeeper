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

func (p *DataProvider) LoadAllDataInfo() {
	red := color.New(color.FgRed).SprintFunc()

	if !p.state.IsAuthorized() {
		fmt.Println(red("You are not authorized, please use 'login' or 'register'"))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	allUserDataList, err := p.dataService.LoadAllDataInfo(p.state.GetToken())
	logrus.Info(allUserDataList)
	if err != nil {
		logrus.WithError(err).Error("All data info load failed")
		fmt.Println(red("All data info load failed"), "please try again")
		lib.UnpackGRPCError(err)
		return
	}

	fmt.Println("-------------------------------------")

	green := color.New(color.FgGreen).SprintFunc()

	var sb strings.Builder
	for _, userData := range allUserDataList {
		if userData != nil {
			sb.WriteString("Data ID: " + strconv.FormatUint(userData.DataId, 10) + "\n")
			sb.WriteString("Data type: " + userData.DataType + "\n")
			sb.WriteString("Description : " + userData.Description + "\n")
			sb.WriteString("-------------------------------------" + "\n")
		}
	}

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
