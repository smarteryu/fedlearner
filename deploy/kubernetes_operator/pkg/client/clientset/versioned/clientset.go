/* Copyright 2020 The FedLearner Authors. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	fedlearnerv1alpha1 "github.com/bytedance/fedlearner/deploy/kubernetes_operator/pkg/client/clientset/versioned/typed/fedlearner.k8s.io/v1alpha1"
	fedlearnerv1alpha2 "github.com/bytedance/fedlearner/deploy/kubernetes_operator/pkg/client/clientset/versioned/typed/fedlearner.k8s.io/v1alpha2"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	FedlearnerV1alpha1() fedlearnerv1alpha1.FedlearnerV1alpha1Interface
	FedlearnerV1alpha2() fedlearnerv1alpha2.FedlearnerV1alpha2Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	fedlearnerV1alpha1 *fedlearnerv1alpha1.FedlearnerV1alpha1Client
	fedlearnerV1alpha2 *fedlearnerv1alpha2.FedlearnerV1alpha2Client
}

// FedlearnerV1alpha1 retrieves the FedlearnerV1alpha1Client
func (c *Clientset) FedlearnerV1alpha1() fedlearnerv1alpha1.FedlearnerV1alpha1Interface {
	return c.fedlearnerV1alpha1
}

// FedlearnerV1alpha2 retrieves the FedlearnerV1alpha2Client
func (c *Clientset) FedlearnerV1alpha2() fedlearnerv1alpha2.FedlearnerV1alpha2Interface {
	return c.fedlearnerV1alpha2
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("Burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.fedlearnerV1alpha1, err = fedlearnerv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.fedlearnerV1alpha2, err = fedlearnerv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.fedlearnerV1alpha1 = fedlearnerv1alpha1.NewForConfigOrDie(c)
	cs.fedlearnerV1alpha2 = fedlearnerv1alpha2.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.fedlearnerV1alpha1 = fedlearnerv1alpha1.New(c)
	cs.fedlearnerV1alpha2 = fedlearnerv1alpha2.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
