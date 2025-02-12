package scoring

// IncrementBackLinksCount Increment the backlinks count by 1
//
// If url is not exist into table then it will be inserted automatically
func (t *Cassandra) IncrementBackLinksCount(url string) (err error) {
	err = t.session.Query(`
		UPDATE backlinks 
		SET backlink_count = backlink_count + 1 
		WHERE url = ?;
	`, url).Exec()
	return
}
