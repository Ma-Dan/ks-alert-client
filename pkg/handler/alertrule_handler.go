package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

type GetAlertRuleParams struct {
	UserID             string
	AlertConfigID      string
	AlertConfigName    string
	AlertRuleGroupID   string
	AlertRuleGroupName string
	ResourceType       string
	ResourceName       string
	ParentResourceType string
	ParentResourceName string
}

func HandlerAlertRule(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewAlertRuleHandlerClient(conn)

	method := request.Request.Method

	switch method {
	case http.MethodGet, http.MethodDelete:

		params := parseGetParams(request)
		ruleGroupIDs, err := DoGetAlertRuleGroupID(params)
		if err != nil {
			glog.Errorln(err.Error())
		}

		if method == http.MethodGet {
			var ruleGroupSpec *pb.AlertRuleGroupSpec
			if len(ruleGroupIDs) == 0 {
				if params.ResourceType != "" {
					ruleGroupSpec = &pb.AlertRuleGroupSpec{
						SystemRule:     true,
						ResourceTypeId: params.ResourceType,
					}
					// find build-in rule group
					rsp, err := cli.GetAlertRule(context.Background(), ruleGroupSpec)

					if err != nil {
						glog.Errorln(err)
						response.WriteHeaderAndEntity(http.StatusInternalServerError, &[]pb.AlertRuleGroupResponse{
							{
								Error: &pb.Error{Text: err.Error()},
							}})
						return
					} else {
						response.WriteHeaderAndEntity(http.StatusOK, &[]pb.AlertRuleGroupResponse{*rsp})
					}

				} else {
					// error invalid params

				}
			} else {

				var rsps []pb.AlertRuleGroupResponse
				for i := 0; i < len(ruleGroupIDs); i++ {
					ruleGroupSpec = &pb.AlertRuleGroupSpec{
						AlertRuleGroupId: ruleGroupIDs[i],
					}

					rsp, err := cli.GetAlertRule(context.Background(), ruleGroupSpec)

					if err != nil {
						glog.Errorln(err)
						rsp.Error = &pb.Error{Text: err.Error(), Code: pb.Error_ACCESS_DENIED}
					}
					rsps = append(rsps, *rsp)
				}

				response.WriteHeaderAndEntity(http.StatusOK, rsps)

			}
		} else {

			if len(ruleGroupIDs) != 1 {
				errStr := "can only delete one rule group at once"
				response.WriteHeaderAndEntity(http.StatusOK, &pb.AlertRuleGroupResponse{Error: &pb.Error{Text: errStr, Code: pb.Error_ACCESS_DENIED}})
				return

			} else {
				rsp, err := cli.DeleteAlertRule(context.Background(), &pb.AlertRuleGroupSpec{
					AlertRuleGroupId: ruleGroupIDs[0],
				})

				if err != nil {
					glog.Errorln(err)
					response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.AlertRuleGroupResponse{Error: &pb.Error{Text: err.Error(), Code: pb.Error_ACCESS_DENIED}})
					return
				} else {
					response.WriteHeaderAndEntity(http.StatusOK, rsp)
					return
				}

			}
		}

	case http.MethodPost, http.MethodPut:
		var ruleGroup = pb.AlertRuleGroup{}
		err := request.ReadEntity(&ruleGroup)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.AlertRuleGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if method == http.MethodPost {
			rsp, err := cli.CreateAlertRule(context.Background(), &ruleGroup)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.AlertRuleGroupResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		} else {

			rsp, err := cli.UpdateAlertRule(context.Background(), &ruleGroup)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.AlertRuleGroupResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		}
	}
}

func DoGetAlertRuleGroupID(p *GetAlertRuleParams) ([]string, error) {
	var ids []string

	if p.UserID != "" {
		// get alert rule group ids  by userID

	} else if p.AlertConfigID != "" {
		// get alert rule group id by alert config id

	} else if p.AlertConfigName != "" {
		// get alert config id by alert config name

	} else if p.AlertRuleGroupID != "" {
		ids = append(ids, p.AlertRuleGroupID)

	} else if p.ResourceType != "" && p.ResourceName != "" {
		// get alert group ids by resource type and resource name

	} else if p.ParentResourceType != "" && p.ParentResourceName != "" {

	} else {
		// error invalid params

	}

	return ids, nil
}

func parseGetParams(request *restful.Request) *GetAlertRuleParams {
	return &GetAlertRuleParams{
		UserID:             request.QueryParameter("user_id"),
		AlertConfigID:      request.QueryParameter("alert_config_id"),
		AlertConfigName:    request.QueryParameter("alert_config_name"),
		AlertRuleGroupID:   request.QueryParameter("alert_rule_group_id"),
		ResourceType:       request.QueryParameter("resource_type"),
		ResourceName:       request.QueryParameter("resource_name"),
		ParentResourceName: request.QueryParameter("parent_resource_type"),
		ParentResourceType: request.QueryParameter("parent_resource_name"),
	}
}
