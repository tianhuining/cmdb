package conf_test

import (
	"cmdb/conf"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadTomlToConfig(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadTomlToConfig("../etc/config.toml")
	if should.NoError(err){
		 fmt.Println(conf.Conf().App.Name)
	}
}

func TestMySQL_GetDB(t *testing.T) {
	should :=assert.New(t)

	err := conf.LoadTomlToConfig("../etc/config.toml")
	if should.NoError(err) {
		conf.Conf().MySQL.GetDB()
	}

}

