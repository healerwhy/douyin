
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	// data.DeletedTime = time.Unix(0,0)
	{{if .withCache}}{{.keys}}
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	if session != nil{
		return session.ExecCtx(ctx,query,{{.expressionValues}})
	}
	return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}
	query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	if session != nil{
		return session.ExecCtx(ctx,query,{{.expressionValues}})
	}
	return m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
}

func (m *default{{.upperStartCamelObject}}Model)  InsertOrUpdate(ctx context.Context, session sqlx.Session, field string, setStatus string, userId, objId, opt int64) (sql.Result, error) {
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%d,%d,?) ON DUPLICATE KEY UPDATE %s=?", m.table, field, userId, objId, setStatus)
		if session != nil {
			return session.ExecCtx(ctx, query, opt, opt)
		}
		return conn.ExecCtx(ctx, query, opt, opt)
	})
}