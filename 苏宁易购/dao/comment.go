package dao

import (
	"fmt"
	"main.go/model"
)

func InsertComment(gid int, username string, content string) (err error) {
	sqlStr := "insert into comment (gid,username,content) values (?,?,?)"
	_, err = DB.Exec(sqlStr, gid, username, content)
	if err != nil {
		fmt.Printf("插入评论失败，err:%v", err)
		return err
	}
	return err
}

func SelectGoodsComment(gid int) (g model.Goods, err error) {
	sqlStr := "select comment from goods where gid=?"
	err = DB.Get(&g, sqlStr, gid)
	if err != nil {
		fmt.Printf("查询商品失败，err:%v", err)
		return g, err
	}
	return g, err
}

func UpdateCommentCount(gid int, num int) (err error) {
	sqlStr := "update goods set comment=? where gid=?"
	_, err = DB.Exec(sqlStr, num, gid)
	if err != nil {
		fmt.Printf("修改评论数失败，err:%v", err)
		return err
	}
	return err
}

func SelectComment(gid int) (comments []model.Comment, err error) {
	sqlStr := "select cid,gid,username,ctime,content from comment where gid=?"
	err = DB.Select(&comments, sqlStr, gid)
	if err != nil {
		fmt.Printf("查询评论失败,err:%v", err)
		return nil, err
	}
	return comments, err
}
