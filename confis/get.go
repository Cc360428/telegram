/**
 * @Author: Cc
 * @Description: 描述
 * @File: get
 * @Version: 1.0.0
 * @Date: 2022/10/12 17:35
 * @Software : GoLand
 */

package confis

func GetServer() (r [][]string) {
	Conf.config.Viper.UnmarshalKey("server", &r)
	return
}

func GetBaseURL() string {
	return Conf.config.Viper.GetString("baseURL")
}

func GetBotFunItem() (r []string) {
	Conf.config.Viper.UnmarshalKey("bot.FunItem", &r)
	return
}

/////////// gameAll

func GetGameAll() (r [][]string) {
	Conf.config.Viper.UnmarshalKey("GameAll.Info", &r)
	return
}
