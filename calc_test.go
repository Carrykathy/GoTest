package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 should be 3")
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Error("2*3 should be 6")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("failed")
		}
	})
}

/*
运行go test命令，可以运行当前目录下所有的*_test.go文件中的测试用例，并输出测试结果

运行go test -v ./...命令，可以运行所有的测试用例，并输出详细的测试结果

单独运行某个测试用例，可以使用go test -run="TestAdd" ./...命令，只运行TestAdd函数中的测试用例
*/

//go的测试文件和源文件必须同名，并且必须以_test.go结尾
//go语言的测试文件必须以_test.go结尾，以区分测试文件和源文件
//go语言的测试文件中，必须包含Test开头的函数，并且函数必须有两个参数，一个是t *testing.T类型，另一个是*testing.T类型的方法
//go语言的测试文件中，可以使用assert包来进行断言，assert包提供了很多断言方法，可以方便的进行测试
//go语言的测试文件中，可以使用t.Errorf()方法来输出错误信息，并中断测试流程
//go语言的测试文件中，可以使用t.Run()方法来定义子测试，t.Run()方法可以为同一个测试定义多个子测试，并分别执行
//go语言的测试文件中，可以使用t.Parallel()方法来并发执行测试，提高测试效率
//go语言的测试文件中，可以使用t.Cleanup()方法来定义清理函数，在测试结束后，会自动执行清理函数
//go语言的测试文件中，可以使用t.Helper()方法来标记为辅助函数，辅助函数不会被计入测试的覆盖率统计中
//go语言的测试文件中，可以使用t.Log()方法来输出日志信息，不会影响测试结果
//go语言的测试文件中，可以使用t.FailNow()方法来中断当前测试流程，并标记为失败
//go语言的测试文件中，可以使用t.Skip()方法来跳过当前测试，不会影响测试结果
