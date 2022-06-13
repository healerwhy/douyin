// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"douyin/common/globalkey"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFollowListFieldNames          = builder.RawFieldNames(&UserFollowList{})
	userFollowListRows                = strings.Join(userFollowListFieldNames, ",")
	userFollowListRowsExpectAutoSet   = strings.Join(stringx.Remove(userFollowListFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userFollowListRowsWithPlaceHolder = strings.Join(stringx.Remove(userFollowListFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheDouyin2UserFollowListIdPrefix             = "cache:douyin2:userFollowList:id:"
	cacheDouyin2UserFollowListUserIdFollowIdPrefix = "cache:douyin2:userFollowList:userId:followId:"
)

type (
	userFollowListModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *UserFollowList) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserFollowList, error)
		FindOneByUserIdFollowId(ctx context.Context, userId int64, followId int64) (*UserFollowList, error)
		Update(ctx context.Context, session sqlx.Session, data *UserFollowList) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *UserFollowList) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserFollowListModel struct {
		sqlc.CachedConn
		table string
	}

	UserFollowList struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"`
		FollowId   int64     `db:"follow_id"`
		DelState   int64     `db:"del_state"`
		CreateTime time.Time `db:"create_time"`
	}
)

func newUserFollowListModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserFollowListModel {
	return &defaultUserFollowListModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_follow_list`",
	}
}

func (m *defaultUserFollowListModel) Insert(ctx context.Context, session sqlx.Session, data *UserFollowList) (sql.Result, error) {
	douyin2UserFollowListIdKey := fmt.Sprintf("%s%v", cacheDouyin2UserFollowListIdPrefix, data.Id)
	douyin2UserFollowListUserIdFollowIdKey := fmt.Sprintf("%s%v:%v", cacheDouyin2UserFollowListUserIdFollowIdPrefix, data.UserId, data.FollowId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userFollowListRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.UserId, data.FollowId, data.DelState)
		}
		return conn.ExecCtx(ctx, query, data.UserId, data.FollowId, data.DelState)
	}, douyin2UserFollowListIdKey, douyin2UserFollowListUserIdFollowIdKey)
}

func (m *defaultUserFollowListModel) FindOne(ctx context.Context, id int64) (*UserFollowList, error) {
	douyin2UserFollowListIdKey := fmt.Sprintf("%s%v", cacheDouyin2UserFollowListIdPrefix, id)
	var resp UserFollowList
	err := m.QueryRowCtx(ctx, &resp, douyin2UserFollowListIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userFollowListRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserFollowListModel) FindOneByUserIdFollowId(ctx context.Context, userId int64, followId int64) (*UserFollowList, error) {
	douyin2UserFollowListUserIdFollowIdKey := fmt.Sprintf("%s%v:%v", cacheDouyin2UserFollowListUserIdFollowIdPrefix, userId, followId)
	var resp UserFollowList
	err := m.QueryRowIndexCtx(ctx, &resp, douyin2UserFollowListUserIdFollowIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `follow_id` = ? and del_state = ? limit 1", userFollowListRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, followId, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserFollowListModel) Update(ctx context.Context, session sqlx.Session, data *UserFollowList) (sql.Result, error) {
	douyin2UserFollowListIdKey := fmt.Sprintf("%s%v", cacheDouyin2UserFollowListIdPrefix, data.Id)
	douyin2UserFollowListUserIdFollowIdKey := fmt.Sprintf("%s%v:%v", cacheDouyin2UserFollowListUserIdFollowIdPrefix, data.UserId, data.FollowId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userFollowListRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.UserId, data.FollowId, data.DelState, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.UserId, data.FollowId, data.DelState, data.Id)
	}, douyin2UserFollowListIdKey, douyin2UserFollowListUserIdFollowIdKey)
}

func (m *defaultUserFollowListModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *UserFollowList) error {

	return nil

}

func (m *defaultUserFollowListModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	douyin2UserFollowListIdKey := fmt.Sprintf("%s%v", cacheDouyin2UserFollowListIdPrefix, id)
	douyin2UserFollowListUserIdFollowIdKey := fmt.Sprintf("%s%v:%v", cacheDouyin2UserFollowListUserIdFollowIdPrefix, data.UserId, data.FollowId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, douyin2UserFollowListIdKey, douyin2UserFollowListUserIdFollowIdKey)
	return err
}

func (m *defaultUserFollowListModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDouyin2UserFollowListIdPrefix, primary)
}
func (m *defaultUserFollowListModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userFollowListRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultUserFollowListModel) tableName() string {
	return m.table
}
