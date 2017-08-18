package main

import ()

type InstanceCreator interface {
	Create(instanceID string) error
	Destroy(instanceID string) error
	InstanceExists(instanceID string) (bool, error)
}

type creator struct {
	instances map[string]string
}

func NewCreator() *creator {
	var c creator
	c.instances = make(map[string]string)
	return &c
}

func (c *creator) Create(instanceID string) error {
	c.instances[instanceID] = instanceID
	return nil
}

func (c *creator) Destroy(instanceID string) error {
	delete(c.instances, instanceID)
	return nil
}

func (c *creator) InstanceExists(instanceID string) (bool, error) {
	_, exists := c.instances[instanceID]
	return exists, nil
}
