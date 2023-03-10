package Models

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Webpage struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Keywords []string           `json:"keywords"`
}
type Page struct {
	Title    string   `json:"title"`
	Keywords []string `json:"keywords"`
}
type Keys struct {
	Keywords []string `json:"keywords"`
}

func (w Webpage) Check() bool {
	var title = w.Title
	title = strings.Trim(title, " ")
	if len(title) == 0 {
		return true
	}
	if w.Title == "" || len(w.Keywords) == 0 {
		return true
	}
	return false
}

// to allow not morethan 10 keywords
func (w *Webpage) ModifyKeysLength() {
	if len(w.Keywords) > 10 {
		w.Keywords = w.Keywords[:10]
	}
}
