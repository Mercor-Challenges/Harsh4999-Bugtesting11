/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
Copyright 2017 The Kubernetes Authors.

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

package server

// This file exists to force the desired plugin implementations to be linked.
// This should probably be part of some configuration fed into the build for a
// given binary target.
import (
	"k8s.io/apiserver/pkg/admission"

	// Admission controllers
	"github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/broker/authsarcheck"
	"github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/namespace/lifecycle"
	siclifecycle "github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/servicebindings/lifecycle"
	"github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/serviceplan/changevalidator"
	"github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/serviceplan/defaultserviceplan"
)

// registerAllAdmissionPlugins registers all admission plugins
func registerAllAdmissionPlugins(plugins *admission.Plugins) {
	lifecycle.Register(plugins)
	defaultserviceplan.Register(plugins)
	siclifecycle.Register(plugins)
	changevalidator.Register(plugins)
	authsarcheck.Register(plugins)
}
