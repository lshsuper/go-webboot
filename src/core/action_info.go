package core

import "strings"

type ActionInfo struct {
	ControllerName string
	ActionName string
	attribute string
}

func NewActionInfo() *ActionInfo {
	return &ActionInfo{}
}

func (a *ActionInfo)SetAttribute(attr string)  {
	attr=strings.TrimLeft(attr,"//")
	a.attribute=attr
}