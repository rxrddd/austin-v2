package httpHelper

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintRoute(srv *http.Server) {
	// 声明一个二维数组来存放table的内容
	var tableData [][]string

	//初始化tableWriter
	table := tablewriter.NewWriter(os.Stdout)

	//上面的data为表格内容，还需要定义表格头部
	table.SetHeader([]string{"path", "method"})

	_ = srv.WalkRoute(func(info http.RouteInfo) error {
		tableData = append(tableData, []string{info.Path, info.Method})
		return nil
	})
	//将数据添加到table
	table.AppendBulk(tableData)

	//输出
	table.Render()
}
