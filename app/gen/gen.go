package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var g *gen.Generator
var db *gorm.DB

func init() {
	db, _ = gorm.Open(mysql.Open("root:root@tcp(192.168.127.128:3306)/austin-v2?parseTime=true&collation=utf8mb4_unicode_ci&loc=Asia%2FShanghai&charset=utf8mb4"), &gorm.Config{})
}

//model 查询生成器  gorm gen
func main() {
	g = gen.NewGenerator(gen.Config{
		OutPath: "./common/dal/query",
		Mode:    gen.WithoutContext,
	})

	dataMap := map[string]func(dtype string) string{
		"smallint":  func(dType string) string { return "int32" },
		"tinyint":   func(dType string) string { return "int32" },
		"mediumint": func(dType string) string { return "int32" },
		"bigint":    func(dType string) string { return "int64" },
	}
	g.WithDataTypeMap(dataMap)

	g.UseDB(db)
	var tableList []string
	tableList, _ = db.Migrator().GetTables()

	tableList = relationship(tableList) //需要处理关系的表

	//其他默认的表
	for _, v := range tableList {
		g.ApplyBasic(g.GenerateModel(v))
	}
	//g.ApplyInterface(func(CommonDao) {}, g.GenerateModel("la_user"))
	g.Execute()
}

//在这定义各种表关系
func relationship(tableList []string) []string {
	adminRoleModel := g.GenerateModel("la_system_admin_role",
		hasOne("Role", "la_system_auth_role", "id", "role_id"),
	)

	messageTemplateModel := g.GenerateModel("message_template",
		hasOne("SendAccountItem", "send_account", "id", "send_channel"),
	)
	adminModel := g.GenerateModel("la_system_auth_admin",
		gen.FieldRelate(field.HasMany, "AuthRoles",
			adminRoleModel,
			hasManyCfg("admin_id", "id"),
		),
		gen.FieldType("role", "SplitSlice"),
	)

	//给model加一个字段
	menuModel := g.GenerateModel("la_system_auth_menu",
		gen.FieldNew("Children", "[]*LaSystemAuthMenu", "gorm:\"-\""),
	)
	roleModel := g.GenerateModel("la_system_auth_role",
		gen.FieldType("menu_ids", "SplitSlice"),
	)

	g.ApplyBasic([]interface{}{
		adminModel,
		adminRoleModel,
		menuModel,
		roleModel,
		messageTemplateModel,
	}...)

	return clearTable(tableList, []string{
		adminModel.TableName,
		adminRoleModel.TableName,
		roleModel.TableName,
		menuModel.TableName,
		messageTemplateModel.TableName,
	})
}

func clearTable(oldList []string, delTable []string) []string {
	if len(delTable) <= 0 {
		return oldList
	}
	var newTables []string
	for _, s := range oldList {
		if !inArr(s, delTable) {
			newTables = append(newTables, s)
		}
	}
	return newTables
}
func inArr(table string, list []string) bool {
	for _, s := range list {
		if s == table {
			return true
		}
	}
	return false
}

func hasOne(fieldName, tableName, other, self string) gen.ModelOpt {
	return gen.FieldRelate(field.HasOne, fieldName,
		g.GenerateModel(tableName),
		hasOneCfg(other, self),
	)
}
func hasMany(fieldName, tableName, other, self string) gen.ModelOpt {
	return gen.FieldRelate(field.HasMany, fieldName,
		g.GenerateModel(tableName),
		hasManyCfg(other, self),
	)
}

/**
other 对方表字段
self 我方表字段
*/
func hasOneCfg(other, self string) *field.RelateConfig {
	return &field.RelateConfig{
		RelatePointer: true,
		//foreignKey:关联表的结构体字段 references:当前表的结构体字段;
		GORMTag: fmt.Sprintf("foreignKey:%s;references:%s", other, self),
	}
}

/**
other 对方表字段
self 我方表字段
*/
func hasManyCfg(other, self string) *field.RelateConfig {
	return &field.RelateConfig{
		RelateSlice: true,
		//foreignKey:关联表的结构体字段 references:当前表的结构体字段;
		GORMTag: fmt.Sprintf("foreignKey:%s;references:%s", other, self),
	}
}
