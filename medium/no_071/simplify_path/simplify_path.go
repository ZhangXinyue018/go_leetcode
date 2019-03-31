package simplify_path

func simplifyPath(path string) string {
	path = path + "/"
	stack := &Stack{}
	temp := []rune{}
	tempStr := ""
	for _, value := range path{
		if value == '/'{
			tempStr = string(temp)
			if tempStr == ".."{
				stack.Pop()
			}else if tempStr != "" && tempStr != "."{
				stack.Push(tempStr)
			}
			temp = []rune{}
		}else{
			temp = append(temp, value)
		}
	}
	if len(stack.Items) == 0{
		return "/"
	}
	result := ""
	for _, value := range stack.Items{
		result = result + "/" + value
	}
	return result
}

type Stack struct{
	Items []string
}

func (stack *Stack) Push(str string){
	stack.Items = append(stack.Items, str)
}

func (stack *Stack) Pop(){
	if len(stack.Items) == 0{
		return
	}
	stack.Items = stack.Items[:len(stack.Items)-1]
}
