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
// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	v1alpha1 "k8s.io/api/scheduling/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	scheduling "k8s.io/kubernetes/pkg/apis/scheduling"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_PriorityClass_To_scheduling_PriorityClass,
		Convert_scheduling_PriorityClass_To_v1alpha1_PriorityClass,
		Convert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList,
		Convert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList,
	)
}

func autoConvert_v1alpha1_PriorityClass_To_scheduling_PriorityClass(in *v1alpha1.PriorityClass, out *scheduling.PriorityClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Value = in.Value
	out.GlobalDefault = in.GlobalDefault
	out.Description = in.Description
	return nil
}

// Convert_v1alpha1_PriorityClass_To_scheduling_PriorityClass is an autogenerated conversion function.
func Convert_v1alpha1_PriorityClass_To_scheduling_PriorityClass(in *v1alpha1.PriorityClass, out *scheduling.PriorityClass, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityClass_To_scheduling_PriorityClass(in, out, s)
}

func autoConvert_scheduling_PriorityClass_To_v1alpha1_PriorityClass(in *scheduling.PriorityClass, out *v1alpha1.PriorityClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Value = in.Value
	out.GlobalDefault = in.GlobalDefault
	out.Description = in.Description
	return nil
}

// Convert_scheduling_PriorityClass_To_v1alpha1_PriorityClass is an autogenerated conversion function.
func Convert_scheduling_PriorityClass_To_v1alpha1_PriorityClass(in *scheduling.PriorityClass, out *v1alpha1.PriorityClass, s conversion.Scope) error {
	return autoConvert_scheduling_PriorityClass_To_v1alpha1_PriorityClass(in, out, s)
}

func autoConvert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList(in *v1alpha1.PriorityClassList, out *scheduling.PriorityClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]scheduling.PriorityClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList is an autogenerated conversion function.
func Convert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList(in *v1alpha1.PriorityClassList, out *scheduling.PriorityClassList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList(in, out, s)
}

func autoConvert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList(in *scheduling.PriorityClassList, out *v1alpha1.PriorityClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.PriorityClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList is an autogenerated conversion function.
func Convert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList(in *scheduling.PriorityClassList, out *v1alpha1.PriorityClassList, s conversion.Scope) error {
	return autoConvert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList(in, out, s)
}
