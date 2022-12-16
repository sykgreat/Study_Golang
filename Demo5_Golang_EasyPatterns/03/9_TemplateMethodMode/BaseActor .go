package __TemplateMethodMode

import (
	"bytes"
	"fmt"
)

// IActor 演员接口
type IActor interface {
	DressUp() string // 装扮
}

// dressBehavior 装扮的多个行为，这里多个行为是私有的，通过DressUp模版方法调用
type dressBehavior interface {
	makeUp() string // 化妆
	clothe() string // 穿衣
	wear() string   // 配饰
}

// BaseActor 演员基类
type BaseActor struct {
	roleName      string // 扮演角色
	dressBehavior        // 装扮行为
}

// DressUp 统一实现演员接口的DressUp模版方法，装扮过程通过不同装扮行为进行扩展
func (b *BaseActor) DressUp() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("扮演%s的", b.roleName))
	buf.WriteString(b.makeUp())
	buf.WriteString(b.clothe())
	buf.WriteString(b.wear())
	return buf.String()
}
