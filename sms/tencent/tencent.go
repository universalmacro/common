package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/universalmacro/common/sms/models"
)

func NewSmsSender(region, secretId, secretKey string) *SmsSender {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	cpf := profile.NewClientProfile()

	return &SmsSender{
		credential: credential,
		profile:    cpf,
	}
}

type Config struct {
	Region     *string
	AppId      string
	SignName   *string
	TemplateId string
}

type SmsSender struct {
	credential *common.Credential
	profile    *profile.ClientProfile
	Region     string
}

func (s *SmsSender) SendWithConfig(to models.PhoneNumber, config Config, vars []string) error {
	region := "ap-singapore"
	if config.Region != nil {
		region = *config.Region
	}
	client, _ := sms.NewClient(s.credential, region, s.profile)
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(config.AppId)
	if config.SignName == nil {
		request.SignName = common.StringPtr(*config.SignName)
	}
	request.TemplateParamSet = common.StringPtrs(vars)
	request.TemplateId = common.StringPtr(config.TemplateId)
	request.PhoneNumberSet = common.StringPtrs([]string{"+" + to.AreaCode + to.Number})
	_, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err
	}
	return nil
}
