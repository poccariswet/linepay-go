package linepay

type LinePay struct {
	ChannelID     string
	ChannelSecret string
	Hostname      string
	IsSandBox     bool
}

const (
	ApiVersion         = "v2"
	SandboxApiHostname = "sandbox-api-pay.line.me"
)

func New(channelID, channelSecret string) *LinePay {
	return &LinePay{
		ChannelID:     channelID,
		ChannelSecret: channelSecret,
		IsSandBox:     true,
	}
}
