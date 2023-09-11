package models


type Inventory_SKU struct {
	Sku      string  `json:"sku" bson:"sku"`
	Price    Price   `json:"price" bson:"price"`
	Quantity float32 `json:"quantity" bson:"quantity,truncate"`
	Options  Options `json:"options" bson:"options"`
}

type Price struct {
	Base     float32 `json:"base" bson:"base,truncate"`
	Currency string  `json:"currency" bson:"currency"`
	Discount float32 `json:"discount" bson:"discount,truncate"`
}
type Options struct {
	Size     Size     `json:"size" bson:"size"`
	Features []string `json:"features" bson:"features"`
	Colors   []string `json:"colors" bson:"colors"`
	Ruling   string   `json:"ruling" bson:"ruling"`
	Image    string   `json:"image" bson:"image"`
}

type Size struct {
	H float32 `json:"h" bson:"h,truncate"`
	L float32 `json:"l" bson:"l,truncate"`
	W float32 `json:"w" bson:"w,truncate"`
}

type UpdatedInventory struct {
	Sku      string  `json:"sku" bson:"sku"`
	Quantity float32 `json:"quantity" bson:"quantity,truncate"`
}
