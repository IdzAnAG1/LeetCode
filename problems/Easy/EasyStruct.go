package Easy

import (
	"LeetCode/utils"
)

type Easy struct {
	DescriptionRU string
	DescriptionEN string
	InputData     interface{}
	Func          func() (output []interface{})
}

func (e *Easy) Play() {
	utils.OutputTaskCard(e.DescriptionEN, e.DescriptionRU, e.InputData, e.Func())
}
