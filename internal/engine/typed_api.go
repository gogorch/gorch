package engine

// RegisterTyped 注册一个强类型对象到 Context 容器中。
func RegisterTyped[T any](ctx *Context, ins *T, replace bool) error {
	return registerTyped(ctx.container, ins, replace)
}

// MutableTyped 从 Context 容器中读取一个强类型对象。
func MutableTyped[T any](ctx *Context, val *T) error {
	return mutableTyped(ctx.container, val)
}
