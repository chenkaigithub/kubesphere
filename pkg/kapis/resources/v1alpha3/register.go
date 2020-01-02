/*

 Copyright 2019 The KubeSphere Authors.

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
package v1alpha3

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/api/resource/v1alpha2"
	"kubesphere.io/kubesphere/pkg/apiserver/query"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
	"kubesphere.io/kubesphere/pkg/simple/client/k8s"
	"net/http"
)

const (
	GroupName = "resources.kubesphere.io"

	tagComponentStatus    = "Component Status"
	tagNamespacedResource = "Namespaced Resource"

	ok = "OK"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}

func AddWebService(c *restful.Container, client k8s.Client) error {

	webservice := runtime.NewWebService(GroupVersion)
	handler := New(client)

	webservice.Route(webservice.GET("/namespaces/{namespace}/{resources}").
		To(handler.handleGetNamespacedResource).
		Metadata(restfulspec.KeyOpenAPITags, []string{tagNamespacedResource}).
		Doc("Namespace level resource query").
		Param(webservice.PathParameter("namespace", "the name of the project")).
		Param(webservice.PathParameter("resources", "namespace level resource type, e.g. pods,jobs,configmaps,services.")).
		Param(webservice.QueryParameter(query.ParameterName, "name used to do filtering").Required(false)).
		Param(webservice.QueryParameter(query.ParameterPage, "page").Required(false).DataFormat("page=%d").DefaultValue("page=0")).
		Param(webservice.QueryParameter(query.ParameterLimit, "limit").Required(false)).
		Param(webservice.QueryParameter(query.ParameterAscending, "sort parameters, e.g. reverse=true").Required(false).DefaultValue("ascending=false")).
		Param(webservice.QueryParameter(query.ParameterOrderBy, "sort parameters, e.g. orderBy=createTime")).
		Returns(http.StatusOK, ok, api.ListResult{}))

	webservice.Route(webservice.GET("/components").
		To(handler.handleGetComponents).
		Metadata(restfulspec.KeyOpenAPITags, []string{tagComponentStatus}).
		Doc("List the system components.").
		Returns(http.StatusOK, ok, []v1alpha2.ComponentStatus{}))
	webservice.Route(webservice.GET("/components/{component}").
		To(handler.handleGetComponentStatus).
		Metadata(restfulspec.KeyOpenAPITags, []string{tagComponentStatus}).
		Doc("Describe the specified system component.").
		Param(webservice.PathParameter("component", "component name")).
		Returns(http.StatusOK, ok, v1alpha2.ComponentStatus{}))
	webservice.Route(webservice.GET("/componenthealth").
		To(handler.handleGetSystemHealthStatus).
		Metadata(restfulspec.KeyOpenAPITags, []string{tagComponentStatus}).
		Doc("Get the health status of system components.").
		Returns(http.StatusOK, ok, v1alpha2.HealthStatus{}))

	c.Add(webservice)

	return nil
}
