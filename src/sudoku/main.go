package main

import "fmt"

//数独中的单元格
type Box struct {
	Figure uint8
	Lock   bool
	//未填数的格可以填写的数字数目
	nOp int
	//未填数的格可以填写的数字列表
	Op []int
}

//数独题
type Form [9][9]Box

var (
	Num     int
	readbuf string = "800000000003600000070090200050007000000045700000100030001000068008500010090000400"
)

func main() {
	form := new(Form)
	form.InitShudu(readbuf)
	fmt.Println("已读取数独题：")
	form.Print()
	if form.Answer() {
		fmt.Println("成功解答，共计算了", Num, "次！")
		form.Print()
	}
}

//填充Box
func (form *Form) InitShudu(buf string) {
	//计算次数初始化
	Num = 0
	//填充数列
	i := 0
	j := 0
	for _, v := range buf {
		form[i][j].Figure = uint8(v) - 48
		if v != '0' {
			form[i][j].Lock = true
		}
		j++
		if j > 8 {
			j = 0
			i++
		}
	}
}

//验证是否数独是否成立
func (form *Form) CheckAll() bool {
	row := [9]uint8{}
	//横向检查
	for _, vform := range form {
		for j, v := range vform {
			row[j] = v.Figure
		}
		if !CheckLine(row) {
			return false
		}
		row = [9]uint8{}
	}
	//竖向检查
	n := 0
	for i := 0; i < 9; i++ {
		for _, vform := range form {
			for j, v := range vform {
				if j == i {
					row[n] = v.Figure
					n++
				}
			}
		}
		if !CheckLine(row) {
			return false
		}
		row = [9]uint8{}
		n = 0
	}
	//块检查
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			nx := getx(i)
			ny := gety(j)
			for _, vx := range nx {
				for _, vy := range ny {
					row[n] = form[vx][vy].Figure
					n++
				}
			}
			if !CheckLine(row) {
				return false
			}
			row = [9]uint8{}
			n = 0
		}
	}
	return true
}

//检查某一序列的数字是否合格
func CheckLine(line [9]uint8) bool {
	for _, v := range line {
		if v == 0 {
			return false
		}
	}
	state := [9]bool{}
	for _, v := range line {
		state[v-1] = true
	}
	for _, v := range state {
		if v == false {
			return false
		}
	}
	return true
}
func getx(x int) (n [3]int) {
	switch x {
	case 0, 1, 2:
		{
			n = [3]int{0, 1, 2}
			return
		}
	case 3, 4, 5:
		{
			n = [3]int{3, 4, 5}
			return
		}
	case 6, 7, 8:
		{
			n = [3]int{6, 7, 8}
			return
		}
	}
	return
}
func gety(y int) (n [3]int) {
	switch y {
	case 0, 1, 2:
		{
			n = [3]int{0, 1, 2}
			return
		}
	case 3, 4, 5:
		{
			n = [3]int{3, 4, 5}
			return
		}
	case 6, 7, 8:
		{
			n = [3]int{6, 7, 8}
			return
		}
	}
	return
}

//填充空格的可选数字列表
func (form *Form) FillOption() {
	form.ReBox()
	for i, vform := range form {
		for j, _ := range vform {
			if !form[i][j].Lock {
				form.GetOption(i, j)
			}
		}
	}
}

//获取某个空格的可选数字列表
func (form *Form) GetOption(x, y int) {
	state := [9]bool{}
	for i := 0; i < 9; i++ {
		//横向判断
		if form[x][i].Figure != 0 {
			state[form[x][i].Figure-1] = true
		}
		//竖向判断
		if form[i][y].Figure != 0 {
			state[form[i][y].Figure-1] = true
		}
	}
	nx := getx(x)
	ny := gety(y)
	for _, vx := range nx {
		for _, vy := range ny {
			if form[vx][vy].Figure != 0 {
				state[form[vx][vy].Figure-1] = true
			}
		}
	}
	n := 0
	for i, v := range state {
		if !v {
			form[x][y].Op = append(form[x][y].Op, i+1)
			n++
		}
	}
	form[x][y].nOp = n
}

func (form *Form) ReBox() {
	for j, vform := range form {
		for i, _ := range vform {
			if !form[i][j].Lock {
				form[i][j].Figure = 0
				form[i][j].nOp = 0
				l := len(form[i][j].Op)
				//删除form[i][j].State.Option切片中的所有元素
				form[i][j].Op = append(form[i][j].Op[:0], form[i][j].Op[l:]...)
			}
		}
	}
}
func (form *Form) Getmin() (x, y int) {
	init := false
	for i, vform := range form {
		for j, _ := range vform {
			if form[i][j].Figure == 0 {
				if !init {
					x, y = i, j
					init = true
				}
				if form[i][j].nOp < form[x][y].nOp {
					x, y = i, j
				}
			}
		}
	}
	return
}
func (form *Form) End() bool {
	for _, vform := range form {
		for _, v := range vform {
			if v.Figure == 0 {
				return false
			}
		}
	}
	return true
}

//递归求值
func (form *Form) Answer() bool {
	form.FillOption()
	x, y := form.Getmin()
	var vop []int
	vop = append(vop, form[x][y].Op...)
	for _, v := range vop {
		form[x][y].Figure = uint8(v)
		form[x][y].Lock = true
		if !form.End() {
			if form.Answer() {
				return true
			} else {
				form[x][y].Lock = false
				Num++
			}
		} else {
			fmt.Println("结束")
			if form.CheckAll() {
				return true
			}
		}
	}
	return false
}

//打印结果
func (form *Form) Print() {
	for _, vform := range form {
		for _, v := range vform {
			fmt.Print(v.Figure, "  ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}
