/**
 * @Author: Cc
 * @Description: 描述
 * @File: utils_test.go
 * @Version: 1.0.0
 * @Date: 2022/11/24 17:25
 * @Software : GoLand
 */

package utils

import (
	"telegram_bot/confis"
	"testing"
)

func TestQuery(t *testing.T) {
	confis.InitConfigStart()
	t.Log(Query("bqq_current_player_Count", "bandarqq_pro_a"))
}
