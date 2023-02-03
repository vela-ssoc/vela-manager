package sendto

import "github.com/vela-ssoc/manager/libkit/httpclient"

type phoneDev struct {
	configure AlertConfigurer
	client    httpclient.Client
}

func (pd *phoneDev) SendPhone(nos []string, content string) error {
	// TODO implement me
	panic("implement me")
}
