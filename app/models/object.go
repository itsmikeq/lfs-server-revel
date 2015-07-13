package models
import (
	"github.com/revel/revel"
	"time"
	"io"
)

type Object struct {
	Oid            string `db:"oid" json:"oid"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time
	Content io.Reader
}

func (o *Object) Validate(v *revel.Validation) {
	// all OIDs are 65 characters long
	v.Check(o.Oid,
		revel.ValidRequired(),
		revel.ValidMaxSize(65))
}

func FindObject(oid string) (*Object) {
	RedisClient.Get(oid)
	return &Object{Oid: oid}
}

func SaveObject(oid string) (*Object) {
	RedisClient.Set(oid, 1, -1)
	return &Object{Oid: oid, CreatedAt: time.Now()}
}
