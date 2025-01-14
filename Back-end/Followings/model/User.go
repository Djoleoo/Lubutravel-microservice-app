package model

import (
	"encoding/json"
	"io"
)

type User struct {
	UserName string `json:"Username"`
}
type Followed []*User

type Users []*User

func (o *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *Followed) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}
