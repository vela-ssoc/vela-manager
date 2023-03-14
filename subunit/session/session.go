package session

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

// DBSess 数据库存放 session 管理器
func DBSess(db *gorm.DB, interval time.Duration) ship.Session {
	if interval <= 0 {
		interval = time.Hour
	} else if interval < time.Minute {
		interval = time.Minute
	}

	return &sessDB{
		db:       db,
		interval: interval,
	}
}

// sessDB 数据库存放 session 管理器
type sessDB struct {
	db       *gorm.DB      // 数据库连接
	interval time.Duration // session 有效活动间隔
}

// GetSession 根据 token 获取 session 信息
func (ssd *sessDB) GetSession(token string) (any, error) {
	info := new(model.Userinfo)
	if err := ssd.db.Where("token = ?", token).First(info).Error; err != nil {
		return nil, ship.ErrSessionNotExist
	}
	// 检查 session 是否有效
	now := time.Now()
	expiredAt := info.LastedAt.Add(ssd.interval)
	if expiredAt.Before(now) {
		return nil, ship.ErrSessionNotExist
	}

	// 如果用户每次获取 session 每次就续约，虽然 session 的精准度会更高，
	// 但是也会增加数据库的操作次数，且用户对 session 续期时间不是很敏感，
	// 所以此处会隔一段时间再对 session 实行续约。
	var renew bool
	diff := now.Sub(info.LastedAt)
	if ssd.interval <= 10*time.Minute {
		renew = diff >= 30*time.Second
	} else {
		renew = diff >= 5*time.Minute
	}
	if renew {
		ssd.db.Model(info).
			Where("token = ?", info.Token).
			UpdateColumn("lasted_at", now)
	}

	return info, nil
}

// SetSession 存放 session
func (ssd *sessDB) SetSession(token string, val any) error {
	info, ok := val.(*model.Userinfo)
	if !ok {
		return ship.ErrInvalidSession
	}
	if token == "" || info.Token != token {
		return ship.ErrInvalidSession
	}
	return ssd.db.Save(info).Error
}

// DelSession 根据 token 删除 session
func (ssd *sessDB) DelSession(token string) error {
	tbl := new(model.Userinfo)
	return ssd.db.
		Delete(tbl, "token = ?", token).
		Error
}

// Issued 签发 session 信息
func Issued(u *model.User) *model.Userinfo {
	now := time.Now()
	id := strconv.FormatInt(u.ID, 36)
	nano := strconv.FormatInt(now.UnixNano(), 36)
	buf := make([]byte, 32)
	_, _ = rand.Read(buf)
	nonce := hex.EncodeToString(buf)
	token := id + "." + nano + "." + nonce // id 时间 随机字符串组成，英文 . 分割

	return &model.Userinfo{
		ID:       u.ID,
		Username: u.Username,
		Nickname: u.Nickname,
		Token:    token,
		IssuedAt: now,
	}
}
