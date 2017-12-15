package cloudware

import "testing"

func TestDriver_AddUser(t *testing.T) {
	d, err := Init(); if err != nil {
		t.Error("driver init failed")
	}

	err = d.AddUser("aa"); if err != nil {
		t.Error("AddUser aa failed:%v", err)
	}
}
