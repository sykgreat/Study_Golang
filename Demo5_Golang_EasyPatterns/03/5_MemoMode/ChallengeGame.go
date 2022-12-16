package __MemoMode

import "fmt"

// Originator 备忘录模式原发器接口
type Originator interface {
	Save(tag string) Memento // 当前状态保存备忘录
}

// RolesPlayGame 支持存档的RPG游戏
type RolesPlayGame struct {
	name          string   // 游戏名称
	rolesState    []string // 游戏角色状态
	scenarioState string   // 游戏场景状态
}

// NewRolesPlayGame 根据游戏名称和角色名，创建RPG游戏
func NewRolesPlayGame(name string, roleName string) *RolesPlayGame {
	return &RolesPlayGame{
		name:          name,
		rolesState:    []string{roleName, "血量100"}, // 默认满血
		scenarioState: "开始通过第一关",                   // 默认第一关开始
	}
}

// Save 保存RPG游戏角色状态及场景状态到指定标签归档
func (r *RolesPlayGame) Save(tag string) Memento {
	return newRPGArchive(tag, r.rolesState, r.scenarioState, r)
}

func (r *RolesPlayGame) SetRolesState(rolesState []string) {
	r.rolesState = rolesState
}

func (r *RolesPlayGame) SetScenarioState(scenarioState string) {
	r.scenarioState = scenarioState
}

// String 输出RPG游戏简要信息
func (r *RolesPlayGame) String() string {
	return fmt.Sprintf("在%s游戏中，玩家使用%s,%s,%s;", r.name, r.rolesState[0], r.rolesState[1], r.scenarioState)
}
