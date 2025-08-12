package ort

import (
	"fmt"
	"reflect"
	"strings"
)

// Field 表示结构体字段的元数据，包含类型、名称、索引以及可选和替换标志。
type Field struct {
	Typ      reflect.Type
	Name     string
	Index    int
	Optional bool
	Replace  bool
}

// OperatorRType 表示算子的反射类型信息，包含类型、是否有 Prepare 方法、注入和提取字段以及所有依赖注入字段。
type OperatorRType struct {
	Typ         reflect.Type
	HasPrepare  bool
	Injects     []*Field
	Extracts    []*Field
	AllDIFields map[string]struct{}
}

// tmpPrepareOperator 定义了一个接口，包含 Prepare 方法，用于判断类型是否实现了 Prepare 方法。
type tmpPrepareOperator interface {
	Prepare() error
}

// AnalyzeOperator 分析泛型类型 T 的结构体信息，提取注入和提取字段，并判断是否实现了 Prepare 方法。
// 返回一个包含分析结果的 OperatorRType 指针。
func AnalyzeOperator[T any]() *OperatorRType {
	// 创建 T 类型的实例
	s := new(T)
	// 获取 T 类型的反射类型
	rt := reflect.TypeOf(s).Elem()

	// 初始化 OperatorRType 实例
	of := &OperatorRType{
		Typ:         rt,
		Injects:     make([]*Field, 0),
		Extracts:    make([]*Field, 0),
		AllDIFields: make(map[string]struct{}),
	}

	// 遍历结构体的所有字段
	for i := range rt.NumField() {
		field := rt.Field(i)
		// 跳过未导出的字段
		if !field.IsExported() {
			continue
		}
		// 跳过没有标签的字段
		if field.Tag == "" {
			continue
		}
		// 检查字段类型是否为指针或接口
		if field.Type.Kind() != reflect.Ptr && field.Type.Kind() != reflect.Interface {
			return handleInvalidFieldType(rt, field)
		}
		// 处理字段标签
		of.structTagFromField(i, field)
	}

	// 判断类型是否实现了 Prepare 方法
	if _, ok := any(s).(tmpPrepareOperator); ok {
		of.HasPrepare = true
	}

	return of
}

// handleInvalidFieldType 处理无效的字段类型，返回错误信息。
func handleInvalidFieldType(rt reflect.Type, field reflect.StructField) *OperatorRType {
	errMsg := fmt.Sprintf("inject/extract object only support ptr or interface: %s.%s %s",
		rt.Name(), field.Name, field.Type)
	panic(errMsg)
}

// structTagFromField 从结构体字段的标签中提取注入和提取信息。
func (a *OperatorRType) structTagFromField(idx int, field reflect.StructField) {
	// 遍历 inject 和 extract 标签
	for _, tagName := range []string{"inject", "extract"} {
		tag, has := field.Tag.Lookup(tagName)
		if !has {
			continue
		}
		// 创建 Field 实例
		st := &Field{
			Typ:   field.Type,
			Index: idx,
			Name:  field.Name,
		}
		// 解析标签选项
		parseTagOptions(tag, st)

		// 记录所有依赖注入字段
		a.AllDIFields[st.Name] = struct{}{}
		// 根据标签类型添加到相应的切片
		switch tagName {
		case "inject":
			a.Injects = append(a.Injects, st)
		case "extract":
			a.Extracts = append(a.Extracts, st)
		}
	}
}

// parseTagOptions 解析标签选项，设置 Field 实例的 Optional 和 Replace 标志。
func parseTagOptions(tag string, st *Field) {
	opts := strings.Split(tag, ",")
	for _, opt := range opts {
		switch opt {
		case "optional":
			st.Optional = true
		case "replace":
			st.Replace = true
		}
	}
}
