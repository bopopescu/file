package mode

// student首字母小写是不能被外面调用的,所以要用工厂模式
type student struct {
	Name  string
	Score float64
}

// 工厂模式.
// 注意函数与方法的区别!!! 下面这个是函数啊,不是方法!
func NewStustudentdent(name string, score float64) *student {
	return &student{
		Name:  name,
		Score: score,
	}
}
