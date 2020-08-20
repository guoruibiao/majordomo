package elasticsearch

import (
	"encoding/json"
	"fmt"
	"github.com/guoruibiao/gorequests"
	dao "github.com/guoruibiao/majordomo/models/dao/elasticsearch"
	"time"
)

type ElasticSearcher struct {

}


const (
	ELASTICSEARCH_URL   = "http://localhost:9200"
	ELASTICSEARCH_INDEX = "clipboard"
	ELASTICSEARCH_TYPE  = "logs"
)

func AddLog(content string) (err error) {
	logid:= time.Now().Unix()
	url := fmt.Sprintf("%s/%s/%s/%d", ELASTICSEARCH_URL, ELASTICSEARCH_INDEX, ELASTICSEARCH_TYPE, logid)


	headers := map[string]string{
		"User-Agent"   : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36",
		"Content-Type" : "application/json",
	}
	payload := map[string]interface{}{
		"content"    : content,
		"timestamp"  : logid,
	}
	response, err := gorequests.NewRequest("POST", url).Headers(headers).Body(payload).DoRequest()
	if err != nil {
		return
	}

	html, err := response.Content()
	if err != nil {
		return
	}

	fmt.Println(html)
	return
}

func GetLog(logid int) (logItem dao.LogItem, err error){
	url := fmt.Sprintf("%s/%s/%s/%d", ELASTICSEARCH_URL, ELASTICSEARCH_INDEX, ELASTICSEARCH_TYPE, logid)

	headers := map[string]string{
		"User-Agent"   : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36",
		"Content-Type" : "application/json",
	}
	response, err := gorequests.NewRequest("GET", url).Headers(headers).DoRequest()
	if err != nil {
		return
	}

	html, err := response.Content()
	if err != nil {
		return
	}

	if unmarshalErr := json.Unmarshal([]byte(html), &logItem); unmarshalErr!= nil {
		err = unmarshalErr
		return
	}

	return
}

func DeleteLog(logid int) (err error) {
	url := fmt.Sprintf("%s/%s/%s/%d", ELASTICSEARCH_URL, ELASTICSEARCH_INDEX, ELASTICSEARCH_TYPE, logid)

	headers := map[string]string{
		"User-Agent"   : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36",
		"Content-Type" : "application/json",
	}
	response, err := gorequests.NewRequest("DELETE", url).Headers(headers).DoRequest()
	if err != nil {
		return
	}

	html, err := response.Content()
	if err != nil {
		return
	}

	fmt.Println(html)
	return
}

func Search(content, ext string) (container dao.LogList, err error){
	url := fmt.Sprintf("%s/%s/%s/%s", ELASTICSEARCH_URL, ELASTICSEARCH_INDEX, ELASTICSEARCH_TYPE, "_search")

	if content == "" && ext == "" {
		return
	}

	searchConditions := make([]map[string]map[string]string, 0)
	if content != "" {
		searchConditions = append(searchConditions, map[string]map[string]string{
			"wildcard": map[string]string{
				"content": fmt.Sprintf("*%s*", content),
			},
		})
	}

	if ext != "" {
		searchConditions = append(searchConditions, map[string]map[string]string{
			"wildcard": map[string]string{
				"content": fmt.Sprintf("*%s*", ext),
			},
		})
	}

	from := 0
	size := 100
	payload := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					map[string]interface{}{
						"bool": map[string]interface{}{
							"must": []map[string]interface{}{
								map[string]interface{}{
									"bool": map[string]interface{}{
										"should": searchConditions,
									},
								},
							},
						},
					},
				},
			},
		},
		"_source": map[string][]string{
			"include": []string{"*"},
		},
		"sort": []map[string]map[string]string{
			map[string]map[string]string{
				"timestamp": map[string]string{
					"order": "DESC",
				},
			},
		},
	}

	headers := map[string]string{
		"User-Agent"   : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36",
		"Content-Type" : "application/json",
	}
	response, err := gorequests.NewRequest("POST", url).Headers(headers).Body(payload).DoRequest()
	if err != nil {
		return
	}

	html, err := response.Content()
	if err != nil {
		return
	}

	fmt.Println(html)
	if unmarshalErr := json.Unmarshal([]byte(html), &container); unmarshalErr!=nil {
		return
	}

	return
}