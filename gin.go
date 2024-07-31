package main

import (
	"github.com/gin-gonic/gin"
)

// ShouldBindGinValidator 参数绑定、Validate参数校验
func ShouldBindGinValidator(ctx *gin.Context, data interface{}) (err error) {
	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	if err = ctx.ShouldBind(data); err != nil {
		return err
	}
	if err = Validate(data); err != nil {
		return
	}
	return err
}

//
//// ValidatorMiddlewareGin .
//func ValidatorMiddlewareGin() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 执行处理请求之前的逻辑
//		data, err := c.GetRawData()
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//
//		// 使用反射创建param的反射对象
//		refVal := reflect.ValueOf(c.Request.GetBody)
//
//		// 获取反射对象的类型
//		paramType := refVal.Type()
//		fmt.Println(2222, paramType)
//
//		body := UnmarshalJsonBody(data)
//		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
//		fmt.Println(11111, body)
//		if err := Validate(data); err != nil {
//			return
//		}
//		fmt.Println("Middleware executing...")
//		// 获取请求参数
//		queryParams := c.Request.URL.Query()
//		formParams := c.Request.PostForm
//
//		// 获取请求体
//		var requestBody []byte
//		if c.Request.Body != nil {
//			// 读取请求体内容
//			//bodyBytes, _ := c.GetRawData()
//			//requestBody = string(bodyBytes)
//			//
//			//// 将请求体内容重新写回请求体，使其可以再次读取
//			//c.Request.Body = bodyBytes
//
//			requestBody, _ = ioutil.ReadAll(c.Request.Body)
//
//			b := reflect.TypeOf(requestBody).Name()
//
//			a := string(requestBody)
//			fmt.Println(11111, a, b)
//			if err := Validate(a); err != nil {
//				return
//			}
//
//			// Restore the request body for later use
//			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(requestBody)))
//
//		}
//		fmt.Printf("Request: Query: %v, Form: %v, Body: %s\n", queryParams, formParams, requestBody)
//		// 记录请求信息
//		fmt.Printf("[%s] %s %s\n", c.Request.Method, c.Request.URL.Path, c.Request.Proto)
//
//		// 处理请求
//		c.Next()
//
//		// 执行处理完请求后的逻辑
//		status := c.Writer.Status()
//		fmt.Printf("Status: %d\n", status)
//	}
//}
//
//// UnmarshalJsonBody 将body Unmarshal 为 Json
//func UnmarshalJsonBody(body []byte) (jsonBody map[string]interface{}) {
//	jsonBody = make(map[string]interface{})
//	err := json.Unmarshal(body, &jsonBody)
//	if err != nil { // 当不为 JsonBody 时
//		jsonBody["_text"] = string(body)
//	}
//	return jsonBody
//}
