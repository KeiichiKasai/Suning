package service

import "main.go/dao"

func AddLike(uid int, gid int) (err error) {
	err = dao.InsertLike(uid, gid)
	return err
}
func MinusLike(uid int, gid int) (err error) {
	err = dao.DelLike(uid, gid)
	return err
}
