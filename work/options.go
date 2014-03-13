package work

import (
	"errors"
	"github.com/relops/amqp"
)

var (
	ErrInvalidOptions = errors.New("invalid options")
)

type Options struct {
	Direction  string  `short:"d" long:"direction" description:"Use rmq to send (-d in) or receive (-d out) messages" required:"true"`
	Exchange   string  `short:"x" long:"exchange" description:"The exchange to send to (-d in) or bind a queue to when receiving (-d out)"`
	Queue      string  `short:"q" long:"queue" description:"The queue to receive from (when used with -d in)"`
	Persistent bool    `short:"P" long:"persistent" description:"Use persistent messaging" default:"false"`
	NoDeclare  bool    `short:"n" long:"no-declare" description:"If set, then don't attempt to declare the queue or bind it" default:"false"`
	Key        string  `short:"k" long:"key" description:"The key to use for routing (-d in) or for queue binding (-d out)"`
	Count      int     `short:"c" long:"count" description:"The number of messages to send" default:"10"`
	Interval   int     `short:"i" long:"interval" description:"The delay (in ms) between sending messages" default:"10"`
	Size       float64 `short:"z" long:"size" description:"Message size in kB" default:"1"`
	StdDev     int     `short:"t" long:"stddev" description:"Standard deviation of message size" default:"0"`
	Renew      bool    `short:"r" long:"renew" description:"Automatically resubscribe when the server cancels a subscription (used for mirrored queues)" default:"false"`
	Username   string  `short:"u" long:"user" description:"The user to connect as" default:"guest"`
	Password   string  `short:"w" long:"pass" description:"The user's password" default:"guest"`
	Host       string  `short:"H" long:"host" description:"The Rabbit host to connect to" default:"localhost"`
	Port       int     `short:"p" long:"port" description:"The Rabbit port to connect on" default:"5672"`
	Entropy    bool    `short:"e" long:"entropy" description:"Display message level entropy information" default:"false"`
	Version    func()  `short:"V" long:"version" description:"Print rmq version and exit"`
}

func (o *Options) Validate() error {
	if o.Direction != "in" && o.Direction != "out" {
		return ErrInvalidOptions
	}
	if o.Size < 1 {
		return ErrInvalidOptions
	}
	if o.StdDev < 0 {
		return ErrInvalidOptions
	}
	return nil
}

func (o *Options) IsSender() bool {
	return o.Direction == "in"
}

func (o *Options) uri() string {
	u := &amqp.URI{
		Username: o.Username,
		Password: o.Password,
		Host:     o.Host,
		Port:     o.Port,
		Scheme:   "amqp",
	}
	return u.String()
}
