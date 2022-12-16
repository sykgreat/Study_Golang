package DataStructure

import (
	"Study_Golang/Demo7_Golang_Blockchain/DataStructure/ArrayList"
	"Study_Golang/Demo7_Golang_Blockchain/DataStructure/Queue"
	"Study_Golang/Demo7_Golang_Blockchain/DataStructure/Stack"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestStack_Recursion1(t *testing.T) {
	stack := ArrayList.NewArrayListStack(10)
	stack.Push(5)
	result := 0
	for !stack.IsEmpty() {
		data := stack.Pop().(int)
		if data != 0 {
			result += data
			stack.Push(data - 1)
		}
	}
	t.Log(result)
}

func TestStack_Recursion2(t *testing.T) {
	stack := ArrayList.NewArrayListStack(10)
	stack.Push(3)
	result := 0
	for !stack.IsEmpty() {
		data := stack.Pop().(int)
		if data == 1 || data == 2 {
			result += 1
		} else {
			stack.Push(data - 2)
			stack.Push(data - 1)
		}
	}
	t.Log(result)
}

func Test_ErgodicFile_Stack(t *testing.T) {
	path := "E:\\code\\go\\study\\Study_Golang"
	var files []string

	stack := Stack.NewStack(10000)
	stack.Push(path)

	stackF := Stack.NewStack(10000)
	stackF.Push(path)

	level := 0

	for !stack.IsEmpty() {

		prefix := "|"
		for i := 0; i < level; i++ {
			prefix += "-"
		}

		path = stack.Pop().(string)
		files = append(files, stackF.Pop().(string))
		read, _ := ioutil.ReadDir(path)

		for _, fi := range read {
			if fi.IsDir() {
				fullFile := path + "\\" + fi.Name()
				stack.Push(fullFile)
				stackF.Push(prefix + fi.Name())
				level++
			} else {
				files = append(files, prefix+fi.Name())
			}
		}

	}
	for _, v := range files {
		t.Log(v)
	}
}

func Test_ErgodicFile_Queue(t *testing.T) {
	path := "E:\\code\\go\\study\\Study_Golang"
	var files []string

	queue := Queue.NewQueue()
	queue.EnQueue(path)

	for !queue.IsEmpty() {
		path = queue.DeQueue().(string)

		read, _ := ioutil.ReadDir(path)
		for _, fi := range read {
			if fi.IsDir() {
				fullFile := path + "\\" + fi.Name()
				queue.EnQueue(fullFile)
			} else {
				files = append(files, path+"\\"+fi.Name())
			}
		}
	}
	for _, v := range files {
		t.Log(v)
	}
}

func Test_ErgodicFile_Recursion(t *testing.T) {
	files := WalkDir(`E:\code\go\study\Study_Golang`, 1)
	for _, v := range files {
		fmt.Println(v)
	}
}

func WalkDir(filepath string, level int) []string {
	prefix := "|"
	for i := 0; i < level; i++ {
		prefix += "-----"
	}
	files, _ := ioutil.ReadDir(filepath)
	var allFile []string
	for _, v := range files {

		if v.IsDir() {
			allFile = append(allFile, prefix+v.Name())
			allFile = append(allFile, WalkDir(filepath+"\\"+v.Name(), level+1)...)
		} else {
			allFile = append(allFile, prefix+v.Name())
		}
	}

	return allFile
}
