package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/itachilee/ginblog/global"
	"log"
)

var (
	e *casbin.Enforcer
)

func InitEnforce() *casbin.Enforcer {
	var err error
	//a, err := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/blog", true)
	a, err := gormadapter.NewAdapterByDB(global.DBEngine)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	m, err := model.NewModelFromFile("./conf/rbac_model.conf")
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err = casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	return e
}
