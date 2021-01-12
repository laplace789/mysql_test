package sql

import "github.com/laplace789/mysql_test/config"

type mysql struct {
	cfg config.ServiceCfg
}

//Init will init mysql connection
func Init(c *config.ServiceCfg) {
}
