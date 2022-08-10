package main

import (
	"encoding/csv"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := "https://blog.csdn.net/qq_21514303/article/details/84619957"
	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	//直接访问 scheme。
	fmt.Println(u.Scheme)
	//User 包含了所有的认证信息，这里调用 Username和 Password 来获取独立值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	//Host 同时包括主机名和端口信息，如过端口存在的话，使用 strings.Split() 从 Host 中手动提取端口。
	fmt.Println(u.Host, "---")
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])
	//这里我们提出路径和查询片段信息。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	//要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。你也可以将查询参数解析为一个map。已解析的查询参数 map 以查询字符串为键，对应值字符串切片为值，所以如何只想得到一个键对应的第一个值，将索引位置设置为 [0] 就行了。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}

func GetDomain(url string) string {
	return ""
}

func ImportCsv(data map[string]int) {
	dir, _ := os.Getwd()
	f, err := os.Create(dir + "/data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//写入UTF-8 BOM，避免使用Microsoft Excel打开乱码
	_, _ = f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)
	for k, v := range data {
		str := strconv.Itoa(v)
		_ = writer.Write([]string{k, " ", str})
	}

}
