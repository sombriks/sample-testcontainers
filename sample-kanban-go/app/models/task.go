package models

import "time"

type Task struct {
	Id          int64      `db:"id" goqu:"skipinsert"`
	StatusId    int64      `db:"status_id"`
	Description string     `db:"description"`
	Created     *time.Time `db:"created" goqu:"omitnil,skipupdate"`
	Status      *Status
	Messages    *[]Message
	People      *[]Person
}

func (t *Task) SafeMessageCount() int {
	if t.Messages != nil {
		return len(*t.Messages)
	}
	return 0
}

func (t *Task) SafeMessageList() *[]Message {
	if t.Messages == nil {
		return &[]Message{}
	}
	return t.Messages
}

func (t *Task) SafePeopleCount() int {
	if t.People != nil {
		return len(*t.People)
	}
	return 0
}

func (t *Task) SafePeopleList() *[]Person {
	if t.People == nil {
		return &[]Person{}
	}
	return t.People
}

func (t *Task) MemberById(id int64) *Person {
	for _, person := range *t.SafePeopleList() {
		if person.Id == id {
			return &person
		}
	}
	return nil
}
