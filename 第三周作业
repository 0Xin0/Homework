package main

import (
	"fmt"
)

type commodity struct {
	Name      string
	Price     float64
	Inventory int
} //商品结构体
type inventory_control interface {
	check()
	renew()
	print()
} //库存管理接口

func (c commodity) check() string {
	result := fmt.Sprintf("%s 的库存为 %v 件", c.Name, c.Inventory)
	return result
} //普通商品库存检查
func (c *commodity) renew(newinventory int) string {
	c.Inventory = newinventory
	result := fmt.Sprintf("%v 更新后的库存为 %v 件", c.Name, c.Inventory)
	return result
} //普通商品库存更新
func (c commodity) print() string {
	result := fmt.Sprintf("产品名：%v 价格：%v 元 库存数量：%v 件", c.Name, c.Price, c.Inventory)
	return result
} //普通商品信息打印

type electronic_goods struct {
	commodity
	Brand, Model string
} //电子产品结构体
type eg_inventory_control interface {
	echeck()
	erenew()
	eprint()
} //电子产品库存管理接口

func (e electronic_goods) echeck() string {
	result := fmt.Sprintf("%v 的库存为 %v 件", e.Name, e.Inventory)
	return result
} //电子产品库存检查
func (e *electronic_goods) erenew(newinventory int) string {
	e.Inventory = newinventory
	result := fmt.Sprintf("%v 更新后的库存为 %v 件", e.Name, e.Inventory)
	return result
} //电子产品库存更新
func (e electronic_goods) eprint() string {
	result := fmt.Sprintf("产品名：%s 品牌：%s 型号：%s 价格：%v 元 库存数量：%v 件", e.Name, e.Brand, e.Model, e.Price, e.Inventory)
	return result
} //电子产品信息打印

func main() {
	c := commodity{
		"被子",
		50,
		10,
	}
	e := electronic_goods{
		commodity{
			"手机",
			6888,
			5,
		},
		"华为",
		"mate 60 Pro",
	}
	fmt.Println(c)
	fmt.Println(e)
	c.renew(5)
	e.renew(0)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Printf("%s\n", c.check())
	fmt.Printf("%s\n", c.print())
	fmt.Printf("%s\n", e.echeck())
	fmt.Printf("%s\n", e.eprint())
}
