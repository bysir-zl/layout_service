package model

// Page
type Page struct {
	Id     int64       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name   string      `json:"name" xorm:"varchar(20)"`
	Layout *LayoutItem `json:"layout" xorm:"json"`
	UserId int64       `json:"user_id" xorm:"INT(11)"`
	SiteId int64       `json:"site_id" xorm:"INT(11)"`

	CreatedAt int64 `json:"created_at" xorm:"int(11) created"`
	UpdatedAt int64 `json:"updated_at" xorm:"int(11) updated"`
}

func GetPage(id int64) (bool, *Page, error) {
	page := Page{}
	exist, err := engine.ID(id).Get(&page)
	if err != nil {
		return false, nil, err
	}

	return exist, &page, nil
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

// 添加Page
func CreatePage(s *Page) (error) {
	_, err := engine.Insert(s)
	if err != nil {
		return err
	}

	return nil
}
