package models

import (
	"time"
	"fmt"
	"context"
	"bytes"
	"strings"
	"encoding/json"
)

type Posts struct {
	ID uint `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Content string `gorm:"type:text" json:"content"`
	Desc string `gorm:"type:text" json:"desc"`
	AddTime time.Time `gorm:"column:addTime;default:now()" json:"addTime"`
	Author int `json:"author"`
	Tags string `json:"tags"`
	Imgs string `json:"imgs"`
	View uint `json:"view"`
	NavId int `gorm:"column:navId" json:"navId"`
}

/**
 * @Author: ruke
 * @Date: 2018-12-04 15:46:40
 * @Desc: 文章过滤结构
 */
type PostFilter struct {
	NavId int
	TagId int
	Author int
	Page int
	PageSize int
	Title string
}

/**
 * @Author: ruke
 * @Date: 2018-12-04 15:45:51
 * @Desc: 获取所有列表包括过滤分页
 */
func (post *Posts) GetAll(filter *PostFilter) ([]*Posts, int, int, bool) {
	posts := []*Posts{}
	findPost := &Posts{}
	findDb := db
	if filter.NavId != 0 {
		findPost.NavId = filter.NavId
	}

	if filter.Author != 0 {
		findPost.Author = filter.Author
	}

	if filter.TagId != 0 {
		findDb = findDb.Where("find_in_set(?, tags)", filter.TagId)
	}
	count := 0
	current := filter.Page
	hasNext := false
	findDb.Where(findPost).Find(&posts).Count(&count)
	if filter.Page > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		if filter.Page * filter.PageSize < count {
			hasNext = true
		}
		findDb.Where(findPost).Order("addTime desc").Offset(offset).Limit(filter.PageSize).Find(&posts)
	}
	return posts, current, count, hasNext
}

type PostsEsInfo struct {
	Index string
	Type string
}

func (post *Posts) GetEs() PostsEsInfo {
	mapping := `
	{
		"settings":{
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
		"mappings":{
			"posts":{
				"properties":{
					"id": {
						"type": "long"
					},
					"title":{
						"type": "text",
		                "analyzer": "ik_max_word",
		                "search_analyzer": "ik_max_word"
					},
					"desc":{
						"type": "text",
		                "analyzer": "ik_max_word",
		                "search_analyzer": "ik_max_word"
					},
					"addTime":{
						"type":"date"
					},
					"tags":{
						"type": "text",
		                "analyzer": "ik_max_word",
		                "search_analyzer": "ik_max_word"
					},
					"content":{
						"type": "text",
		                "analyzer": "ik_max_word",
		                "search_analyzer": "ik_max_word"
					}
				}
			}
		}
	}`

	info := PostsEsInfo{Index: "itruke", Type: "posts"}
	ctx := context.Background()
	exists, err := esClient.IndexExists(info.Index).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := esClient.CreateIndex(info.Index).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
	
	return info
}

type PostsEs struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	AddTime uint64 `json:"addTime"`
	Tags string `json:"tags"`
}

/**
 * @Author: ruke
 * @Date: 2018-12-05 21:28:08
 * @Desc: 创建文件病写入到es
 */
func (post *Posts) Create(article Posts) Posts {
	sDb := db
	db.Create(&article)
	//存入到es
	addTime := uint64(article.AddTime.Unix()) * 1000
	//查找tags
	tags := []*Tags{}
	sDb.Where("id in (?)", strings.Split(article.Tags, ",")).Find(&tags)
	var buffer bytes.Buffer
	for _, value := range tags {
		buffer.WriteString(value.Name)
		buffer.WriteString(",")
	}
	esData := PostsEs{article.ID, article.Title, article.Desc, article.Content, addTime , strings.Trim(buffer.String(), ",")}
	jsonStr, _ := json.Marshal(esData)
	esInfo := post.GetEs()
	ctx := context.Background()
	esClient.Index().
		Index(esInfo.Index).
		Type(esInfo.Type).
		Id(fmt.Sprint(article.ID)).
		BodyString(string(jsonStr)).
		Do(ctx)
	return article 
}