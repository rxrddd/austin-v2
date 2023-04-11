package idHelper

import (
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/cast"
)

var node *snowflake.Node

func init() {
	node, _ = snowflake.NewNode(1)
}
func NextID() int64 {
	// 应用id 生成雪花随机数
	// 生成雪花id 读取其中的9位
	snowflakeId := node.Generate().Int64()
	return snowflakeId
}
func NextIDStr() string {
	return cast.ToString(NextID())
}
