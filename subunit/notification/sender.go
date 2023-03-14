package notification

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/backend-common/opurl"
)

type Sender interface {
	Dong([]string, string, string) error
	Email([]string, string, string) error
	Wechat([]string, string) error
	SMS([]string, string) error
	Phone([]string, string) error
}

type YunWeiLoader interface {
	LoadDong() (*model.Dong, error)
	LoadYunWei() (*model.Alert, error)
}

type ywAlert struct {
	client opurl.Client
	load   YunWeiLoader
}

func (yw *ywAlert) Dong(ids []string, title string, content string) error {
	dc, err := yw.load.LoadDong()
	if err != nil {
		return err
	}
	req := &dongRequest{
		UserIDs: strings.Join(ids, ","),
		Title:   title,
		Detail:  content,
	}
	header := http.Header{
		"Host":         []string{dc.Host},
		"Account":      []string{dc.Account},
		"Token":        []string{dc.Token},
		"Accept":       []string{"application/json"},
		"Content-Type": []string{"application/json; charset=utf-8"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reply := new(dongResponse)
	if err = yw.client.FetchJSON(ctx, http.MethodPost, dc.Addr, header, req, reply); err != nil {
		return err
	}

	return reply.Error()
}

func (yw *ywAlert) Email(ids []string, title string, content string) error {
	// TODO implement me
	panic("implement me")
}

func (yw *ywAlert) Wechat(ids []string, content string) error {
	// TODO implement me
	panic("implement me")
}

func (yw *ywAlert) SMS(ids []string, content string) error {
	// TODO implement me
	panic("implement me")
}

func (yw *ywAlert) Phone(ids []string, content string) error {
	// TODO implement me
	panic("implement me")
}
