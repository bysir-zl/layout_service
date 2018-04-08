package model

import (
	"log"
	"strconv"
	"strings"
)

// item, 包括普通组件与布局组件
type Item struct {
	Id     string                 `json:"id"`
	Type   string                 `json:"type"`
	Data   map[string]interface{} `json:"data,omitempty"`
	Design map[string]interface{} `json:"design,omitempty"`

	Name   string      `json:"name,omitempty"`
	Layout *LayoutItem `json:"layout,omitempty"`

	Items map[string]*Item `json:"items,omitempty"`
}

type LayoutItem struct {
	Id       string       `json:"i"`
	Children []LayoutItem `json:"c,omitempty"`
}

// GetLayout 获取一个布局，包含这个布局组件所需要的所有item
func GetLayoutPage(id int64) (l *Item, err error) {
	exist, p, err := GetPage(id)
	if err != nil {
		return
	}
	if !exist {
		err = ErrNotFind.Append("page")
		return
	}

	i := strconv.FormatInt(p.Id, 10)
	item, err := GetLayout(i, p.Layout)
	if err != nil {
		return
	}

	return item, nil
}

// GetLayout 获取布局组件， 包含前端渲染需要的所有数据
func GetLayout(layoutId string, layout *LayoutItem) (item *Item, err error) {
	if layout == nil {
		return nil, ErrDefault.Append("组件没有布局")
	}
	itemIds, quoteIds := collectIds(layout)
	cs, err := GetComponentByIds(itemIds)
	if err != nil {
		return
	}
	cqs, err := GetComponentQuoteByIds(quoteIds)
	if err != nil {
		return
	}

	items := make(map[string]*Item, len(cs)+len(cqs))
	for k, v := range cs {
		id := COMPONENT_BASE_PREFIX + "_" + strconv.FormatInt(k, 10)
		items[id] = &Item{
			Id:     id,
			Type:   v.Type,
			Data:   v.Data,
			Design: v.Design,
		}
	}
	for k, v := range cqs {
		id := COMPONENT_QUOTE_PREFIX + "_" + strconv.FormatInt(k, 10)
		it, err := GetLayout(id, v.Layout)
		if err != nil {
			log.Printf("GetLayout err:%v", err)
			continue
		}
		items[id] = it
	}

	item = &Item{
		Id:     layoutId,
		Type:   "layout",
		Layout: layout,
		Items:  items,
	}

	return
}

const (
	COMPONENT_BASE_PREFIX  = "cb"
	COMPONENT_QUOTE_PREFIX = "cq"
)

// 收集这个layout所需要的所有item, 可能是从component表获取的, 可能是component-quote表获取的.
// 根据id的前缀判断是那个表
func collectIds(layout *LayoutItem) ([]int64, []int64) {
	var itemIds []int64
	var quoteIds []int64

	sp := strings.Split(layout.Id, "_")
	if len(sp) != 2 {
		return itemIds, quoteIds
	}
	prefix := sp[0]
	id, _ := strconv.ParseInt(sp[1], 10, 64)
	switch prefix {
	case COMPONENT_QUOTE_PREFIX:
		quoteIds = append(quoteIds, id)
	case COMPONENT_BASE_PREFIX:
		itemIds = append(itemIds, id)
	default:
		itemIds = append(itemIds, id)
	}

	for _, v := range layout.Children {
		i, l := collectIds(&v)
		itemIds = append(itemIds, i...)
		quoteIds = append(quoteIds, l...)
	}

	return itemIds, quoteIds
}
