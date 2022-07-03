
func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}})  (sql.Result,error) {
	{{if .withCache}}{{.keys}}
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}})
	}
	return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}})
	}
	return m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
}

func (m *default{{.upperStartCamelObject}}Model) UpdateStatus(ctx context.Context,session sqlx.Session, key string, idx string,actionType,id int64)  (sql.Result,error) {

    	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
    		query := fmt.Sprintf("UPDATE %s SET %s=%s+? WHERE %s=?", m.table, key, key, idx)
    		if session != nil {
    			return session.ExecCtx(ctx, query, actionType, id)
    		}
    		return conn.ExecCtx(ctx, query, actionType, id)
    	})
}

