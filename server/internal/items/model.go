package items

import "github.com/darkalit/rlzone/server/pkg/pagination"

type Item struct {
	ID         uint    `gorm:"column:id;primaryKey"`
	Name       string  `gorm:"column:name"`
	Image      string  `gorm:"column:image"`
	Type       string  `gorm:"column:type"`
	Quality    string  `gorm:"column:quality"`
	Hitbox     *string `gorm:"column:hitbox"        json:",omitempty"`
	Reactive   *bool   `gorm:"column:reactive"      json:",omitempty"`
	TradeIn    bool    `gorm:"column:trade_in"`
	Paintable  bool    `gorm:"column:paintable"`
	Blueprints bool    `gorm:"column:blueprints"`
	Released   string  `gorm:"column:released"`
	Platform   string  `gorm:"column:platform"`
	Sideswipe  string  `gorm:"column:sideswipe"`
	Series     string  `gorm:"column:series"`
	Stock      *Stock  `                            json:",omitempty"`
}

type Stock struct {
	ID          uint   `gorm:"primaryKey"`
	Price       uint   `gorm:"column:price"`
	Description string `gorm:"column:description"`
	ItemID      uint   `gorm:"column:item_id"`
}

type GetItemsQuery struct {
	ID          uint   `json:"id"            form:"id"`
	Sort        string `json:"sort"          form:"sort"`
	MinPrice    uint   `json:"min_price"     form:"min_price"`
	MaxPrice    uint   `json:"max_price"     form:"max_price"`
	OnlyInStock bool   `json:"only_in_stock" form:"only_in_stock"`
	Type        string `json:"type"          form:"type"`
	Quality     string `json:"quality"       form:"quality"`
	Page        int    `json:"page"          form:"page"`
	PageSize    int    `json:"page_size"     form:"page_size"`
}

type GetResponse struct {
	Items      []Item
	Pagination pagination.Pagination
}

type CreateStockRequest struct {
	Price       uint   `json:"price"       binding:"required"`
	Description string `json:"description" binding:"required"`
	ItemID      uint   `json:"item_id"     binding:"required"`
}
