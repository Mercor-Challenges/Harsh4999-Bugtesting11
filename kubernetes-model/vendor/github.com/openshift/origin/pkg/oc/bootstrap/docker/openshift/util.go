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
package openshift

import (
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrorutils "k8s.io/apimachinery/pkg/util/errors"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/kubectl/resource"

	configcmd "github.com/openshift/origin/pkg/bulk"
	"github.com/openshift/origin/pkg/oc/bootstrap/docker/errors"
	"github.com/openshift/origin/pkg/oc/cli/util/clientcmd"
	genappcmd "github.com/openshift/origin/pkg/oc/generate/cmd"
	templateinternalclient "github.com/openshift/origin/pkg/template/client/internalversion"
	templateclient "github.com/openshift/origin/pkg/template/generated/internalclientset/typed/template/internalversion"
)

func instantiateTemplate(client templateclient.TemplateInterface, clientFactory *clientcmd.Factory, templateNamespace, templateName, targetNamespace string, params map[string]string, ignoreExistsErrors bool) error {
	template, err := client.Templates(templateNamespace).Get(templateName, metav1.GetOptions{})
	if err != nil {
		return errors.NewError("cannot retrieve template %q from namespace %q", templateName, templateNamespace).WithCause(err)
	}

	// process the template
	templateProcessor := templateinternalclient.NewTemplateProcessorClient(client.RESTClient(), targetNamespace)
	result, err := genappcmd.TransformTemplate(template, templateProcessor, targetNamespace, params, false)
	if err != nil {
		return errors.NewError("cannot process template %s/%s", templateNamespace, templateName).WithCause(err)
	}

	mapper, typer := clientFactory.Object()
	if err != nil {
		return err
	}

	discoveryClient, err := clientFactory.DiscoveryClient()
	if err != nil {
		return errors.NewError("cannot process template %s/%s", templateNamespace, templateName).WithCause(err)
	}
	clientConfig, err := clientFactory.ClientConfig()
	if err != nil {
		return err
	}

	// Create objects
	bulk := &configcmd.Bulk{}

	bulk.DynamicMapper = &resource.Mapper{
		RESTMapper:   mapper,
		ObjectTyper:  typer,
		ClientMapper: resource.ClientMapperFunc(clientFactory.UnstructuredClientForMapping),
	}
	bulk.Mapper = &resource.Mapper{
		RESTMapper:   mapper,
		ObjectTyper:  typer,
		ClientMapper: configcmd.ClientMapperFromConfig(clientConfig),
	}
	bulk.PreferredSerializationOrder = configcmd.PreferredSerializationOrder(discoveryClient)
	bulk.Op = configcmd.Create

	itemsToCreate := &kapi.List{
		Items: result.Objects,
	}
	if errs := bulk.Run(itemsToCreate, targetNamespace); len(errs) > 0 {
		filteredErrs := []error{}
		for _, err := range errs {
			if kerrors.IsAlreadyExists(err) && ignoreExistsErrors {
				continue
			}
			filteredErrs = append(filteredErrs, err)
		}
		if len(filteredErrs) == 0 {
			return nil
		}
		err = kerrorutils.NewAggregate(filteredErrs)
		return errors.NewError("cannot create objects from template %s/%s", templateNamespace, templateName).WithCause(err)
	}

	return nil
}
