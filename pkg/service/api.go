package service

import (
	"github.com/carmanzhang/ks-alert-client/pkg/handler"
	. "github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

type AlertAPI struct{}

func (u AlertAPI) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/alert/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"alert apis"}

	ws.Route(ws.POST("/alerts").To(handler.CreateAlert).
		Doc("create AlertConfig").
		Reads(AlertConfig{}).
		Writes(AlertConfigResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// get batch of alert configs by resource_name
	ws.Route(ws.GET("/alerts").To(handler.RetrieveAlert).
		Doc("retrieve AlertConfig").
		// get alert config by userid
		Param(restful.QueryParameter("user_id", "").Required(false)).
		// get alert config by binding resource_type and resource name
		Param(restful.QueryParameter("resource_type", "").Required(false)).
		Param(restful.QueryParameter("resource_name", "").Required(false)).
		// get alert config by binding super resource_type and super resource name
		// example: list all alert config in namespace 'kube-system',
		// kube-system is super_resource_name, namespace is super_resource_type
		Param(restful.QueryParameter("super_resource_type", "").Required(false)).
		Param(restful.QueryParameter("super_resource_name", "").Required(false)).
		// get alert config by alert_id
		Param(restful.QueryParameter("alert_config_id", "get alert config by id").Required(false)).
		Writes([]AlertConfigResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alerts").To(handler.UpdateAlert).
		Doc("update AlertConfig").
		Reads(AlertConfig{}).
		Writes(AlertConfigResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alerts").To(handler.DeleteAlert).
		Doc("delete AlertConfig").
		Param(restful.QueryParameter("alert_config_id", "delete alert config by id")).
		Writes(AlertConfigResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	tags = []string{"alert history api"}
	// get alert history
	ws.Route(ws.GET("/alerts/history").To(handler.RetrieveAlertHistory).
		Doc("retrieve alert history").
		// get alert history by alert_history_id
		Param(restful.QueryParameter("alert_history_id", "").Required(false)).
		// get alert history by alert_binding_id
		Param(restful.QueryParameter("alert_config_id", "").Required(false)).
		// get alert history by alert_rule_id
		Param(restful.QueryParameter("alert_rule_id", "").Required(false)).
		// get alert history by resource_id or resource_name
		Param(restful.QueryParameter("resource_id", "").Required(false)).
		//Param(restful.QueryParameter("resource_name", "").Required(false)).
		// get alert history by product
		Param(restful.QueryParameter("product_id", "").Required(false)).
		// return specific page of alert history
		Param(restful.QueryParameter("page", "").Required(false)).
		Param(restful.QueryParameter("limit", "").Required(false)).
		// get alert history by field fuzzy query
		Param(restful.QueryParameter("field", "").Required(false)).
		Param(restful.QueryParameter("fuzz", "").Required(false)).
		// get alert history which is in a specific time range
		Param(restful.QueryParameter("start_time", "").Required(false)).
		Param(restful.QueryParameter("end_time", "").Required(false)).
		Writes(AlertHistoryResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	tags = []string{"receiver api"}
	ws.Route(ws.GET("/receivers").To(handler.SearchReceiver).
		Doc("list receivers").
		Param(restful.QueryParameter("operator", "search")).
		Param(restful.QueryParameter("receiver_name", "").Required(false)).
		Param(restful.QueryParameter("email", "").Required(false)).
		Param(restful.QueryParameter("phone", "").Required(false)).
		Param(restful.QueryParameter("wechat", "").Required(false)).
		Writes([]Receiver{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	tags = []string{"resource api"}
	ws.Route(ws.GET("/resource").To(handler.GetResource).
		Doc("get resource spec").
		Param(restful.QueryParameter("cluster", "").Required(false)).
		Param(restful.QueryParameter("node", "").Required(false)).
		Param(restful.QueryParameter("workspace", "").Required(false)).
		Param(restful.QueryParameter("namespace", "").Required(false)).
		Param(restful.QueryParameter("workload", "").Required(false)).
		Param(restful.QueryParameter("resource_name", "")).
		Param(restful.QueryParameter("resource_type", "")).
		Writes(Resource{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	tags = []string{"silence apis"}

	// silence a alert_rule in specific period
	ws.Route(ws.POST("/silence").To(handler.CreateSilence).
		Doc("create silence").
		Reads(Silence{}).
		Writes(SilenceResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/silence").To(handler.RetrieveSilence).
		Doc("retrieve silence").
		Param(restful.QueryParameter("alert_config_id", "")).
		Param(restful.QueryParameter("alert_rule_id", "")).
		Param(restful.QueryParameter("resource_id", "")).
		Writes(SilenceResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/silence").To(handler.UpdateSilence).
		Doc("update silence").
		Reads(Silence{}).
		Writes(SilenceResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/silence").To(handler.DeleteSilence).
		Doc("delete silence").
		Param(restful.QueryParameter("alert_config_id", "")).
		Param(restful.QueryParameter("alert_rule_id", "")).
		Param(restful.QueryParameter("resource_id", "")).
		Writes(SilenceResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// enterprise crud
	tags = []string{"enterprise apis"}

	ws.Route(ws.POST("/enterprise").To(handler.CreateEnterprise).
		Doc("create enterprise").
		Reads(Enterprise{}).
		Writes(EnterpriseResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/enterprise").To(handler.RetrieveEnterprise).
		Doc("retrieve enterprise").
		Param(restful.QueryParameter("enterprise_id", "")).
		Param(restful.QueryParameter("enterprise_name", "")).
		Writes(EnterpriseResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/enterprise").To(handler.UpdateEnterprise).
		Doc("update enterprise").
		Reads(Enterprise{}).
		Writes(EnterpriseResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/enterprise").To(handler.DeleteEnterprise).
		Doc("delete enterprise").
		Param(restful.QueryParameter("enterprise_id", "")).
		Param(restful.QueryParameter("enterprise_name", "")).
		Writes(EnterpriseResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// product crud
	tags = []string{"product apis"}

	ws.Route(ws.POST("/product").To(handler.CreateProduct).
		Doc("create product").
		Reads(Product{}).
		Writes(ProductResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/product").To(handler.RetrieveProduct).
		Doc("retrieve product").
		Param(restful.QueryParameter("product_id", "")).
		Param(restful.QueryParameter("product_name", "")).
		Writes(ProductResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/product").To(handler.UpdateProduct).
		Doc("update product").
		Reads(Product{}).
		Writes(ProductResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/product").To(handler.DeleteProduct).
		Doc("delete product").
		Param(restful.QueryParameter("product_id", "")).
		Param(restful.QueryParameter("product_name", "")).
		Writes(ProductResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// resource_type crud
	tags = []string{"resource type apis"}

	ws.Route(ws.POST("/resource_type").To(handler.CreateResourceType).
		Doc("create resource_type").
		Reads(ResourceType{}).
		Writes(ResourceTypeResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/resource_type").To(handler.RetrieveResourceType).
		Doc("retrieve resource_type").
		Param(restful.QueryParameter("resource_type_id", "")).
		// or get resource type by resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type_name", "")).
		Writes(ResourceTypeResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/resource_type").To(handler.UpdateResourceType).
		Doc("update resource_type").
		Reads(ResourceType{}).
		Writes(ResourceTypeResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/resource_type").To(handler.DeleteResourceType).
		Doc("delete resource_type").
		Param(restful.QueryParameter("resource_type_id", "")).
		// or delete resource type by resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type_name", "")).
		Writes(ResourceTypeResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// alert_rule crud
	tags = []string{"alert rule apis"}

	ws.Route(ws.POST("/alert_rule").To(handler.CreateAlertRule).
		Doc("create alert_rule").
		Reads(AlertRule{}).
		Writes(AlertRuleResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/alert_rule").To(handler.RetrieveAlertRule).
		Doc("retrieve alert_rule").
		Param(restful.QueryParameter("alert_rule_id", "")).
		// or get alert rule by alert_rule_name + resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type_name", "")).
		Param(restful.QueryParameter("alert_rule_name", "")).
		// or get alert rule by resourc_family, example: list all cluster level alert rule
		Param(restful.QueryParameter("resource_family", "")).
		Writes([]AlertRuleResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alert_rule").To(handler.UpdateAlertRule).
		Doc("update alert_rule").
		Reads(AlertRule{}).
		Writes(AlertRuleResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alert_rule").To(handler.DeleteAlertRule).
		Doc("delete alert_rule").
		Param(restful.QueryParameter("alert_rule_id", "")).
		// or delete alert rule by alert_rule_name + resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type_name", "")).
		Param(restful.QueryParameter("alert_rule_name", "")).
		Writes(AlertRuleResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	return ws
}
