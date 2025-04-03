package utils

import (
	"LeetCode/utils/constant"
	"fmt"
	"strings"
)

func OutputTaskCard(DescriptionTaskRU, DescriptionTaskEN string, inputData, result interface{}) {
	fmt.Print(titleSection())
	fmt.Print(descriptionSection(DescriptionTaskRU, DescriptionTaskEN))
	fmt.Print(solutionSection())
	fmt.Print(getInputData(inputData))
	fmt.Print()
}

func Divider() string {
	return strings.Repeat("=", 100)
}

func Tab(n int) string {
	return strings.Repeat("\t", n)
}

func titleSection() string {
	return fmt.Sprintf("%s\n%s%s\n%s\n",
		Divider(),
		Tab(3), constant.DescriptionHeader,
		Divider(),
	)
}
func descriptionSection(en, ru string) string {
	return fmt.Sprintf("%s%s\n%s\n%s%s\n%s\n",
		Tab(6), constant.PointerRuLang,
		ru,
		Tab(6), constant.PointerEnLang,
		en,
	)
}
func solutionSection() string {
	return fmt.Sprintf("%s\n%s%s\n%s\n",
		Divider(),
		Tab(4), constant.SolutionHeader,
		Divider(),
	)
}
func getInputData(in interface{}) string {
	return fmt.Sprintf("%s\t:\t%v\n", constant.InputDataDescription, in)
}

func getResultData(r interface{}) string {
	return fmt.Sprintf("")
}
