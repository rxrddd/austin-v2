// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"austin-v2/common/dal/model"
)

func newMsgRecord(db *gorm.DB, opts ...gen.DOOption) msgRecord {
	_msgRecord := msgRecord{}

	_msgRecord.msgRecordDo.UseDB(db, opts...)
	_msgRecord.msgRecordDo.UseModel(&model.MsgRecord{})

	tableName := _msgRecord.msgRecordDo.TableName()
	_msgRecord.ALL = field.NewAsterisk(tableName)
	_msgRecord.ID = field.NewInt64(tableName, "id")
	_msgRecord.MessageTemplateID = field.NewInt64(tableName, "message_template_id")
	_msgRecord.RequestID = field.NewString(tableName, "request_id")
	_msgRecord.Receiver = field.NewString(tableName, "receiver")
	_msgRecord.MsgID = field.NewString(tableName, "msg_id")
	_msgRecord.Channel = field.NewString(tableName, "channel")
	_msgRecord.Msg = field.NewString(tableName, "msg")
	_msgRecord.SendAt = field.NewString(tableName, "send_at")
	_msgRecord.CreateAt = field.NewTime(tableName, "create_at")
	_msgRecord.StartConsumeAt = field.NewString(tableName, "start_consume_at")
	_msgRecord.EndConsumeAt = field.NewString(tableName, "end_consume_at")
	_msgRecord.ConsumeSinceTime = field.NewString(tableName, "consume_since_time")
	_msgRecord.SendSinceTime = field.NewString(tableName, "send_since_time")
	_msgRecord.TaskInfo = field.NewString(tableName, "task_info")

	_msgRecord.fillFieldMap()

	return _msgRecord
}

type msgRecord struct {
	msgRecordDo

	ALL               field.Asterisk
	ID                field.Int64
	MessageTemplateID field.Int64  // 消息模板ID
	RequestID         field.String // 唯一请求 ID
	Receiver          field.String // 接收人
	MsgID             field.String // 公众号消息id
	Channel           field.String // 渠道
	Msg               field.String // 推送结果信息
	SendAt            field.String // 消息http 发送时间
	CreateAt          field.Time
	StartConsumeAt    field.String // 开始消费时间
	EndConsumeAt      field.String // 结束消费时间
	ConsumeSinceTime  field.String // 消费间距时间
	SendSinceTime     field.String // http->mq消费结束间距时间
	TaskInfo          field.String

	fieldMap map[string]field.Expr
}

func (m msgRecord) Table(newTableName string) *msgRecord {
	m.msgRecordDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m msgRecord) As(alias string) *msgRecord {
	m.msgRecordDo.DO = *(m.msgRecordDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *msgRecord) updateTableName(table string) *msgRecord {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.MessageTemplateID = field.NewInt64(table, "message_template_id")
	m.RequestID = field.NewString(table, "request_id")
	m.Receiver = field.NewString(table, "receiver")
	m.MsgID = field.NewString(table, "msg_id")
	m.Channel = field.NewString(table, "channel")
	m.Msg = field.NewString(table, "msg")
	m.SendAt = field.NewString(table, "send_at")
	m.CreateAt = field.NewTime(table, "create_at")
	m.StartConsumeAt = field.NewString(table, "start_consume_at")
	m.EndConsumeAt = field.NewString(table, "end_consume_at")
	m.ConsumeSinceTime = field.NewString(table, "consume_since_time")
	m.SendSinceTime = field.NewString(table, "send_since_time")
	m.TaskInfo = field.NewString(table, "task_info")

	m.fillFieldMap()

	return m
}

func (m *msgRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *msgRecord) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 14)
	m.fieldMap["id"] = m.ID
	m.fieldMap["message_template_id"] = m.MessageTemplateID
	m.fieldMap["request_id"] = m.RequestID
	m.fieldMap["receiver"] = m.Receiver
	m.fieldMap["msg_id"] = m.MsgID
	m.fieldMap["channel"] = m.Channel
	m.fieldMap["msg"] = m.Msg
	m.fieldMap["send_at"] = m.SendAt
	m.fieldMap["create_at"] = m.CreateAt
	m.fieldMap["start_consume_at"] = m.StartConsumeAt
	m.fieldMap["end_consume_at"] = m.EndConsumeAt
	m.fieldMap["consume_since_time"] = m.ConsumeSinceTime
	m.fieldMap["send_since_time"] = m.SendSinceTime
	m.fieldMap["task_info"] = m.TaskInfo
}

func (m msgRecord) clone(db *gorm.DB) msgRecord {
	m.msgRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m msgRecord) replaceDB(db *gorm.DB) msgRecord {
	m.msgRecordDo.ReplaceDB(db)
	return m
}

type msgRecordDo struct{ gen.DO }

func (m msgRecordDo) Debug() *msgRecordDo {
	return m.withDO(m.DO.Debug())
}

func (m msgRecordDo) WithContext(ctx context.Context) *msgRecordDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m msgRecordDo) ReadDB() *msgRecordDo {
	return m.Clauses(dbresolver.Read)
}

func (m msgRecordDo) WriteDB() *msgRecordDo {
	return m.Clauses(dbresolver.Write)
}

func (m msgRecordDo) Session(config *gorm.Session) *msgRecordDo {
	return m.withDO(m.DO.Session(config))
}

func (m msgRecordDo) Clauses(conds ...clause.Expression) *msgRecordDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m msgRecordDo) Returning(value interface{}, columns ...string) *msgRecordDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m msgRecordDo) Not(conds ...gen.Condition) *msgRecordDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m msgRecordDo) Or(conds ...gen.Condition) *msgRecordDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m msgRecordDo) Select(conds ...field.Expr) *msgRecordDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m msgRecordDo) Where(conds ...gen.Condition) *msgRecordDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m msgRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *msgRecordDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m msgRecordDo) Order(conds ...field.Expr) *msgRecordDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m msgRecordDo) Distinct(cols ...field.Expr) *msgRecordDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m msgRecordDo) Omit(cols ...field.Expr) *msgRecordDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m msgRecordDo) Join(table schema.Tabler, on ...field.Expr) *msgRecordDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m msgRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) *msgRecordDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m msgRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) *msgRecordDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m msgRecordDo) Group(cols ...field.Expr) *msgRecordDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m msgRecordDo) Having(conds ...gen.Condition) *msgRecordDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m msgRecordDo) Limit(limit int) *msgRecordDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m msgRecordDo) Offset(offset int) *msgRecordDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m msgRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *msgRecordDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m msgRecordDo) Unscoped() *msgRecordDo {
	return m.withDO(m.DO.Unscoped())
}

func (m msgRecordDo) Create(values ...*model.MsgRecord) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m msgRecordDo) CreateInBatches(values []*model.MsgRecord, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m msgRecordDo) Save(values ...*model.MsgRecord) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m msgRecordDo) First() (*model.MsgRecord, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgRecord), nil
	}
}

func (m msgRecordDo) Take() (*model.MsgRecord, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgRecord), nil
	}
}

func (m msgRecordDo) Last() (*model.MsgRecord, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgRecord), nil
	}
}

func (m msgRecordDo) Find() ([]*model.MsgRecord, error) {
	result, err := m.DO.Find()
	return result.([]*model.MsgRecord), err
}

func (m msgRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MsgRecord, err error) {
	buf := make([]*model.MsgRecord, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m msgRecordDo) FindInBatches(result *[]*model.MsgRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m msgRecordDo) Attrs(attrs ...field.AssignExpr) *msgRecordDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m msgRecordDo) Assign(attrs ...field.AssignExpr) *msgRecordDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m msgRecordDo) Joins(fields ...field.RelationField) *msgRecordDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m msgRecordDo) Preload(fields ...field.RelationField) *msgRecordDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m msgRecordDo) FirstOrInit() (*model.MsgRecord, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgRecord), nil
	}
}

func (m msgRecordDo) FirstOrCreate() (*model.MsgRecord, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgRecord), nil
	}
}

func (m msgRecordDo) FindByPage(offset int, limit int) (result []*model.MsgRecord, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m msgRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m msgRecordDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m msgRecordDo) Delete(models ...*model.MsgRecord) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *msgRecordDo) withDO(do gen.Dao) *msgRecordDo {
	m.DO = *do.(*gen.DO)
	return m
}
