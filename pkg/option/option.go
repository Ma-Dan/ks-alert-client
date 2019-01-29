package option

import "flag"

var (
	MysqlHost = flag.String("mysql", "127.0.0.1", "")
	MysqlPort = flag.String("mysql_port", "3306", "")
	Database  = flag.String("database", "alert_client", "")
	User      = flag.String("user", "root", "")
	Password  = flag.String("password", "password", "")
	//DispatcherServiceHost = flag.String("dispatcher_service", "alerting-dispatcher-server.kubesphere-monitoring-system.svc", "")
	//DispatcherServicePort = flag.String("dispatcher_port", "50000", "")

	DispatcherServiceHost = flag.String("dispatcher_service", "139.198.120.226", "")
	DispatcherServicePort = flag.String("dispatcher_port", "35000", "")

	//DispatcherServiceHost = flag.String("dispatcher_service", "172.31.140.133", "")
	//DispatcherServicePort = flag.String("dispatcher_port", "50000", "")
)

func init() {
	flag.Parse()
}
