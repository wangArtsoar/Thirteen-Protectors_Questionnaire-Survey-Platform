package project

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"testing"
	"time"
)

//type label struct {
//	Id             int       `json:"id"`
//	CreatedAt      time.Time `json:"created_at"`
//	UpdatedAt      time.Time `json:"updated_at"`
//	RepoFileId     string    `json:"repo_file_id"`
//	FunctionId     string    `json:"function_id"`
//	Signature      string    `json:"signature"`
//	LabelStatus    int       `json:"label_status"`
//	EnglishComment string    `json:"english_comment"`
//	ChineseComment string    `json:"chinese_comment"`
//	ForceUpdate    bool      `json:"force_update"`
//	ReviewedAt     time.Time `json:"reviewed_at"`
//	CheckApprove   int       `json:"check_approve"`
//	CheckReason    string    `json:"check_reason"`
//	CreatedBy      string    `json:"created_by"`
//	ChineseSummary string    `json:"chinese_summary"`
//	EnglishSummary string    `json:"english_summary"`
//	LabeledAt      time.Time `json:"labeled_at"`
//	Ext            string    `json:"ext"`
//	ChatgptContent string    `json:"chatgpt_content"`
//	ReviewBatch    string    `json:"review_batch"`
//}

func TestCountCodeLebelOn0816(t *testing.T) {
	url := "https://codelabel.tencent.com/api/function/label-by-self"
	// 创建HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	cookies := []*http.Cookie{
		{
			Name:  "_ga_6WSZ0YS5ZQ",
			Value: "GS1.1.1687913212.2.0.1687913212.0.0.0",
		},
		{
			Name:  "sensorsdata2015jssdkcross",
			Value: "%7B%22distinct_id%22%3A%221892520c35b63c-0a8f7c860143a28-7e565473-3686400-1892520c35c852%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E8%87%AA%E7%84%B6%E6%90%9C%E7%B4%A2%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTg5MjUyMGMzNWI2M2MtMGE4ZjdjODYwMTQzYTI4LTdlNTY1NDczLTM2ODY0MDAtMTg5MjUyMGMzNWM4NTIifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%22%2C%22value%22%3A%22%22%7D%2C%22%24device_id%22%3A%221892520c35b63c-0a8f7c860143a28-7e565473-3686400-1892520c35c852%22%7D"},
		{
			Name:  "_ga",
			Value: "GA1.2.1334629688.1687855736",
		},
		{
			Name:  "_gcl_au",
			Value: "1.1.1587327710.1688545053",
		},
		{
			Name: "token",
			//Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb2RlbGFiZWwiLCJzdWIiOiJ4aWFveWlfd3l4QGljbG91ZC5jb20iLCJleHAiOjE2OTIyMzYyMTcsInVzZXJfbmFtZSI6InhpYW95aV93eXhAaWNsb3VkLmNvbSJ9.yJ-wCMZjd5rvldOYdgX98z3L8XCWHydt0F96QCbHJKY",
			//Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb2RlbGFiZWwiLCJzdWIiOiIxODYwNzA0ODY0OEAxNjMuY29tIiwiZXhwIjoxNjkyMjM0MjUzLCJ1c2VyX25hbWUiOiIxODYwNzA0ODY0OEAxNjMuY29tIn0.oNP7skuAxzKO8CPci7JQkH7Lc9lBQhb_EIVBfWs6Djs",
			Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb2RlbGFiZWwiLCJzdWIiOiJ4aWFveWlfd3l4QGljbG91ZC5jb20iLCJleHAiOjE2OTIzMzYyMTcsInVzZXJfbmFtZSI6InhpYW95aV93eXhAaWNsb3VkLmNvbSJ9.PeLKA2oMRrPt8c44PTewUc3eiAWdc9fUn_trbOhdn0o",
		},
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	// 创建HTTP客户端
	client := &http.Client{}
	// 发送HTTP GET请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 解析JSON数据
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	// 输出结果
	res := data["data"].([]interface{})
	var goodCount, midCount, chaCount int
	fmt.Printf("sum : %d\n", len(res))
	mapCount := make(map[time.Time]int)
	for _, re := range res {
		m := re.(map[string]interface{})
		labeledAtStr := m["labeled_at"].(string)
		checkApprove := m["check_approve"].(float64)
		switch checkApprove {
		case 5:
			goodCount++
		case 0:
			chaCount++
		case 3:
			midCount++
		}
		var labeledAt time.Time
		if labeledAt, err = time.Parse(time.RFC3339, labeledAtStr); err != nil {
			t.Error(err)
		}
		var parse time.Time
		if parse, err = time.Parse(time.DateOnly, labeledAt.Format("2006-01-02")); err != nil {
			t.Error(err)
		}
		mapCount[parse]++
	}
	type kv struct {
		key   time.Time
		Value int
	}
	var kvs []kv
	for k, v := range mapCount {
		kvs = append(kvs, kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].key.After(kvs[j].key)
	})
	fmt.Println(goodCount, midCount, chaCount)
	for _, v := range kvs {
		fmt.Printf("日期 : %s, 标注数量 : %d\n", v.key.String()[0:10], v.Value)
	}
}
