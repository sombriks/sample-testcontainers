package models

import (
	"fmt"
	"strings"
	"time"
)

type Person struct {
	Id      int64     `db:"id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created" goqu:"omitnil"`
}

func UserFromCookie(cookie string) *Person {
	var user Person
	var kv = strings.Split(cookie, "&")
	for _, e := range kv {
		if strings.Contains(e, "id=") {
			fmt.Sscan(strings.Replace(e, "id=", "", 1), &user.Id)
		}
		if strings.Contains(e, "name=") {
			fmt.Sscan(strings.Replace(e, "name=", "", 1), &user.Name)
		}
	}
	return &user
}

func (p *Person) UserToCookie() string {
	return fmt.Sprintf("id=%d&name=%s", p.Id, p.Name)
}
