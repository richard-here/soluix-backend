package model

type UpdateBody struct {
	RetailPrice int  `json:"retail_price" validate:"omitempty,required,numeric,min=1,max=50000000"`
	Status      bool `json:"status" validate:"omitempty,required"`
}
