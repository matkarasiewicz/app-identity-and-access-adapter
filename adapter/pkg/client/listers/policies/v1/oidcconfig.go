/*
Copyright 2019 The Kubernetes Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/ibm-cloud-security/policy-enforcer-mixer-adapter/adapter/pkg/apis/policies/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OidcConfigLister helps list OidcConfigs.
type OidcConfigLister interface {
	// List lists all OidcConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1.OidcConfig, err error)
	// OidcConfigs returns an object that can list and get OidcConfigs.
	OidcConfigs(namespace string) OidcConfigNamespaceLister
	OidcConfigListerExpansion
}

// oidcConfigLister implements the OidcConfigLister interface.
type oidcConfigLister struct {
	indexer cache.Indexer
}

// NewOidcConfigLister returns a new OidcConfigLister.
func NewOidcConfigLister(indexer cache.Indexer) OidcConfigLister {
	return &oidcConfigLister{indexer: indexer}
}

// List lists all OidcConfigs in the indexer.
func (s *oidcConfigLister) List(selector labels.Selector) (ret []*v1.OidcConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.OidcConfig))
	})
	return ret, err
}

// OidcConfigs returns an object that can list and get OidcConfigs.
func (s *oidcConfigLister) OidcConfigs(namespace string) OidcConfigNamespaceLister {
	return oidcConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OidcConfigNamespaceLister helps list and get OidcConfigs.
type OidcConfigNamespaceLister interface {
	// List lists all OidcConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.OidcConfig, err error)
	// Get retrieves the OidcConfig from the indexer for a given namespace and name.
	Get(name string) (*v1.OidcConfig, error)
	OidcConfigNamespaceListerExpansion
}

// oidcConfigNamespaceLister implements the OidcConfigNamespaceLister
// interface.
type oidcConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OidcConfigs in the indexer for a given namespace.
func (s oidcConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.OidcConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.OidcConfig))
	})
	return ret, err
}

// Get retrieves the OidcConfig from the indexer for a given namespace and name.
func (s oidcConfigNamespaceLister) Get(name string) (*v1.OidcConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("oidcconfig"), name)
	}
	return obj.(*v1.OidcConfig), nil
}
