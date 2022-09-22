package models
type User struct {
	Id     int    `json:"id"`
	Name  string `json:"name"`
	// DOB primitive.DateTime `bson:"dob,omitempty" json:"dob,omitempty"`
	Address string `json:address`
	Desc   string `json:"desc"`
	// createdAt primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}