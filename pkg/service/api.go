package service

import (
	"github.com/carmanzhang/ks-alert-client/pkg/handler"
	. "github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

type AlertAPI struct{}

func (u AlertAPI) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		//Path("/apis/alerting.kubesphere.io/v1").
		Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"alert apis"}

	ws.Route(ws.POST("/alert").To(handler.HandlerAlertConfig).
		Doc("create AlertConfig").
		Reads(AlertConfig{}).
		Writes(AlertConfigResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// get batch of alert configs by resource_name
	ws.Route(ws.GET("/alert").To(handler.HandlerAlertConfig).
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
		// get alert config by alert_name + super_resource_type + super_resource_name,
		// alert_name is unique in any super_resource_name
		Param(restful.QueryParameter("alert_name", "get alert config by id").Required(false)).
		Writes([]AlertConfigResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alert").To(handler.HandlerAlertConfig).
		Doc("update AlertConfig").
		Reads(AlertConfig{}).
		Writes(AlertConfigResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alert").To(handler.HandlerAlertConfig).
		Doc("delete AlertConfig").
		Param(restful.QueryParameter("alert_config_id", "delete alert config by id")).
		Writes(AlertConfigResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	tags = []string{"alert history api"}
	// get alert history
	ws.Route(ws.GET("/alert/history").To(handler.RetrieveAlertHistory).
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

	//****************************************************************************************************
	tags = []string{"silence apis"}

	// silence a alert_rule in specific period
	ws.Route(ws.POST("/silence").To(handler.HandlerSilence).
		Doc("create silence").
		Reads(Silence{}).
		Writes(SilenceResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/silence").To(handler.HandlerSilence).
		Doc("retrieve silence").
		//Param(restful.QueryParameter("alert_config_id", "")).
		Param(restful.QueryParameter("alert_rule_id", "")).
		Param(restful.QueryParameter("resource_id", "")).
		Writes(SilenceResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/silence").To(handler.HandlerSilence).
		Doc("update silence").
		Reads(Silence{}).
		Writes(SilenceResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/silence").To(handler.HandlerSilence).
		Doc("delete silence").
		//Param(restful.QueryParameter("alert_config_id", "")).
		Param(restful.QueryParameter("alert_rule_id", "")).
		Param(restful.QueryParameter("resource_id", "")).
		Writes(SilenceResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	tags = []string{"severity apis"}

	// silence a alert_rule in specific period
	ws.Route(ws.POST("/severity").To(handler.HandlerSeverity).
		Doc("create silence").
		Reads(Severity{}).
		Writes(SeverityResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/severity").To(handler.HandlerSeverity).
		Doc("retrieve silence").
		Param(restful.QueryParameter("severity_id", "")).
		Param(restful.QueryParameter("product_id", "")).
		Writes(SeveritiesResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/severity").To(handler.HandlerSeverity).
		Doc("update silence").
		Reads(Severity{}).
		Writes(SeverityResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/severity").To(handler.HandlerSeverity).
		Doc("delete silence").
		Param(restful.QueryParameter("severity_id", "")).
		Param(restful.QueryParameter("product_id", "")).
		Writes(SeverityResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	tags = []string{"suggestion apis"}

	ws.Route(ws.GET("/suggestion").To(handler.HandlerSuggestion).
		Doc("retrieve silence").
		Param(restful.QueryParameter("alert_config_id", "")).
		Param(restful.QueryParameter("alert_rule_id", "")).
		Param(restful.QueryParameter("resource_id", "")).
		Writes(SuggestionResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/suggestion").To(handler.HandlerSuggestion).
		Doc("update silence").
		Reads(Suggestion{}).
		Writes(SuggestionResponse{}).
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

	ws.Route(ws.POST("/product").To(handler.HandlerProduct).
		Doc("create product").
		Reads(Product{}).
		Writes(ProductResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/product").To(handler.HandlerProduct).
		Doc("retrieve product").
		Param(restful.QueryParameter("product_id", "")).
		Param(restful.QueryParameter("product_name", "")).
		Writes(ProductResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/product").To(handler.HandlerProduct).
		Doc("update product").
		Reads(Product{}).
		Writes(ProductResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/product").To(handler.HandlerProduct).
		Doc("delete product").
		Param(restful.QueryParameter("product_id", "")).
		Param(restful.QueryParameter("product_name", "")).
		Writes(ProductResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// resource_type crud
	tags = []string{"resource type apis"}

	ws.Route(ws.POST("/resource_type").To(handler.HandlerResourceType).
		Doc("create resource_type").
		Reads(ResourceType{}).
		Writes(ResourceTypeResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/resource_type").To(handler.HandlerResourceType).
		Doc("retrieve resource_type").
		Param(restful.QueryParameter("resource_type_id", "")).
		// or get resource type by resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type_name", "")).
		Writes(ResourceTypeResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/resource_type").To(handler.HandlerResourceType).
		Doc("update resource_type").
		Reads(ResourceType{}).
		Writes(ResourceTypeResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/resource_type").To(handler.HandlerResourceType).
		Doc("delete resource_type").
		Param(restful.QueryParameter("resource_type_id", "")).
		// or delete resource type by resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type_name", "")).
		Writes(ResourceTypeResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// alert_rule crud
	tags = []string{"alert rule apis"}

	ws.Route(ws.POST("/alert_rule").To(handler.HandlerAlertRule).
		Doc("create alert_rule").
		Reads(AlertRuleGroup{}).
		Writes(AlertRuleGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/alert_rule").To(handler.HandlerAlertRule).
		Doc("retrieve alert_rule").
		Param(restful.QueryParameter("user_id", "")).
		Param(restful.QueryParameter("alert_rule_group_id", "")).
		Param(restful.QueryParameter("alert_rule_group_name", "")).
		Param(restful.QueryParameter("alert_rule_id", "")).
		Param(restful.QueryParameter("alert_rule_name", "")).
		// or get alert rule by alert_rule_name + resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type", "")).
		Param(restful.QueryParameter("resource_name", "")).
		// or get alert rule by resourc_family, example: list all cluster level alert rule
		Param(restful.QueryParameter("parent_resource_type", "")).
		Param(restful.QueryParameter("parent_resource_name", "")).
		Writes([]AlertRuleGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alert_rule").To(handler.HandlerAlertRule).
		Doc("update alert_rule").
		Reads(AlertRuleGroup{}).
		Writes(AlertRuleGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alert_rule").To(handler.HandlerAlertRule).
		Doc("delete alert_rule").
		Param(restful.QueryParameter("alert_rule_group_id", "")).
		Param(restful.QueryParameter("alert_rule_id", "")).
		// or delete alert rule by alert_rule_name + resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type_name", "")).
		Param(restful.QueryParameter("alert_rule_name", "")).
		Param(restful.QueryParameter("resource_family", "")).
		Writes(AlertRuleGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	tags = []string{"receiver api"}
	ws.Route(ws.POST("/receiver").To(handler.HandlerReceiver).
		Doc("create receiver group").
		Operation("create receivers").
		Reads(ReceiverGroup{}).
		Writes(ReceiverGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/receiver").To(handler.HandlerReceiver).
		Doc("retrieve receiver").
		Operation("get receivers").
		Param(restful.QueryParameter("receiver_group_id", "")).
		Param(restful.QueryParameter("operator", "search")).
		Param(restful.QueryParameter("receiver_name", "").Required(false)).
		Param(restful.QueryParameter("email", "").Required(false)).
		Param(restful.QueryParameter("phone", "").Required(false)).
		Param(restful.QueryParameter("wechat", "").Required(false)).
		Writes(ReceiverGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/receiver").To(handler.HandlerReceiver).
		Doc("update receiver").
		Operation("update receivers").
		Reads(ReceiverGroup{}).
		Writes(ReceiverGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/receiver").To(handler.HandlerReceiver).
		Doc("delete receiver").
		Operation("delete receivers").
		Param(restful.QueryParameter("receiver_group_id", "")).
		Writes(ReceiverGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	tags = []string{"resource api"}
	ws.Route(ws.POST("/resource").To(handler.HandlerResource).
		Doc("create resource group").
		Reads(ResourceGroup{}).
		Writes(ResourceGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/resource").To(handler.HandlerResource).
		Doc("retrieve resource").
		Param(restful.QueryParameter("cluster", "").Required(false)).
		Param(restful.QueryParameter("node", "").Required(false)).
		Param(restful.QueryParameter("workspace", "").Required(false)).
		Param(restful.QueryParameter("namespace", "").Required(false)).
		Param(restful.QueryParameter("workload", "").Required(false)).
		Param(restful.QueryParameter("resource_name", "")).
		Param(restful.QueryParameter("resource_type", "")).
		// add this parameter temporary
		Param(restful.QueryParameter("resource_group_id", "")).
		Writes(ResourceGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/resource").To(handler.HandlerResource).
		Doc("update resource").
		Reads(ResourceGroup{}).
		Writes(ResourceGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/resource").To(handler.HandlerResource).
		Doc("delete resource").
		Param(restful.QueryParameter("resource_id", "")).
		// or delete alert rule by resource_name + resource_type_name + enterprise + product
		Param(restful.QueryParameter("resource_type_name", "")).
		Param(restful.QueryParameter("resource_name", "")).
		// add this parameter temporary
		Param(restful.QueryParameter("resource_group_id", "")).
		Writes(ResourceGroupResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	return ws
}
