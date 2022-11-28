package module

func DataStruct() {

	/*
		//创建栈
		stack := make([]int, 0)
		//push压入栈
		stack = append(stack, 10)
		//pop弹出
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//检查栈空
		len(stack) == 0
	*/

	/*
		//创建队列
		queue := make([]int, 0)
		//enqueue入队
		queue = append(queue, 10)
		//dequeue出队
		v := queue[0]
		queue = queue[1:]
		//检查队列为空
		len(queue) == 0
	*/

	/*
		// int排序
		sort.Ints([]int{})
		// 字符串排序
		sort.Strings([]string{})
		// 自定义排序
		sort.Slice(s,func(i,j int)bool{return s[i]<s[j]})
	*/

	/*
		// int32 最大最小值
		math.MaxInt32 // 实际值：1<<31-1
		math.MinInt32 // 实际值：-1<<31
		// int64 最大最小值（int默认是int64）
		math.MaxInt64
		math.MinInt64
	*/

	/*
		// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
		copy(a[i:],a[i+1:])
		a=a[:len(a)-1]

		// make创建长度，则通过索引赋值
		a:=make([]int,n)
		a[n]=x
		// make长度为0，则通过append()赋值
		a:=make([]int,0)
		a=append(a,x)
	*/

	/*
		// byte转数字
		s := "12345"                        // s[0] 类型是byte
		num := int(s[0] - '0')              // 1
		str := string(s[0])                 // "1"
		b := byte(num + '0')                // '1'
		fmt.Printf("%d%s%c\n", num, str, b) // 111

		// 字符串转数字
		num, _ := strconv.Atoi("123")

		//数字转字符串
		str := strconv.Itoa(123)
	*/
}
