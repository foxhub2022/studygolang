package main

import (
	"fmt"
)

type  student struct {
	id 		int64
	name 	string
}
func newStudent(id int64,name string) *student {
	return &student{
		id:		id,
		name: 	name,
	}
}
var (
	allStudent map[int64]*student //变量声明
)

func showAllStudent() {
	// 把所有学生都打印出来
	for k,v := range allStudent {
		fmt.Printf("学号：%d 姓名： %s\n",k,v.name)
	}
}

func addStudent()  {
	// 向allStudent中添加一个学生
	// 1. 创建一个新学生
	// 1.1.获取用户输入
	var (
	 	id 	int64
	 	name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	// 1.2 造学生（调用student的构造函数）
	newStu := newStudent(id, name)
	// 2. 追加allStudent这个map中
	allStudent[id] = newStu
}

func deletestudent(){
	var (
		deleteID int64
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&deleteID)
	// 2. 去allstudent这个map里根据学号删除对应的键值对
	delete(allStudent, deleteID)
}


func main() {
	allStudent = make(map[int64]*student,40) // 初始化（开辟内存空间

	for {
		var (
			name string
			age  int
		)
		fmt.Print("输入姓名和年龄，使用空格分隔：")
		fmt.Scanln(&name, &age)
		fmt.Printf("name: %s\nage: %d\n", name, age)


		// 1.打印菜单
		//fmt.Println("欢迎光临学生管理系统！")
		//fmt.Println(`
		//1.查看所有学生
		//2.新增学生
		//3.删除学生
		//4.退出
		//`)
		//fmt.Print("请输入你要干啥：")
		//// 2.等待用户选择要做什么
		//var choice int
		//fmt.Scan(choice)
		//fmt.Scanln(choice)
		//fmt.Printf("你选择了%d这个选择！\n", choice)
		////3. 执行对应函数
		//switch choice {
		//	case 1:
		//		showAllStudent()
		//	case 2:
		//		addStudent()
		//	case 3:
		//		deletestudent()
		//	case 4:
		//		os.Exit(1)
		//	default:
		//		fmt.Println("滚~")
		//}


	}



}
