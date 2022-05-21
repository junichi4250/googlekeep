// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package server

import (
	"fmt"
	"io"
	"strconv"

	"github.com/volatiletech/authboss/v3"
)

type Label struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Todos  []*Todo `gorm:"many2many:todos_labels"` // many-to-many
	UserID string  `sql:"type:TEXT REFERENCES users(id) ON DELETE CASCADE"`
}

type LabelAction struct {
	Action Action `json:"action"`
	Label  *Label `json:"label"`
}

type Note struct {
	ID          string `gorm:"primary_key"`
	TodoID      string `sql:"type:TEXT REFERENCES todos(id) ON DELETE CASCADE"`
	Text        string `json:"text"`
	IsCompleted bool   `json:"isCompleted"`
}

type NotesInput struct {
	Text        string `json:"text"`
	IsCompleted bool   `json:"isCompleted"`
}

type Todo struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Notes          []*Note  `json:"notes" gorm:"foreignkey:TodoID"`       // has-many
	Labels         []*Label `json:"labels" gorm:"many2many:todos_labels"` // many-to-many
	Color          string   `json:"color"`
	IsCheckboxMode bool     `json:"isCheckboxMode"`
	UserID         string   `sql:"type:TEXT REFERENCES users(id) ON DELETE CASCADE"`
}

type TodoAction struct {
	Action Action `json:"action"`
	Todo   *Todo  `json:"todo"`
}

type User struct {
	authboss.ArbitraryUser
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	ListMode bool     `json:"listMode"`
	DarkMode bool     `json:"darkMode"`
	Todos    []*Todo  `gorm:"foreignkey:UserID"` // has-many
	Labels   []*Label `gorm:"foreignkey:UserID"` // has-many
}

func (u *User) GetPID() string {
	return u.ID
}

func (u *User) PutPID(pid string) {
	u.ID = pid
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) PutPassword(password string) {
	u.Password = password
}

func (u *User) GetArbitrary() map[string]string {
	return map[string]string{
		"name":  u.Name,
		"email": u.Email,
	}
}

func (u *User) PutArbitrary(values map[string]string) {
	if name, ok := values["name"]; ok {
		u.Name = name
	}
	if email, ok := values["email"]; ok {
		u.Email = email
	}
}

func (u *User) Validate() []error {
	return nil
}

type Action string

const (
	ActionCreated Action = "CREATED"
	ActionDeleted Action = "DELETED"
	ActionUpdated Action = "UPDATED"
)

var AllAction = []Action{
	ActionCreated,
	ActionDeleted,
	ActionUpdated,
}

func (e Action) IsValid() bool {
	switch e {
	case ActionCreated, ActionDeleted, ActionUpdated:
		return true
	}
	return false
}

func (e Action) String() string {
	return string(e)
}

func (e *Action) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Action(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Action", str)
	}
	return nil
}

func (e Action) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
