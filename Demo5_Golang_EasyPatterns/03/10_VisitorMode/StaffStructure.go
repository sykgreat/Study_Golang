package _0_VisitorMode

import "fmt"

// Employee 员工接口
type Employee interface {
	KPI() string                    // 完成kpi信息
	Accept(visitor EmployeeVisitor) // 接受访问者对象
}

// productManager 产品经理
type productManager struct {
	name         string // 名称
	productNum   int    // 上线产品数
	satisfaction int    // 平均满意度
}

func NewProductManager(name string, productNum int, satisfaction int) *productManager {
	return &productManager{
		name:         name,
		productNum:   productNum,
		satisfaction: satisfaction,
	}
}

func (p *productManager) KPI() string {
	return fmt.Sprintf("产品经理%s，上线%d个产品，平均满意度为%d", p.name, p.productNum, p.satisfaction)
}

func (p *productManager) Accept(visitor EmployeeVisitor) {
	visitor.VisitProductManager(p)
}

// softwareEngineer 软件工程师
type softwareEngineer struct {
	name           string // 姓名
	requirementNum int    // 完成需求数
	bugNum         int    // 修复问题数
}

func NewSoftwareEngineer(name string, requirementNum int, bugNum int) *softwareEngineer {
	return &softwareEngineer{
		name:           name,
		requirementNum: requirementNum,
		bugNum:         bugNum,
	}
}

func (s *softwareEngineer) KPI() string {
	return fmt.Sprintf("软件工程师%s，完成%d个需求，修复%d个问题", s.name, s.requirementNum, s.bugNum)
}

func (s *softwareEngineer) Accept(visitor EmployeeVisitor) {
	visitor.VisitSoftwareEngineer(s)
}

// hr 人力资源
type hr struct {
	name       string // 姓名
	recruitNum int    // 招聘人数
}

func NewHR(name string, recruitNum int) *hr {
	return &hr{
		name:       name,
		recruitNum: recruitNum,
	}
}

func (h *hr) KPI() string {
	return fmt.Sprintf("人力资源%s，招聘%d名员工", h.name, h.recruitNum)
}

func (h *hr) Accept(visitor EmployeeVisitor) {
	visitor.VisitHR(h)
}
