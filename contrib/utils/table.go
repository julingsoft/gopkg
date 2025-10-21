package utils

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

func GetTableComment(ctx context.Context, tableName string) string {
	var (
		db      = g.DB()
		sql     = `SELECT TABLE_COMMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`
		comment = ""
	)

	result, err := db.GetAll(ctx, sql, db.GetConfig().Name, tableName)
	if err != nil {
		return ""
	}

	if len(result) > 0 {
		comment = result[0]["TABLE_COMMENT"].String()
	}

	if gstr.SubStrRune(comment, gstr.LenRune(comment)-1) == "表" {
		comment = gstr.SubStrRune(comment, 0, gstr.LenRune(comment)-1)
	}

	return comment
}

func GetTableColumns(ctx context.Context, table string) []map[string]string {
	tableFields, err := g.DB().TableFields(ctx, table)
	if err != nil {
		panic(err)
	}

	var allColumns = make([]map[string]string, len(tableFields))
	for _, fieldItem := range tableFields {
		fieldItem.Comment = gstr.ReplaceByMap(fieldItem.Comment, map[string]string{
			"：":  ":",
			"，":  ":",
			",":  ":",
			"(":  ":",
			"（":  ":",
			" ":  ":",
			"\t": ":",
		})

		if gstr.Contains(fieldItem.Comment, ":") {
			fieldItem.Comment = gstr.Split(fieldItem.Comment, ":")[0]
		}

		if g.IsEmpty(fieldItem.Comment) && fieldItem.Name == "id" {
			fieldItem.Comment = "ID"
		}

		allColumns[fieldItem.Index] = map[string]string{
			"Name":               fieldItem.Name,
			"Comment":            fieldItem.Comment,
			"CaseCamelName":      gstr.CaseCamel(fieldItem.Name),
			"CaseCamelLowerName": gstr.CaseCamelLower(fieldItem.Name),
			"Type":               GetFieldType(fieldItem.Type),
			"Key":                fieldItem.Key,
		}
	}

	return allColumns
}

func GetPriKey(ctx context.Context, table string) string {
	keyColumns := GetKeyColumns(ctx, table)

	var priKey = "id"
	for _, fieldInfo := range keyColumns {
		if gstr.InArray([]string{"PRI"}, fieldInfo["Key"]) {
			priKey = fieldInfo["Name"]
		}
	}

	return priKey
}

func GetKeyColumns(ctx context.Context, table string) []map[string]string {
	var allColumns = GetTableColumns(ctx, table)

	var keyColumns = make([]map[string]string, 0)
	for _, fieldInfo := range allColumns {
		if !g.IsEmpty(fieldInfo["Key"]) {
			keyColumns = append(keyColumns, fieldInfo)
		}
	}

	return keyColumns
}

func GetFieldType(fieldType string) string {
	m, err := gregex.MatchString(`(\w+)\(`, fieldType)
	if err != nil {
		panic(err)
	}

	if len(m) > 1 {
		fieldType = m[1]
	}

	var unsigned = gstr.ContainsI(fieldType, "unsigned")
	fieldType = gstr.ReplaceByMap(fieldType, map[string]string{
		" unsigned": "",
	})

	if gstr.InArray([]string{"bigint"}, fieldType) {
		if unsigned {
			return "uint64"
		}
		return "int64"
	}

	if gstr.InArray([]string{"bit", "int", "mediumint", "smallint", "tinyint", "enum"}, fieldType) {
		if unsigned {
			return "uint"
		}
		return "int"
	}

	if gstr.InArray([]string{"decimal", "float", "double"}, fieldType) {
		return "float64"
	}

	if gstr.InArray([]string{"blob", "binary"}, fieldType) {
		return "[]byte"
	}

	if gstr.InArray([]string{"date", "datetime", "timestamp", "time"}, fieldType) {
		return "*gtime.Time"
	}

	return "string"
}
