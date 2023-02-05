package dao

func InsertLike(uid int, gid int) (err error) {
	sqlStr := "insert into likes (uid,gid) values (?,?)"
	_, err = DB.Exec(sqlStr, uid, gid)
	return err
}
func DelLike(uid int, gid int) (err error) {
	sqlStr := "delete from likes where uid=? and gid=?"
	_, err = DB.Exec(sqlStr, uid, gid)
	return err
}
