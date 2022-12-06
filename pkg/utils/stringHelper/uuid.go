package stringHelper

import "github.com/bwmarrin/snowflake"

func GenerateUUID() string {
	// 应用id 生成雪花随机数
	node, _ := snowflake.NewNode(1)
	// 生成雪花id 读取其中的9位
	snowflakeId := node.Generate().String()
	return snowflakeId
}
