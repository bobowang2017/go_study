package main

import "fmt"

/*出行方式接口*/
type ITravel interface {
	Starting() //起点
	Travel()   //出行方式
	End()      //终点
}

//定义父类
type Person struct {
	Name    string  //姓名
	ITravel ITravel //出行方式接口
}

func (p *Person) SetName(name string) {
	p.Name = name
}

//算法骨架
func (p Person) Start() {
	if p.ITravel == nil {
		fmt.Println("ITravel is nil")
		return
	}
	fmt.Printf("%s", p.Name)
	p.ITravel.Starting()
	p.ITravel.Travel()
	p.ITravel.End()
	fmt.Print("\n")
}

//具体类：张三
type ZhangSan struct {
	Person //匿名字段实现继承
}

func NewZhangSan() *ZhangSan {
	c := new(ZhangSan)
	c.Person = Person{ITravel: c}
	return c
}

func (_ ZhangSan) Starting() {
	fmt.Printf("从%s出发", "北京")
}

func (_ ZhangSan) Travel() {
	fmt.Printf(" %s", "火车")
}
func (_ ZhangSan) End() {
	fmt.Printf(" %s", "到达上海")
}

//具体类：李四
type LiSi struct {
	Person
}

func (_ LiSi) Starting() {
	fmt.Printf("从%s出发", "广州")
}

func (_ LiSi) Travel() {
	fmt.Printf(" %s", "坐飞机")
}
func (_ LiSi) End() {
	fmt.Printf(" %s", "到达重庆")
}

func main() {

	//张三
	zhangsan := NewZhangSan()
	zhangsan.Start()

	persion := &Person{}
	persion.ITravel = ZhangSan{}
	persion.Start()

	//李四
	persion.ITravel = &LiSi{}
	persion.Start()

	lisi := LiSi{}
	lisi.ITravel = lisi
	lisi.Start()
}
