package utils

import (
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

func YesNo(label string) bool {
	prompt := promptui.Select{
		Label: label+"[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		logrus.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}