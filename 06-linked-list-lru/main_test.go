package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("single value")
	c := &cache{}
	c.new(2)
	c.set("hello1", "world1")
	cacheVal, err := c.get("hello1")
	if err != nil || cacheVal != "world1" {
		t.Errorf("testing failed.")
	}
}

func TestMultiValue(t *testing.T) {
	fmt.Println("multiple value")
	c := &cache{}
	c.new(2)
	c.set("hello1", "world1")
	c.set("hello2", "world2")
	cacheVal, err := c.get("hello1")
	if err != nil || cacheVal != "world1" {
		t.Errorf("testing failed.")
	}
	cacheVal, err = c.get("hello2")
	if err != nil || cacheVal != "world2" {
		t.Errorf("testing failed.")
	}
}

func TestTripleValue(t *testing.T) {
	fmt.Println("triple value")
	c := &cache{}
	c.new(5)
	c.set("hello1", "world1")
	c.set("hello2", "world2")
	c.set("hello3", "world3")
	cacheVal, err := c.get("hello1")
	if err != nil || cacheVal != "world1" {
		t.Errorf("testing failed.")
	}
	cacheVal, err = c.get("hello2")
	if err != nil || cacheVal != "world2" {
		t.Errorf("testing failed.")
	}
	cacheVal, err = c.get("hello3")
	if err != nil || cacheVal != "world3" {
		t.Errorf("testing failed.")
	}
}

func TestResetValue(t *testing.T) {
	fmt.Println("reset same value")
	c := &cache{}
	c.new(5)
	c.set("hello1", "world1")
	c.set("hello2", "world2")
	c.set("hello3", "world3")
	c.set("hello2", "world2")
	cacheVal, err := c.get("hello1")
	if err != nil || cacheVal != "world1" {
		t.Errorf("testing failed.")
	}
	cacheVal, err = c.get("hello2")
	if err != nil || cacheVal != "world2" {
		t.Errorf("testing failed.")
	}
	cacheVal, err = c.get("hello3")
	if err != nil || cacheVal != "world3" {
		t.Errorf("testing failed.")
	}
}

func TestFullValue(t *testing.T) {
	fmt.Println("set full value")
	c := &cache{}
	c.new(5)
	c.set("hello1", "world1")
	c.set("hello2", "world2")
	c.set("hello3", "world3")
	c.set("hello4", "world4")
	c.set("hello5", "world5")
	c.set("hello6", "world6")
	c.debug()
	cacheVal, err := c.get("hello6")
	if err != nil || cacheVal != "world6" {
		t.Errorf("testing failed.")
	}
	c.debug()
	cacheVal, err = c.get("hello1")
	if err == nil {
		t.Errorf("testing failed.")
	}
	c.debug()
	cacheVal, err = c.get("hello2")
	if err != nil || cacheVal != "world2" {
		t.Errorf("testing failed.")
	}
	c.debug()
	cacheVal, err = c.get("hello3")
	if err != nil || cacheVal != "world3" {
		t.Errorf("testing failed.")
	}
	c.debug()
	cacheVal, err = c.get("hello5")
	if err != nil || cacheVal != "world5" {
		t.Errorf("testing failed.")
	}
	c.debug()
}
