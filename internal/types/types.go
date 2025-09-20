package types



type OperateSystem = string

const (
	Windows OperateSystem = "windows"
	Linux   OperateSystem = "linux"
	Mac     OperateSystem = "darwin"
)

type NameFunc = func(OperateSystem, string) string