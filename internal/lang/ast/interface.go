package ast

import "strings"

type DescriptorI interface {
	Description(*strings.Builder, string)
}

type LeafSnippetI interface {
	DescriptorI
	isLeafSnippet()
}

// ExedescI 执行结构描述
// 包含serial、concurrent、switch、wrap等等
type ExedescI interface {
	DescriptorI
	isExedesc()
}
