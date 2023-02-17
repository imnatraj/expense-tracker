package models

type MetaDataMode string
type MetaDataChannel string

const (
	// mode
	MetaDataModeWeb MetaDataMode = "web"
	MetaDataModeApp MetaDataMode = "app"

	// channel
	MetaDataChannelAdmin  MetaDataChannel = "admin"
	MetaDataChannelClient MetaDataChannel = "client"
)


type MetaData struct {
	Mode    MetaDataMode
	Channel MetaDataChannel
}
