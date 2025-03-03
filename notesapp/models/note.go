package models

import(
	"time"
)


type Note struct{
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedOn time.Time `json:"createdon"`
}