// Code generated by goctl. DO NOT EDIT!

package message

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	messagesFieldNames          = builder.RawFieldNames(&Messages{})
	messagesRows                = strings.Join(messagesFieldNames, ",")
	messagesRowsExpectAutoSet   = strings.Join(stringx.Remove(messagesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	messagesRowsWithPlaceHolder = strings.Join(stringx.Remove(messagesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	messagesModel interface {
		Insert(ctx context.Context, data *Messages) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Messages, error)
		Update(ctx context.Context, data *Messages) error
		Delete(ctx context.Context, id int64) error
		FindCode(ctx context.Context, id int64) (code string, phone string, timestamp int64, err error)
		InsertCode(ctx context.Context, phone string, code string) (sql.Result, error)
	}

	defaultMessagesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Messages struct {
		Id        int64          `db:"id"`
		Timestamp int64          `db:"timestamp"`
		Phone     sql.NullString `db:"phone"`
		Type      int64          `db:"type"`
		Content   sql.NullString `db:"content"`
		Code      sql.NullString `db:"code"`
	}
)

func newMessagesModel(conn sqlx.SqlConn) *defaultMessagesModel {
	return &defaultMessagesModel{
		conn:  conn,
		table: "`messages`",
	}
}

func (m *defaultMessagesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultMessagesModel) FindOne(ctx context.Context, id int64) (*Messages, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", messagesRows, m.table)
	var resp Messages
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessagesModel) Insert(ctx context.Context, data *Messages) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, messagesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Timestamp, data.Phone, data.Type, data.Content, data.Code)
	return ret, err
}

func (m *defaultMessagesModel) Update(ctx context.Context, data *Messages) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, messagesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Timestamp, data.Phone, data.Type, data.Content, data.Code, data.Id)
	return err
}

func (m *defaultMessagesModel) tableName() string {
	return m.table
}

func (m *defaultMessagesModel) FindCode(ctx context.Context, id int64) (code string, phone string, timestamp int64, err error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", messagesRows, m.table)
	var resp Messages
	err = m.conn.QueryRowCtx(ctx, &resp, query, id)

	if err != nil {
		return "", "", 0, err
	}

	return resp.Code.String, resp.Phone.String, resp.Timestamp, nil

}

func (m *defaultMessagesModel) InsertCode(ctx context.Context, phone string, code string) (sql.Result, error) {
	data := &Messages{
		Phone: sql.NullString{String: phone, Valid: true},
		Code: sql.NullString{
			String: code,
			Valid:  true,
		},
	}

	logx.Info(data)

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, messagesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, time.Now().Unix(), data.Phone, data.Type, data.Content, data.Code)
	return ret, err
}