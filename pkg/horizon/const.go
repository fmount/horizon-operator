/*

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

package horizon

import (
	"github.com/openstack-k8s-operators/lib-common/modules/storage"
)

const (
	// ServiceName -
	ServiceName = "horizon"

	// DatabaseName -
	DatabaseName = "horizon"

	// HorizonPort -
	HorizonPort int32 = 80

	// HorizonPortTLS -
	HorizonPortTLS int32 = 443

	// HorizonPortName -
	HorizonPortName = "horizon"

	// HorizonExtraVolTypeUndefined can be used to label an extraMount which is
	// not associated to anything in particular
	HorizonExtraVolTypeUndefined storage.ExtraVolType = "Undefined"
	// Horizon is the global ServiceType that refers to all the components deployed
	// by the horizon-operator
	Horizon storage.PropagationType = "Horizon"

	// HorizonUID
	HorizonUID int64 = 42420
)

// HorizonPropagation is the  definition of the Horizon propagation service
var HorizonPropagation = []storage.PropagationType{Horizon}
