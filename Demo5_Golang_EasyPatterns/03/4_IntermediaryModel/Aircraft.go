package __IntermediaryModel

import "fmt"

// Aircraft 飞机接口
type Aircraft interface {
	ApproachAirport() // 抵达机场空域
	DepartAirport()   // 飞离机场
}

// airliner 客机
type airliner struct {
	name            string          // 客机型号
	airportMediator AirportMediator // 机场调度
}

// NewAirliner 根据指定型号及机场调度创建客机
func NewAirliner(name string, mediator AirportMediator) *airliner {
	return &airliner{
		name:            name,
		airportMediator: mediator,
	}
}

func (a *airliner) ApproachAirport() {
	if !a.airportMediator.CanLandAirport(a) { // 请求塔台是否可以降落
		fmt.Printf("机场繁忙，客机%s继续等待降落;\n", a.name)
		return
	}
	fmt.Printf("客机%s成功滑翔降落机场;\n", a.name)
}

func (a *airliner) DepartAirport() {
	fmt.Printf("客机%s成功滑翔起飞，离开机场;\n", a.name)
	a.airportMediator.NotifyWaitingAircraft() // 通知等待的其他飞机
}

// helicopter 直升机
type helicopter struct {
	name            string
	airportMediator AirportMediator
}

// NewHelicopter 根据指定型号及机场调度创建直升机
func NewHelicopter(name string, mediator AirportMediator) *helicopter {
	return &helicopter{
		name:            name,
		airportMediator: mediator,
	}
}

func (h *helicopter) ApproachAirport() {
	if !h.airportMediator.CanLandAirport(h) { // 请求塔台是否可以降落
		fmt.Printf("机场繁忙，直升机%s继续等待降落;\n", h.name)
		return
	}
	fmt.Printf("直升机%s成功垂直降落机场;\n", h.name)
}

func (h *helicopter) DepartAirport() {
	fmt.Printf("直升机%s成功垂直起飞，离开机场;\n", h.name)
	h.airportMediator.NotifyWaitingAircraft() // 通知其他等待降落的飞机
}
