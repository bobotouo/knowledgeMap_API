package repositories

import (
	"bobo/models"
	"fmt"
)

func tempStr(str1, str2 string) string {
	var temp string
	if len(str2) > 0 {
		temp = fmt.Sprintf("%s(%s)", str1, str2)
	} else {
		temp = str1
	}
	return temp
}

// / 查询所有
func GetAllData(entityId string) []models.ApiAllResults {
	var data models.Tomaha
	var list []models.Tomaha
	Orm.Where("value = ?", entityId).First(&data)
	Orm.Where("entityId = ? ", data.EntityId).Find(&list)

	entiInfo := models.Entity{
		Eid:         data.EntityId,
		RealName:    data.Entity,
		Name:        tempStr(data.Entity, data.EntityType),
		SemanticTag: data.EntityType,
	}

	var all []models.ApiAllResults

	for index, nd := range list {

		rel := models.Relation{
			Name:       nd.Property,
			RelationId: "",
		}
		val := models.ValueNode{
			Name:     tempStr(nd.Value, nd.ValueType),
			RealName: nd.Value,
			ValId:    nd.ValueId,
		}

		model := models.ApiAllResults{
			EntityInfo:  []models.Entity{entiInfo},
			Relation:    []models.Relation{rel},
			SemanticTag: nd.ValueType,
			Val:         []models.ValueNode{val},
			Groups:      nd.Group,
			Status:      false,
		}

		found := false

		/// 去重
		for _, v := range all {
			if v.Val[0].ValId == nd.ValueId {
				found = true
				break
			}
		}

		// 继续查询下层查看是否有数据
		if index == 0 {
			all = append(all, model)
		} else {
			var next []models.Tomaha
			Orm.Where("entityId = ?", nd.ValueId).Find(&next)
			if len(next) == 0 {
				all = append(all, model)
			} else if next[0].EntityId != "" && next[0].ValueId == "" && !found {
				all = append(all, model)
			}
		}

	}

	return all
}

// 查询节点内
func QueryNode(entityId string) []models.ApiAllResults {
	var list []models.Tomaha
	var results []models.ApiAllResults
	var first models.Tomaha
	Orm.Where("entityId = ?", entityId).Find(&list)

	for index, nd := range list {

		if index == 0 {
			first.EntityId = nd.EntityId
			first.Entity = nd.Entity
			first.EntityType = nd.EntityType
			first.Group = nd.Group
			first.Property = nd.Property
			first.Source = nd.Source
			first.Value = nd.Value
			first.ValueId = nd.ValueId
			first.ValueType = nd.ValueType
		}

		entiInfo := models.Entity{
			Eid:         first.EntityId,
			RealName:    first.Entity,
			Name:        tempStr(first.Entity, first.EntityType),
			SemanticTag: first.EntityType,
		}

		rel := models.Relation{
			Name:       nd.Property,
			RelationId: "",
		}
		val := models.ValueNode{
			Name:     tempStr(nd.Value, nd.ValueType),
			RealName: nd.Value,
			ValId:    nd.ValueId,
		}

		model := models.ApiAllResults{
			EntityInfo:  []models.Entity{entiInfo},
			Relation:    []models.Relation{rel},
			SemanticTag: nd.ValueType,
			Val:         []models.ValueNode{val},
			Groups:      nd.Group,
			Status:      false,
		}

		found := false

		/// 去重
		for _, v := range results {
			if v.Val[0].ValId == nd.ValueId {
				found = true
				break
			}
		}

		if !found {
			results = append(results, model)
		}

	}

	return results
}

// // 搜索内容
func SearchContent(search string) []models.Tomaha {

	var list []models.Tomaha
	if len(search) == 0 {
		return list
	}

	Orm.Where("entity like ? and property = '七巧板医学术语集概念ID' AND entityType != '药品'", search+"%").Find(&list)

	var temps []models.Tomaha
	for _, v := range list {
		var count int64
		/// 查询下一层
		Orm.Model(&models.Tomaha{}).Where("entityId = ?", v.EntityId).Count(&count)
		if count > 2 {
			temps = append(temps, v)
		}
	}

	return temps
}
