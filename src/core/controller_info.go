package core

import (
	"fmt"
	"strings"
)

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
	//检测一下
	c.checkAttr()
}
//checkAttr 检测属性注解的合法性
func (c *ControllerInfo)checkAttr()  {

	if !strings.HasPrefix(c.attribute,"@"){
		panic(fmt.Sprintf("[位置]:ctrl->%s-[信息]:控制器注解不正确",c.Name))
	}


	for _,v:=range []string{"ApiController"}{
		if strings.Index(c.attribute,v)>0{
              break
		}
	}

	//TODO 其它各种类型注解(暂时不支持)


}
