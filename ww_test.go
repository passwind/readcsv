package main

import (
	"strconv"
	"testing"
)

func Test_F(t *testing.T) {
	convRes1, err := strconv.ParseFloat("125.56美元", 10)
	if err != nil {
		t.Errorf("err to: %s", err)
	}
	t.Logf("%+v", convRes1)
}
