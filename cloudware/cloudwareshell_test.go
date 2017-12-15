package cloudware

import (
	"testing"
	"fmt"
)

func TestDriver_AddUser(t *testing.T) {
	d, err := Init(); if err != nil {
		t.Error("driver init failed")
	}

	err = d.AddUser("aa"); if err != nil {
		t.Error("AddUser aa failed:%v", err)
	}
}

func TestDriver_DelUser(t *testing.T) {
	d, err := Init(); if err != nil {
		t.Error("driver init failed")
	}

	err = d.DelUser("aa"); if err != nil {
		t.Error("DelUser aa failed:%v", err)
	}

}

func TestDriver_ListUser(t *testing.T) {
	d, err := Init(); if err != nil {
		t.Error("driver init failed")
	}

	users, err := d.ListUser(); if err != nil {
		t.Error("DelUser aa failed:%v", err)
	}

	for user := range users.Users {
		fmt.Printf("user:%v", user)
	}
}
