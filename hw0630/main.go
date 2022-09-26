package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func main() {
	r := gin.Default()
	//git flow test
	// 仅读取文件
	r.GET("/read", func(c *gin.Context) {
		cont, err := ioutil.ReadFile("./permission_schema.yml")
		if err != nil {
			fmt.Println("读取文件失败：", err)
		}
		c.String(http.StatusOK, "%s", cont)
	})

	//解析yaml文件，可以对其中出现的错误进行识别
	r.GET("/parse", func(c *gin.Context) {
		yamlContent, err := ioutil.ReadFile("./permission_schema.yml")
		if err != nil {
			fmt.Println("读取文件失败", err)
		}
		// 存储解析数据
		var result yaml.Node
		err = yaml.Unmarshal([]byte(yamlContent), &result)
		if err != nil {
			c.String(http.StatusBadRequest, "%s", err)
		}
		text, err := yaml.Marshal(&result)
		if err != nil {
			c.String(http.StatusBadRequest, "%s", err)
		}
		c.String(http.StatusOK, "%s", text)
	})
	r.Run()
}
