package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BindingResult struct {
	RandomNumber string
}

type InstanceBinder interface {
	Bind(instanceID string, bindingID string) (BindingResult, error)
	Unbind(instanceID string, bindingID string) error
	InstanceExists(instanceID string) (bool, error)
}

type binder struct {
	instances map[string]string
	generator *rand.Rand
}

func (b *binder) Bind(instanceID, bindingID string) (BindingResult, error) {
	b.instances[instanceID] = bindingID
	return BindingResult{RandomNumber: fmt.Sprintf("%d", b.generator.Int63())}, nil
}

func (b *binder) Unbind(instanceID string, bindingID string) error {
	delete(b.instances, instanceID)
	return nil
}

func (b *binder) InstanceExists(instanceID string) (bool, error) {
	_, exists := b.instances[instanceID]
	return exists, nil
}

func NewBinder() *binder {
	return &binder{
		instances: make(map[string]string),
		generator: rand.New(rand.NewSource(int64(time.Now().Nanosecond()))),
	}
}
