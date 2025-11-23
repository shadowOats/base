package base

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

// 时间函数
func Sleep(args ...int) {
	var min, max int

	fmt.Println(len(args))

	switch len(args) {
	case 0:
		time.Sleep(time.Duration(1000) * time.Millisecond)
		return
	case 1:
		time.Sleep(time.Duration(args[0]) * time.Millisecond)
		return
	case 2:
		min, max = args[0], args[1] // 两个参数，随机范围
		if min < 0 || max < 0 {
			fmt.Printf("参数只能接收正整数")
			return
		}
		if min == max || min > max {
			time.Sleep(time.Duration(args[0]) * time.Millisecond)
			return
		}
	default:
		fmt.Println("Sleep函数仅支持 0、1、2 个参数。")
		fmt.Println("例如：base.Sleep()、base.Sleep(2000)、base.Sleep(1000,3000)")
		os.Exit(0)
		return
	}

	// 1000-3000 毫秒
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn((max - min)) + min*1

	fmt.Printf("⏳ 随机延迟 %d 毫秒...\n", n)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// 封装 input 函数
func Input(args ...string) string {
	switch len(args) {
	case 0:
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		return scanner.Text()

	case 1:
		fmt.Print(args[0])
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		return scanner.Text()
	default:
		fmt.Println("Input函数仅接收0、1个参数。")
		fmt.Println("例如：base.Input(\"请输入您的名字：\")")
		os.Exit(0)
		return ""
	}
}

// 颜色函数
func RedPrint(format string, a ...interface{}) {
	fmt.Println(color.New(color.FgRed).Sprintf(format, a...))
}
func GreenPrint(format string, a ...interface{}) {
	fmt.Println(color.New(color.FgGreen).Sprintf(format, a...))
}
func YellowPrint(format string, a ...interface{}) {
	fmt.Println(color.New(color.FgYellow).Sprintf(format, a...))
}
func PinkPrint(format string, a ...interface{}) {
	fmt.Println(color.New(color.FgMagenta).Sprintf(format, a...))
}

func RedStr(a ...interface{}) string    { return color.New(color.FgHiMagenta).Sprint(a...) }
func GreenStr(a ...interface{}) string  { return color.New(color.FgGreen).Sprint(a...) }
func YellowStr(a ...interface{}) string { return color.New(color.BgYellow).Sprint(a...) }
func PinkStr(a ...interface{}) string   { return color.New(color.FgHiMagenta).Sprint(a...) }

// 创建目录，如果路径中包含文件名，则创建上级目录
func Mkdir(path string) {
	dirPath := path
	if !strings.HasSuffix(path, string(os.PathSeparator)) && !strings.Contains(filepath.Base(path), ".") {
		dirPath = filepath.Dir(path)
	}

	if dirPath != "" {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}
}

// 追加内容到文件，如果目录不存在则创建
func AppendFile(filePath string, fileContent interface{}) {
	dirPath := filepath.Dir(filePath)
	Mkdir(dirPath) // Ensure the directory exists

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	switch content := fileContent.(type) {
	case []string:
		for _, line := range content {
			file.WriteString(line + "\n")
		}
	case string:
		file.WriteString(content + "\n")
	}
}

// 写入内容到文件，如果目录不存在则创建
func WriteFile(filePath string, fileContent interface{}) {
	dirPath := filepath.Dir(filePath)
	Mkdir(dirPath) // Ensure the directory exists

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	switch content := fileContent.(type) {
	case []string:
		for _, line := range content {
			file.WriteString(line + "\n")
		}
	case string:
		file.WriteString(content + "\n")
	}
}

// 读取文件内容
func ReadFile(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(data), "\n")
}

// 获取当前时间
func NowTime() string {
	return time.Now().Format("20060102-150405")
}

// 发包（get、post）
