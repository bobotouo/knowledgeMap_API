package models

type Entity struct {
	Eid         string `json:"eid"`
	Name        string `json:"name"`
	RealName    string `json:"realName"`
	SemanticTag string `json:"semanticTag"`
}

type Relation struct {
	Name       string `json:"name"`
	RelationId string `json:"relationId"`
}

type ValueNode struct {
	Name     string `json:"name"`
	RealName string `json:"realName"`
	ValId    string `json:"valId"`
}

type ApiAllResults struct {
	EntityInfo  []Entity    `json:"entityInfo"`
	Relation    []Relation  `json:"relation"`
	SemanticTag string      `json:"semanticTag"`
	Val         []ValueNode `json:"val"`
	Groups      string      `json:"groups"`
	Status      bool        `json:"status"`
}

type Tomaha struct {
	EntityId   string `gorm:"Column:entityId;type:varchar(255); not null;default:'';index:entity_id"`
	Entity     string `gorm:"type:varchar(255); not null;default:'';"`
	EntityType string `gorm:"Column:entityType;type:varchar(255); not null;default:'';"`
	Property   string `gorm:"type:varchar(255); not null;default:'';"`
	ValueId    string `gorm:"Column:valueId;type:varchar(255); not null;default:'';"`
	Value      string `gorm:"type:varchar(500); not null;default:'';index:val_id"`
	ValueType  string `gorm:"Column:valueType;type:varchar(255); not null;default:'';"`
	Group      string `gorm:"type:varchar(255); not null;default:'';"`
	Source     string `gorm:"type:varchar(255); not null;default:'';"`
}

func (Tomaha) TableName() string {
	return "t_omaha"
}
