package def


//请求方法
type Method string

const(
	GET  Method ="GET"
	POST        ="POST"
)
var(
	allMethod=[]Method{GET,POST}
)

//String 转化字符串
func (e Method)String()string  {
	return string(e)
}

//MethodAttr 对应请求方法属性对应的具体值
func (e Method)MethodAttr()string  {

	switch e {
	case GET:
		return "HttpGet"
	case POST:
		return "HttpPost"
	default:
		return ""
	}
}

//GetAllMethod 获取所有方式的Method
func GetAllMethod()[]Method  {
	return allMethod
}
