package model

import (
	"encoding/json"
	"log"
)

// Page
type Page struct {
	Id     int64  `json:"id" xorm:"not null pk autoincr INT(10)"`
	Type   string `json:"type" xorm:"varchar(20)"`
	Data   string `json:"data" xorm:"text"`
	Layout string `json:"layout" xorm:"text"`
	UserId int64  `json:"user_id" xorm:"INT(11)"`
	SiteId int64  `json:"site_id" xorm:"INT(11)"`

	CreatedAt int64 `json:"created_at" xorm:"int(11) created"`
	UpdatedAt int64 `json:"updated_at" xorm:"int(11) updated"`
}

type LayoutItem struct {
	Id       int64        `json:"i"`
	Children []LayoutItem `json:"c,omitempty"`
	Layout   bool         `json:"layout,omitempty"`
}

func GetPage(id int64) (bool, *Page, error) {
	Page := Page{}
	exist, err := engine.ID(id).Get(&Page)
	if err != nil {
		return false, nil, err
	}

	return exist, &Page, nil
}

// 根据多个id获取Page
func GetPageByIds(ids []int64) (map[int64]*Page, error) {
	Pages := map[int64]*Page{}
	err := engine.In("id", ids).Find(&Pages)
	if err != nil {
		return nil, err
	}

	return Pages, nil
}

// 删除Page
func DelPage(id int64) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	_, err := engine.ID(id).Delete(new(Page))
	if err != nil {
		return err
	}

	return nil
}

// 更新Page
func UpdatePage(id int64, s *Page) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	exist, _, err := GetPage(s.Id)
	if err != nil {
		return err
	}
	if !exist {
		return ErrNotFind.Append("没找到Page")
	}

	col := []string{"data", "layout"}
	_, err = engine.ID(s.Id).Cols(col...).Update(s)
	if err != nil {
		return err
	}

	return nil
}

type Layout struct {
	Id      int64             `json:"id"`
	Layout  LayoutItem        `json:"layout"`
	Items   map[int64]*Item   `json:"items"`
	Layouts map[int64]*Layout `json:"layouts"`
}

func collectIds(layout *LayoutItem) ([]int64, []int64) {
	var itemIds []int64
	var layoutIds []int64

	if layout.Layout {
		layoutIds = append(layoutIds, layout.Id)
	} else {
		itemIds = append(itemIds, layout.Id)
	}
	for _, v := range layout.Children {
		i, l := collectIds(&v)
		itemIds = append(itemIds, i...)
		layoutIds = append(layoutIds, l...)
	}

	return itemIds, layoutIds
}

func GetLayoutPage(id int64) (l *Layout, err error) {
	exist, p, err := GetPage(id)
	if err != nil {
		return
	}
	if !exist {
		err = ErrNotFind.Append("page")
		return
	}
	layout := LayoutItem{}
	json.Unmarshal([]byte(p.Layout), &layout)

	itemIds,layoutIds := collectIds(&layout)
	items, err := GetItemByIds(itemIds)
	if err != nil {
		return
	}

	log.Print(layoutIds)
	l = &Layout{
		Id:     id,
		Layout: layout,
		Items:  items,
	}

	return
}

// 添加Page
func CreatePage(s *Page) (error) {
	if s.Type == "" {
		return ErrBadParams.Append("type can't be empty")
	}
	_, err := engine.Insert(s)
	if err != nil {
		return err
	}

	return nil
}
