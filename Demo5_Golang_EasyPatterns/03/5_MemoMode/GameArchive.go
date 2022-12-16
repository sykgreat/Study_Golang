package __MemoMode

import "fmt"

// Memento 备忘录接口
type Memento interface {
	Tag() string // 备忘录标签
	Restore()    // 根据备忘录存储数据状态恢复原对象
}

// rpgArchive rpg游戏存档，
type rpgArchive struct {
	tag           string         // 存档标签
	rolesState    []string       // 存档的角色状态
	scenarioState string         // 存档游戏场景状态
	rpg           *RolesPlayGame // rpg游戏引用
}

// newRPGArchive 根据标签，角色状态，场景状态，rpg游戏引用，创建游戏归档备忘录
func newRPGArchive(tag string, rolesState []string, scenarioState string, rpg *RolesPlayGame) *rpgArchive {
	return &rpgArchive{
		tag:           tag,
		rolesState:    rolesState,
		scenarioState: scenarioState,
		rpg:           rpg,
	}
}

func (r *rpgArchive) Tag() string {
	return r.tag
}

// Restore 根据归档数据恢复游戏状态
func (r *rpgArchive) Restore() {
	r.rpg.SetRolesState(r.rolesState)
	r.rpg.SetScenarioState(r.scenarioState)
}

// RPGArchiveManager RPG游戏归档管理器
type RPGArchiveManager struct {
	archives map[string]Memento // 存储归档标签对应归档
}

func NewRPGArchiveManager() *RPGArchiveManager {
	return &RPGArchiveManager{
		archives: make(map[string]Memento),
	}
}

// Reload 根据标签重新加载归档数据
func (r *RPGArchiveManager) Reload(tag string) {
	if archive, ok := r.archives[tag]; ok {
		fmt.Printf("重新加载%s;\n", tag)
		archive.Restore()
	}
}

// Put 保存归档数据
func (r *RPGArchiveManager) Put(memento Memento) {
	r.archives[memento.Tag()] = memento
}
