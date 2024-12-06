package main

import (
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/siddontang/go-log/log"
)

/*
docker run -d \
--name mysql \
-e MYSQL_ROOT_PASSWORD=root \
-e MYSQL_LOG_BIN=ON mysql:latest
*/

/*
CREATE TABLE IF NOT EXISTS movie (

	uid INT not null auto_increment comment 'id',
	code varchar(50) default null comment '编号',
	kind INT default null comment '电影类型',
	name varchar(20) default null comment '电影名称',
	ranks INT default null comment '排名',
	remark varchar(200) default null comment '描述',
	primary key (uid)

) engine=innodb default charset=utf8mb4;
*/

type Movie struct {
	Uid    int    `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`             //id
	Code   string `gorm:"column:code" db:"code" json:"code" form:"code"`         //编号
	Kind   int    `gorm:"column:kind" db:"kind" json:"kind" form:"kind"`         //电影类型
	Name   string `gorm:"column:name" db:"name" json:"name" form:"name"`         //电影名称
	Ranks  int    `gorm:"column:ranks" db:"ranks" json:"ranks" form:"ranks"`     //排名
	Remark string `gorm:"column:remark" db:"remark" json:"remark" form:"remark"` //描述
}
type MyEventHandler struct {
	canal.DummyEventHandler
}

func NewMyEventHandler() *MyEventHandler {
	return &MyEventHandler{}
}
func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {

	log.Infof("table: %s rowevent: %s %v\n", e.Table.Name, e.Action, e.Rows)
	data := e.Rows[0]
	var movie Movie
	movie.Uid = data[0].(int)

	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}
