package core

import (
	"fmt"
	"strings"
)

type ControllerMapper struct {
	Controller *ControllerInfo
	Actions[]*ActionInfo
}

func (c *ControllerMapper)Print()  {


	fmt.Println("===========================================================")

	strBuilder:=new(strings.Builder)
	strBuilder.WriteString(fmt.Sprintf("[controoler]:%s\r\n",c.Controller.Name))

	for _,v:=range c.Actions{
		strBuilder.WriteString(fmt.Sprintf("--[action]:name%s->|method->%s\r\n",v.ActionName,v.attribute))
	}

	fmt.Println(strings.TrimRight(strBuilder.String(),"\r\n"))


}

func NewControllerMapper() *ControllerMapper {
	return &ControllerMapper{}
}

func (a *ControllerMapper)AddAction(action *ActionInfo)  {
	a.Actions=append(a.Actions,action)
}