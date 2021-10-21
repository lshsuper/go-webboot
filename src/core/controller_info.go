package core

import "strings"

type ControllerInfo struct {
	Name string
	attribute string
}

func NewControllerInfo() *ControllerInfo {
	return &ControllerInfo{}
}

func (c *ControllerInfo)SetAttribute(attr string)  {
	attr=strings.TrimLeft(attr,"//")
	c.attribute=attr
}
