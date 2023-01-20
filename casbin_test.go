package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

func Test_casbin(t *testing.T) {

	suite.Run(t, new(MyTestSuit))
	// 使用 MySQL 数据库初始化一个 Xorm 适配器
	a, err := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/blog", true)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	m, err := model.NewModelFromFile("./conf/rbac_model.conf")
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	added, err := e.AddPolicy("alice", "data1", "read")
	added, err = e.AddPolicy("domain1", "data1", "read")
	added, err = e.AddGroupingPolicy("domain1", "data1")
	if added {

	}
	sub := "alice" // 想要访问资源的用户
	obj := "data1" // 将要被访问的资源
	act := "read"  // 用户对资源实施的操作

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理错误
		fmt.Printf("%s", err)
	}

	if ok == true {
		// 允许 alice 读取 data1
		fmt.Println("allow")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("deny")
	}
}
