package pkg

import (
	"github.com/speps/go-hashids/v2"
)

func HashId(id uint) string {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(id)})
	return e
}
