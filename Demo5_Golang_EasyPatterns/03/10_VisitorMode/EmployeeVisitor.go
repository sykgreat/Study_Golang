package _0_VisitorMode

import (
	"fmt"
	"sort"
)

// EmployeeVisitor 员工访问者接口
type EmployeeVisitor interface {
	VisitProductManager(pm *productManager)     // 访问产品经理
	VisitSoftwareEngineer(se *softwareEngineer) // 访问软件工程师
	VisitHR(hr *hr)                             // 访问人力资源
}

// kpi kpi对象
type kpi struct {
	name string // 完成kpi姓名
	sum  int    // 完成kpi总数量
}

// kpiTopVisitor 员工kpi排名访问者
type kpiTopVisitor struct {
	top []*kpi
}

func (k *kpiTopVisitor) VisitProductManager(pm *productManager) {
	k.top = append(k.top, &kpi{
		name: pm.name,
		sum:  pm.productNum + pm.satisfaction,
	})
}

func (k *kpiTopVisitor) VisitSoftwareEngineer(se *softwareEngineer) {
	k.top = append(k.top, &kpi{
		name: se.name,
		sum:  se.requirementNum + se.bugNum,
	})
}

func (k *kpiTopVisitor) VisitHR(hr *hr) {
	k.top = append(k.top, &kpi{
		name: hr.name,
		sum:  hr.recruitNum,
	})
}

// Publish 发布KPI排行榜
func (k *kpiTopVisitor) Publish() {
	sort.Slice(k.top, func(i, j int) bool {
		return k.top[i].sum > k.top[j].sum
	})
	for i, curKPI := range k.top {
		fmt.Printf("第%d名%s：完成KPI总数%d\n", i+1, curKPI.name, curKPI.sum)
	}
}

// salaryVisitor 薪酬访问者
type salaryVisitor struct{}

func (s *salaryVisitor) VisitProductManager(pm *productManager) {
	fmt.Printf("产品经理基本薪资：1000元，KPI单位薪资：100元，")
	fmt.Printf("%s，总工资为%d元\n", pm.KPI(), (pm.productNum+pm.satisfaction)*100+1000)
}

func (s *salaryVisitor) VisitSoftwareEngineer(se *softwareEngineer) {
	fmt.Printf("软件工程师基本薪资：1500元，KPI单位薪资：80元，")
	fmt.Printf("%s，总工资为%d元\n", se.KPI(), (se.requirementNum+se.bugNum)*80+1500)
}

func (s *salaryVisitor) VisitHR(hr *hr) {
	fmt.Printf("人力资源基本薪资：800元，KPI单位薪资：120元，")
	fmt.Printf("%s，总工资为%d元\n", hr.KPI(), hr.recruitNum*120+800)
}
