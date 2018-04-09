package model

// ComponentQuote
type ComponentQuote struct {
	Id     int64       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name   string      `json:"name" xorm:"varchar(20)"`
	Layout *LayoutItem `json:"layout" xorm:"json"`

	CreatedAt int64 `json:"created_at" xorm:"int(11) created"`
}

func GetComponentQuote(id int64) (bool, *ComponentQuote, error) {
	item := ComponentQuote{}
	exist, err := engine.ID(id).Get(&item)
	if err != nil {
		return false, nil, err
	}

	return exist, &item, nil
}

// 根据多个id获取item
func GetComponentQuoteByIds(ids []int64) (map[int64]*ComponentQuote, error) {
	items := map[int64]*ComponentQuote{}
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
func DelComponentQuote(id int64) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}
	_, err := engine.ID(id).Delete(new(ComponentQuote))
	if err != nil {
		return err
	}

	return nil
}

// 更新Item
func UpdateComponentQuote(id int64, s *ComponentQuote) (error) {
	if id == 0 {
		return ErrBadParams.Append("id is empty")
	}

	exist, _, err := GetComponentQuote(s.Id)
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
func CreateComponentQuote(s *ComponentQuote) (error) {

	_, err := engine.Insert(s)
	if err != nil {
		return err
	}

	return nil
}

// 添加item
func CreateComponentQuotes(s []*ComponentQuote) (error) {
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
