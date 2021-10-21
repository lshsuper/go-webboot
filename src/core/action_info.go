package core

import (
	"fmt"
	"github.com/lshsuper/go-webboot/src/core/def"
	"strings"
)

type ActionInfo struct {
	ControllerName string
	ActionName string
	attribute string
	method    def.Method
}

func NewActionInfo() *ActionInfo {
	return &ActionInfo{}
}

//SetAttribute 设置注解
func (a *ActionInfo)SetAttribute(attr string)  {
	attr=strings.TrimLeft(attr,"//")
	a.attribute=attr

	//检测一下
	a.checkAttr()
}

//checkAttr 检测注解
func (a *ActionInfo)checkAttr()  {

	//注解格式标清
    if !strings.HasPrefix(a.attribute,"@"){
            panic(fmt.Sprintf("[位置]:ctrl->%s|action->%s-[信息]:注解使用不当",a.ControllerName,a.ActionName))
	}

	//解析一下方法请求方式
	for _,v:=range def.GetAllMethod(){
		if strings.Index(a.attribute,v.MethodAttr())>=0{
			a.method=v
			break
		}
	}
	//TODO 其它各种类型注解(暂时不支持)

}