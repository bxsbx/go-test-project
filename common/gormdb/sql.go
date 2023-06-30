package gormdb

import (
	"fmt"
	"strings"
)

// 批量插入(默认不跳过错误插入）
func BatchInsert[T any](tableName string, fields []string, values []T, f func(v T) []interface{}, ignore ...bool) (sql string, args []interface{}) {

	ignoreStr := ""
	if len(ignore) > 0 && ignore[0] {
		ignoreStr = "IGNORE"
	}

	repeat := "(" + strings.TrimSuffix(strings.Repeat("?,", len(fields)), ",") + "),"
	repeats := strings.TrimSuffix(strings.Repeat(repeat, len(values)), ",")
	sql = fmt.Sprintf("INSERT %s INTO %s (%s) VALUES %s", ignoreStr, tableName, strings.Join(fields, ","), repeats)
	args = make([]interface{}, len(fields)*len(values))
	for i, vi := range values {
		start := i * len(fields)
		for j, vj := range f(vi) {
			args[start+j] = vj
		}
	}
	return
}

type UpdateSet[T any] struct {
	UpdateFiled   string
	UpdateSign    string
	WhereFiled    string
	WhereSign     string
	FuncUpdateSet func(v T) []interface{}
}

// 批量更新(set方式)(建议是使用id主键或者唯一索引的方式）
func UpdateDifferentSet[T any](tableName string, list []T, updateList []UpdateSet[T]) (sql string, args []interface{}) {
	var sb strings.Builder
	args = make([]interface{}, 0)
	for _, e := range updateList {
		if e.WhereSign == "" {
			e.WhereSign = "?"
		}
		if e.UpdateSign == "" {
			e.UpdateSign = "?"
		}
		sb.WriteString(fmt.Sprintf("\n %s = CASE %s \n", e.UpdateFiled, e.WhereFiled))
		for _, v := range list {
			sb.WriteString(fmt.Sprintf(" WHEN %s THEN %s \n", e.WhereSign, e.UpdateSign))
			args = append(args, e.FuncUpdateSet(v)...)
		}
		sb.WriteString("END,")
	}
	caseWhen := strings.TrimSuffix(sb.String(), ",")
	sql = fmt.Sprintf("UPDATE %s SET %s \n", tableName, caseWhen)
	return
}

// 批量更新(必须包含主键字段以及非空字段，有则更新，无则插入，无主键则插入)（不建议使用）
func UpdateByKey[T any](tableName string, fields []string, values []T, f func(v T) []interface{}) (sql string, args []interface{}) {
	repeat := "(" + strings.TrimSuffix(strings.Repeat("?,", len(fields)), ",") + "),"
	repeats := strings.TrimSuffix(strings.Repeat(repeat, len(values)), ",")
	var sb strings.Builder
	for _, field := range fields {
		sb.WriteString(fmt.Sprintf(" %s = values(%s),", field, field))
	}
	UpdateFiled := strings.TrimSuffix(sb.String(), ",")
	sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES %s ON DUPLICATE KEY UPDATE %s ", tableName, strings.Join(fields, ","), repeats, UpdateFiled)
	args = make([]interface{}, len(fields)*len(values))
	for i, vi := range values {
		start := i * len(fields)
		for j, vj := range f(vi) {
			args[start+j] = vj
		}
	}
	return
}
