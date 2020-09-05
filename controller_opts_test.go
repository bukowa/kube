package kube

import (
	"errors"
	"strconv"
	"testing"
)

// hooks are registered with proper order
func TestOptWithHooksFirst(t *testing.T) {
	c := NewController(NewContainer(),
		WithHooks(Hooks{
			PreCreate: []Hook{
				func(c Container) error { return errors.New("2") },
				func(c Container) error { return errors.New("3") },
			}}),
		WithHooksFirst(Hooks{
			PreCreate: []Hook{
				func(c Container) error { return errors.New("1") },
			}}),
		WithHooks(Hooks{
			PreCreate: []Hook{
				func(c Container) error { return errors.New("4") },
			}}),
		WithHooksFirst(Hooks{
			PreCreate: []Hook{
				func(c Container) error { return errors.New("0") },
			}}),
	)
	if len(c.Hooks()) != 1 {
		t.Error()
	}
	for _, h := range c.Hooks() {
		if len(h) != 5 {
			t.Error()
		}
		for i, f := range h {
			if f(c).Error() != strconv.Itoa(i) {
				t.Error(h, i, f(c).Error())
			}
		}
	}
}
