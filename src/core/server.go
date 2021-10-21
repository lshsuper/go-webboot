package core

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)
//WebBootServer 服务
type WebBootServer struct {

}

func NewWebBootServer(addr string)*WebBootServer {
	return &WebBootServer{}
}

//AutoRegister 自动注册
func (wServer *WebBootServer)AutoRegister()  {

	wServer.parseController()

}

//Start 启动
func (wServer *WebBootServer)Start()  {
	
}
//Stop 停掉
func (wServer *WebBootServer)Stop()  {
	
}


//parseController 转化
func (wServer *WebBootServer)parseController() []*ControllerMapper {

	var controllerMappers []*ControllerMapper

	filepath.Walk("./", func(p string, info os.FileInfo, err error) error {
		if !info.IsDir()&&strings.HasSuffix( info.Name(),"controller.go"){
			fs,_:=ioutil.ReadFile(p)
			t:=token.NewFileSet()
			p,err:=parser.ParseFile(t,"" ,string(fs), parser.ParseComments)
			if err!=nil{
				panic(fmt.Sprintf("位置:%s-文件解析异常",info.Name()))
			}

			//遍历定义

			var (
				curController *ControllerInfo
				curMapper = NewControllerMapper()
			)
			for _,v:=range p.Decls{

				switch v.(type) {
				case *ast.GenDecl:

					//控制器数量限定
					if curController!=nil{
						panic(fmt.Sprintf("位置:%s-Info:%s",info.Name(),"一个控制器文件尽可以定义一个控制器"))
					}

					curDecl:=v.( *ast.GenDecl)
					typeSpec,ok:=curDecl.Specs[0].(*ast.TypeSpec)
					if !ok{
						continue
					}

					//控制器结构体定义标准限定
					if strings.Index(typeSpec.Name.Name,"Controller")<0{
						panic(fmt.Sprintf("位置:%s-Info:%s",info.Name(),"控制器结构体定义不标准,需以Controller后缀结束"))
					}

					//注解定义标准限定
					if curDecl.Doc==nil||curDecl.Doc.List==nil{
						panic(fmt.Sprintf("位置:%s-Info:%s",info.Name(),"控制器结构体定义不标准,请为控制器打上合适的标签"))
					}

					curController= NewControllerInfo()
					curController.Name=typeSpec.Name.Name
					curController.SetAttribute(curDecl.Doc.List[0].Text)
					curMapper.Controller=curController
					//txt:=strings.ToLower(typeSpec.Name.Name)
					//group=txt[0:strings.Index(txt,"controller")]
					//tag:=strings.TrimLeft(curDecl.Doc.List[0].Text,"//")
					//fmt.Println(fmt.Sprintf("【控制器】:%s-【注解】:%s",curDecl.Specs[0].(*ast.TypeSpec).Name.Name,tag))

				case *ast.FuncDecl:
					curDecl:=v.( *ast.FuncDecl)
					if curDecl.Recv==nil{
						continue
					}

					if curController==nil{
						panic(fmt.Sprintf("位置:%s-Info:%s",info.Name(),"不存在对应的控制器,请检查定义相关（eg:控制器结构体需要优先定义）"))
					}

					fromController:=curDecl.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
					if fromController!=curController.Name{
						panic(fmt.Sprintf("位置:%s-Func:%s-Info:%s",info.Name(),curDecl.Name,"存在控制器方法绑定不规范"))
					}

					if curDecl.Doc==nil||curDecl.Doc.List==nil{
						panic(fmt.Sprintf("位置:%s|%s:%s-Info:%s",info.Name(),curDecl.Name,"请为该方法定义合适的注解"))
					}

					curAction:= NewActionInfo()
					curAction.ControllerName=curController.Name
					curAction.ActionName=curDecl.Name.Name
					curAction.SetAttribute(curDecl.Doc.List[0].Text)
					curMapper.AddAction(curAction)
					//tag:=strings.TrimLeft(curDecl.Doc.List[0].Text,"//")
					//fmt.Println(fmt.Sprintf("【Action】:%s-【注解】:%s-【Router】:%s",curDecl.Name,tag,group+"/"+strings.ToLower(curDecl.Name.Name)))
				default:
					continue

				}

			}
			//收纳所有定义
			controllerMappers=append(controllerMappers,curMapper)
		}

		return err
	})

	//控制器映射合法性判断
	if controllerMappers==nil{
		panic("[信息]:未检测到任何控制器及绑定方法")
	}

    return controllerMappers

}




