package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

func main() {

	Hero1 := Hero{
		No:   1,
		Name: "Hero1",
	}

	Hero2 := Hero{
		No:   2,
		Name: "Hero2",
	}

	Hero3 := Hero{
		No:   3,
		Name: "Hero3",
	}

	Hero4 := Hero{
		No:   4,
		Name: "Hero4",
	}

	Hero5 := Hero{
		No:   5,
		Name: "Hero5",
	}

	Hero6 := Hero{
		No:   6,
		Name: "Hero6",
	}

	Hero7 := Hero{
		No:   7,
		Name: "Hero7",
	}

	Hero1.Left = &Hero2
	Hero1.Right = &Hero3
	Hero2.Left = &Hero6
	Hero2.Right = &Hero7
	Hero3.Left = &Hero4
	Hero3.Right = &Hero5

	PreOrder(&Hero1)

	fmt.Println()
	MidOrder(&Hero1)

	fmt.Println()
	PostOrder(&Hero1)
}

func PreOrder(node *Hero) {
	//前序遍历   先root  再左  再又

	if node != nil {
		fmt.Printf("no = %d name=%s", node.No, node.Name)
		fmt.Println()
		PreOrder(node.Left)
		PreOrder(node.Right)
	}

}

//先输出root的 左 再root  再root的右
func MidOrder(node *Hero) {
	if node != nil {
		MidOrder(node.Left)
		fmt.Printf("no = %d name=%s", node.No, node.Name)
		fmt.Println()

		MidOrder(node.Right)
	}
}

//后序遍历
func PostOrder(node *Hero) {
	if node != nil {
		MidOrder(node.Left)

		MidOrder(node.Right)
		fmt.Printf("no = %d name=%s", node.No, node.Name)
		fmt.Println()
	}
}
