Update(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result, error)
UpdateStatus(ctx context.Context,session sqlx.Session, key string, idx string,actionType,id int64)  (sql.Result,error)
