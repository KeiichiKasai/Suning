package service

import (
	"main.go/dao"
	"main.go/model"
)

func CreateComment(gid int, username string, content string) (err error) {
	err = dao.InsertComment(gid, username, content)
	return err
}

func AddComment(gid int) (err error) {
	var g model.Goods
	g, err = dao.SelectGoodsComment(gid)
	if err != nil {
		return err
	}
	now := g.Comment + 1
	err = dao.UpdateCommentCount(gid, now)
	return err
}
func GetComments(gid int) (comments []model.Comment, err error) {
	comments, err = dao.SelectComment(gid)
	return comments, err
}
