// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
//
// This file is part of ewm-cloud-robotics
// (see https://github.com/SAP/ewm-cloud-robotics).
//
// This file is licensed under the Apache Software License, v. 2 except as noted
// otherwise in the LICENSE file (https://github.com/SAP/ewm-cloud-robotics/blob/master/LICENSE)
//

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/apis/ewm/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RobotConfigurationLister helps list RobotConfigurations.
type RobotConfigurationLister interface {
	// List lists all RobotConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RobotConfiguration, err error)
	// RobotConfigurations returns an object that can list and get RobotConfigurations.
	RobotConfigurations(namespace string) RobotConfigurationNamespaceLister
	RobotConfigurationListerExpansion
}

// robotConfigurationLister implements the RobotConfigurationLister interface.
type robotConfigurationLister struct {
	indexer cache.Indexer
}

// NewRobotConfigurationLister returns a new RobotConfigurationLister.
func NewRobotConfigurationLister(indexer cache.Indexer) RobotConfigurationLister {
	return &robotConfigurationLister{indexer: indexer}
}

// List lists all RobotConfigurations in the indexer.
func (s *robotConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.RobotConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RobotConfiguration))
	})
	return ret, err
}

// RobotConfigurations returns an object that can list and get RobotConfigurations.
func (s *robotConfigurationLister) RobotConfigurations(namespace string) RobotConfigurationNamespaceLister {
	return robotConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RobotConfigurationNamespaceLister helps list and get RobotConfigurations.
type RobotConfigurationNamespaceLister interface {
	// List lists all RobotConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RobotConfiguration, err error)
	// Get retrieves the RobotConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RobotConfiguration, error)
	RobotConfigurationNamespaceListerExpansion
}

// robotConfigurationNamespaceLister implements the RobotConfigurationNamespaceLister
// interface.
type robotConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RobotConfigurations in the indexer for a given namespace.
func (s robotConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RobotConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RobotConfiguration))
	})
	return ret, err
}

// Get retrieves the RobotConfiguration from the indexer for a given namespace and name.
func (s robotConfigurationNamespaceLister) Get(name string) (*v1alpha1.RobotConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("robotconfiguration"), name)
	}
	return obj.(*v1alpha1.RobotConfiguration), nil
}
