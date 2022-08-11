package main

import (
	"encoding/csv"
	"net/url"
	"os"
	"strconv"

	"RedisCount/rds"
)

func main() {
	key := "dus:referrer_count:"
	countMap := make(map[string]int)
	var cursor uint64
	for {
		keys, cur, err := rds.RedisMain.HScan(key, cursor, "", 100).Result()
		if err != nil {
			break
		}

		for index := 0; index < len(keys); index += 2 {
			dom := GetDomain(keys[index])
			c, _ := strconv.Atoi(keys[index+1])
			if val, ok := countMap[dom]; ok {
				countMap[dom] = val + c
			} else {
				countMap[dom] = c
			}
		}

		if cur == 0 {
			break
		}
		cursor = cur
	}

	ImportToCsv(countMap)

}

func GetDomain(urls string) string {
	var domain string
	u, err := url.Parse(urls)
	if err != nil {
		panic(err)
	}
	if u.Scheme != "" {
		domain += u.Scheme
		domain += "://"
	}

	if u.Host != "" {
		domain += u.Host
	}

	return domain
}

func ImportToCsv(data map[string]int) {
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
		_ = writer.Write([]string{k, str})
	}
	writer.Flush()
}
