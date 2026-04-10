package engine

// RuntimeStats 运行时观测数据快照。
type RuntimeStats struct {
	// FallbackDIExecCount 反射注入/导出路径执行次数。
	FallbackDIExecCount uint64
	// GeneratedExecCount 生成算子路径执行次数。
	GeneratedExecCount uint64
	// StrictRejectExecCount 严格模式下拒绝执行次数。
	StrictRejectExecCount uint64
}

// SnapshotRuntimeStats 返回当前运行时观测快照。
func SnapshotRuntimeStats() RuntimeStats {
	return RuntimeStats{
		FallbackDIExecCount:   fallbackDIExecCount.Load(),
		GeneratedExecCount:    generatedExecCount.Load(),
		StrictRejectExecCount: strictRejectExecCount.Load(),
	}
}

func resetRuntimeStats() {
	fallbackDIExecCount.Store(0)
	generatedExecCount.Store(0)
	strictRejectExecCount.Store(0)
}
