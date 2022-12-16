package __StatusMode

import "fmt"

// BatteryState 电池状态接口，支持手机充电线插拔事件
type BatteryState interface {
	ConnectPlug(iPhone *IPhone) string
	DisconnectPlug(iPhone *IPhone) string
}

// fullBatteryState 满电状态
type fullBatteryState struct{}

func (s *fullBatteryState) String() string {
	return "满电状态"
}

func (s *fullBatteryState) ConnectPlug(iPhone *IPhone) string {
	return iPhone.pauseCharge()
}

func (s *fullBatteryState) DisconnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(PartBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.consume(), s, PartBatteryState)
}

// emptyBatteryState 空电状态
type emptyBatteryState struct{}

func (s *emptyBatteryState) String() string {
	return "没电状态"
}

func (s *emptyBatteryState) ConnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(PartBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.charge(), s, PartBatteryState)
}

func (s *emptyBatteryState) DisconnectPlug(iPhone *IPhone) string {
	return iPhone.shutdown()
}

// partBatteryState 部分电状态
type partBatteryState struct{}

func (s *partBatteryState) String() string {
	return "有电状态"
}

func (s *partBatteryState) ConnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(FullBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.charge(), s, FullBatteryState)
}

func (s *partBatteryState) DisconnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(EmptyBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.consume(), s, EmptyBatteryState)
}
