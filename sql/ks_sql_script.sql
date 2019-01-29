-- enterprise
INSERT INTO alert.enterprises (enterprise_id, enterprise_name, contacts, email, phone, address, home_page, description, created_at, updated_at) VALUES ('6y19xy9pwm24oo', 'qingcloud', 'Richard', 'yunify@yunify.com', '400-8576-886', '北京优帆科技有限公司', 'https://www.qingcloud.com/', '云计算公司', '2019-01-13 21:50:03', '2019-01-13 21:50:03');

-- product
INSERT INTO alert.products (product_id, product_name, enterprise_id, phone, contacts, email, monitor_center_host, monitor_center_port, home_page, address, description, created_at, updated_at, webhook, webhook_enable) VALUES ('m57o28v7101rwz', 'KubeSphere', '6y19xy9pwm24oo', '400-8576-886', 'Ray', '', 'localhost', 8080, 'https://www.kubesphere.io/', '', '', '2019-01-13 21:55:26', '2019-01-13 21:55:26', 'TODO', 1);

-- severities
INSERT INTO alert.severities (severity_id, product_id, severity_en, severity_ch, created_at, updated_at) VALUES ('0mr4yrrv17pn8v', 'm57o28v7101rwz', 'Warn', '警告', '2019-01-13 22:08:09', '2019-01-13 22:08:09');
INSERT INTO alert.severities (severity_id, product_id, severity_en, severity_ch, created_at, updated_at) VALUES ('75y7ozzw1vq9kk', 'm57o28v7101rwz', 'Minor', '危险', '2019-01-13 22:08:09', '2019-01-13 22:08:09');
INSERT INTO alert.severities (severity_id, product_id, severity_en, severity_ch, created_at, updated_at) VALUES ('78n9p5rv17pn8v', 'm57o28v7101rwz', 'Major', '较危险', '2019-01-13 22:08:09', '2019-01-13 22:08:09');
INSERT INTO alert.severities (severity_id, product_id, severity_en, severity_ch, created_at, updated_at) VALUES ('wq68nl63l4yzj4', 'm57o28v7101rwz', 'Critical', '严重', '2019-01-13 22:08:09', '2019-01-13 22:08:09');

-- resource_type
INSERT INTO alert.resource_types (resource_type_id, product_id, resource_type_name, resource_uri_tmpls, description, enable, monitor_center_host, monitor_center_port, created_at, updated_at) VALUES ('k744yo5vol1zk5', 'm57o28v7101rwz', 'workload', '{"resource_uri_tmpl":[{"uri_tmpl":"/api/v1alpha1/monitoring/namespaces/{ns}/{wl_kind}","path_params":{"ns":"","wl_kind":""}}]}', '', 1, 'http://ks-apiserver.kubesphere-monitoring-system.svc', 80, '2019-01-14 17:55:34', '2019-01-13 22:44:15');
INSERT INTO alert.resource_types (resource_type_id, product_id, resource_type_name, resource_uri_tmpls, description, enable, monitor_center_host, monitor_center_port, created_at, updated_at) VALUES ('l07941wpjzvyll', 'm57o28v7101rwz', 'node', '{"resource_uri_tmpl":[{"uri_tmpl":"/api/v1alpha1/monitoring/nodes"}]}', '', 1, 'http://ks-apiserver.kubesphere-monitoring-system.svc', 80, '2019-01-14 17:55:34', '2019-01-13 22:44:15');
INSERT INTO alert.resource_types (resource_type_id, product_id, resource_type_name, resource_uri_tmpls, description, enable, monitor_center_host, monitor_center_port, created_at, updated_at) VALUES ('lvzkj25vol1zk5', 'm57o28v7101rwz', 'cluster', '{"resource_uri_tmpl":[{"uri_tmpl":"/api/v1alpha1/monitoring/clusters"}]}', '', 1, 'http://ks-apiserver.kubesphere-monitoring-system.svc', 80, '2019-01-14 17:55:33', '2019-01-13 22:44:07');
INSERT INTO alert.resource_types (resource_type_id, product_id, resource_type_name, resource_uri_tmpls, description, enable, monitor_center_host, monitor_center_port, created_at, updated_at) VALUES ('lz6o6vjqxlroww', 'm57o28v7101rwz', 'workspace', '{"resource_uri_tmpl":[{"uri_tmpl":"/api/v1alpha1/monitoring/workspaces"}]}', '', 1, 'http://ks-apiserver.kubesphere-monitoring-system.svc', 80, '2019-01-14 17:55:34', '2019-01-13 22:44:15');
INSERT INTO alert.resource_types (resource_type_id, product_id, resource_type_name, resource_uri_tmpls, description, enable, monitor_center_host, monitor_center_port, created_at, updated_at) VALUES ('o1jxqpk3vp9zo9', 'm57o28v7101rwz', 'container', '{"resource_uri_tmpl":[{"uri_tmpl":"/api/v1alpha1/monitoring/namespaces/{ns_name}/pods/{pod_name}/containers","path_params":{"ns_name":"","pod_name":""}},{"uri_tmpl":"/api/v1alpha1/monitoring/nodes/{node_id}/pods/{pod_name}/containers","path_params":{"node_id":"","pod_name":""}}]}', '', 1, 'http://ks-apiserver.kubesphere-monitoring-system.svc', 80, '2019-01-14 17:55:34', '2019-01-13 22:44:15');
INSERT INTO alert.resource_types (resource_type_id, product_id, resource_type_name, resource_uri_tmpls, description, enable, monitor_center_host, monitor_center_port, created_at, updated_at) VALUES ('q0lw3n1pn04yjp', 'm57o28v7101rwz', 'namespace', '{"resource_uri_tmpl":[{"uri_tmpl":"/api/v1alpha1/monitoring/workspaces/{ws_name}/namespaces","path_params":{"ws_name":""}},{"uri_tmpl":"/api/v1alpha1/monitoring/namespaces"}]}', '', 1, 'http://ks-apiserver.kubesphere-monitoring-system.svc', 80, '2019-01-14 17:55:33', '2019-01-13 22:44:15');
INSERT INTO alert.resource_types (resource_type_id, product_id, resource_type_name, resource_uri_tmpls, description, enable, monitor_center_host, monitor_center_port, created_at, updated_at) VALUES ('z3485jwpjzvyll', 'm57o28v7101rwz', 'pod', '{"resource_uri_tmpl":[{"uri_tmpl":"/api/v1alpha1/monitoring/namespaces/{ns_name}/pods","path_params":{"ns_name":""}},{"uri_tmpl":"/api/v1alpha1/monitoring/nodes/{node_id}/pods","path_params":{"node_id":""}}]}', '', 1, 'http://ks-apiserver.kubesphere-monitoring-system.svc', 80, '2019-01-14 17:55:33', '2019-01-13 22:44:15');

-- rule_group
INSERT INTO alert.alert_rule_groups (alert_rule_group_id, alert_rule_group_name, description, system_rule, resource_type_id, created_at, updated_at) VALUES ('4m636j7o48z3wz', 'ks_cluster_rules', '', 1, 'lvzkj25vol1zk5', '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rule_groups (alert_rule_group_id, alert_rule_group_name, description, system_rule, resource_type_id, created_at, updated_at) VALUES ('57jk07xk54138v', 'ks_node_rules', '', 1, 'l07941wpjzvyll', '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rule_groups (alert_rule_group_id, alert_rule_group_name, description, system_rule, resource_type_id, created_at, updated_at) VALUES ('99w68x9k9ol0j4', 'ks_workload_rules', '', 1, 'k744yo5vol1zk5', '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rule_groups (alert_rule_group_id, alert_rule_group_name, description, system_rule, resource_type_id, created_at, updated_at) VALUES ('v2lmxqjy7p15kk', 'ks_namespace_rules', '', 1, 'q0lw3n1pn04yjp', '2019-01-29 14:47:17', '2019-01-29 14:47:17');

-- rules
-- cluster
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('50w3jzn36w19oo', 'cluster_cpu_utilisation', '4m636j7o48z3wz', 'cluster_cpu_utilisation', '>', 1, 80, '%', 3, 3, 0, 1, '1', 60, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('773x3vjy7p15kk', 'cluster_net_utilisation', '4m636j7o48z3wz', 'cluster_net_utilisation', '>', 1, 200, 'bps', 3, 3, 0, 1, '1', 60, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('jy002j5o9vp8n8', 'cluster_namespace_count', '4m636j7o48z3wz', 'cluster_namespace_count', '>', 1, 25, '', 3, 3, 0, 1, '1', 60, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('m8zqxm7o48z3wz', 'cluster_memory_utilisation', '4m636j7o48z3wz', 'cluster_memory_utilisation', '>', 1, 80, '%', 3, 3, 0, 1, '1', 60, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('xk17xm7o48z3wz', 'cluster_node_offline', '4m636j7o48z3wz', 'cluster_node_offline', '>', 1, 1, '', 3, 3, 0, 1, '1', 60, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');

-- node
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('6zvy42n36w19oo', 'node_memory_utilisation', '57jk07xk54138v', 'node_memory_utilisation', '>', 1, 85, '%', 3, 3, 0, 1, '1', 100, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('jy85z75o9vp8n8', 'node_cpu_utilisation', '57jk07xk54138v', 'node_cpu_utilisation', '>', 1, 60, '%', 4, 3, 0, 1, '0', 120, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('w6n4zy5k9ol0j4', 'node_disk_inode_utilisation', '57jk07xk54138v', 'node_disk_inode_utilisation', '>', 1, 90, '%', 3, 3, 0, 1, '1', 60, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');

-- namespace
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('23vn4l9o9vp8n8', 'namespace_memory_usage_wo_cache', 'v2lmxqjy7p15kk', 'namespace_memory_usage_wo_cache', '>', 1, 500, 'm', 3, 3, 0, 1, '1', 60, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('509qjy5k54138v', 'namespace_cpu_usage', 'v2lmxqjy7p15kk', 'namespace_cpu_usage', '>', 1, 120, 'm', 4, 3, 0, 1, '0', 360, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('7411jy5k54138v', 'namespace_pod_count', 'v2lmxqjy7p15kk', 'namespace_pod_count', '>', 1, 70, '', 3, 3, 0, 1, '1', 100, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');

-- workload
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('4r7xlx9k9ol0j4', 'workload_pod_memory_usage_wo_cache', '99w68x9k9ol0j4', 'workload_pod_memory_usage_wo_cache', '>', 1, 600, 'm', 3, 3, 0, 1, '1', 100, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
INSERT INTO alert.alert_rules (alert_rule_id, alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES ('5x2m7q636w19oo', 'workload_pod_cpu_usage', '99w68x9k9ol0j4', 'workload_pod_cpu_usage', '>', 1, 600, 'm', 4, 3, 0, 1, '0', 120, 10, '2019-01-29 14:47:17', '2019-01-29 14:47:17');
