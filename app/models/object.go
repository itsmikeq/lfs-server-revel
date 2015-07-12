package models
import (
	"github.com/revel/revel"
	"time"
)

type Object struct {
	Id              int64   `db:"id" json:"id"`
	Oid            string  `db:"oid" json:"oid"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
}

func (o *Object) Validate(v *revel.Validation) {
	// all OIDs are 65 characters long
	v.Check(o.Oid,
		revel.ValidRequired(),
		revel.ValidMaxSize(65))
}