package app

import "giligili/app/config"

var Config config.Conf

func Init() error {

	return Config.ConfInit()
}
