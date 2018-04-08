package model

// Component
type Component struct {
	Id     int64                  `json:"id" xorm:"not null pk autoincr INT(10)"`
	Type   string                 `json:"type" xorm:"varchar(20)"`
	Data   map[string]interface{} `json:"data" xorm:"json"`
	Design map[string]interface{} `json:"design" xorm:"json"` // 暂时放map, 后面再优化性能

	CreatedAt int64 `json:"created_at" xorm:"int(11) created"`
	UpdatedAt int64 `json:"updated_at" xorm:"int(11) updated"`
}

func GetComponent(id int64) (bool, *Component, error) {
	item := Component{}
	exist, err := engine.ID(id).Get(&item)
	if err != nil {
		return false, nil, err
	}

	return exist, &item, nil
}

// 根据多个id获取item
func GetComponentByIds(ids []int64) (map[int64]*Component, error) {
	items := map[int64]*Component{}
	if len(ids)==0{
		return items,nil
	}
	err := engine.In("id", ids).Find(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// 删除Item
func DelComponent(id int64) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	_, err := engine.ID(id).Delete(new(Component))
	if err != nil {
		return err
	}

	return nil
}

// 更新Item
func UpdateComponent(id int64, s *Component) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	if s.Type == "" {
		return ErrBadParams.Append("Type can't be empty")
	}
	exist, _, err := GetComponent(s.Id)
	if err != nil {
		return err
	}
	if !exist {
		return ErrNotFind.Append("没找到item")
	}

	col := []string{"data", "design"}
	_, err = engine.ID(s.Id).Cols(col...).Update(s)
	if err != nil {
		return err
	}

	return nil
}

// 添加item
func CreateComponent(s *Component) (error) {
	if s.Type == "" {
		return ErrBadParams.Append("type can't be empty")
	}
	_, err := engine.Insert(s)
	if err != nil {
		return err
	}

	return nil
}

// 添加item
func CreateComponents(s []*Component) (error) {
	it := make([]interface{}, len(s))
	for i := range s {
		it[i] = s[i]
	}
	_, err := engine.Insert(it...)
	if err != nil {
		return err
	}

	return nil
}
