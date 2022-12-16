package __CommandMode

import "fmt"

// ElectricCooker 电饭煲
type ElectricCooker struct {
	fire     string // 火力
	pressure string // 压力
}

// SetFire 设置火力
func (e *ElectricCooker) SetFire(fire string) {
	e.fire = fire
}

// SetPressure 设置压力
func (e *ElectricCooker) SetPressure(pressure string) {
	e.pressure = pressure
}

// Run 持续运行指定时间
func (e *ElectricCooker) Run(duration string) string {
	return fmt.Sprintf("电饭煲设置火力为%s,压力为%s,持续运行%s;", e.fire, e.pressure, duration)
}

// Shutdown 停止
func (e *ElectricCooker) Shutdown() string {
	return "电饭煲停止运行。"
}
