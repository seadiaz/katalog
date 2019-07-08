package domain

import "reflect"

// Deployment ...
type Deployment struct {
	ID   string `json:",omitempty"`
	Name string `json:",omitempty"`
}

// GetID ...
func (s *Deployment) GetID() string {
	return s.ID
}

// GetType ...
func (s *Deployment) GetType() reflect.Type {
	return reflect.TypeOf(s)
}

// GetK8sResource ...
func (s *Deployment) GetK8sResource() interface{} {
	return s
}
