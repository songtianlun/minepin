package cfg

import (
	"fmt"
	"minepin/com/constvar"
	"minepin/com/utils"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestRegisterCfg(t *testing.T) {
	testCases := map[string]interface{}{
		"testKey1":      "testValue",
		"testKey2":      123,
		"testKey3":      true,
		"testKey4":      int64(123),
		"test.testKey4": int64(123),
	}
	fmt.Println(utils.GetTypeString(testCases["testKey3"]))
	cfgStrKey := "testKey"
	cfgValue := "testValue"
	cfgEnvValue := "testEnvValue"

	for k, v := range testCases {
		RegisterCfg(k, v, reflect.TypeOf(v).String())
		os.Setenv(strings.ToUpper(k), fmt.Sprintf("%v", v))
	}

	RegisterCfg(cfgStrKey, cfgValue, "string")
	os.Setenv(strings.ToUpper(constvar.DefaultCfgEnvPrefix+"_"+cfgStrKey), cfgEnvValue)

	err := Init("")
	if err != nil {
		return
	}

	for k, v := range testCases {
		if utils.GetTypeString(v) == "string" {
			res := GetString(k)
			if res != v {
				t.Errorf("GetCfg(%s) = %v, want %v", k, res, v)
			}
		} else if utils.GetTypeString(v) == "int" {
			res := GetInt(k)
			if res != v {
				t.Errorf("GetCfg(%s) = %v, want %v", k, res, v)
			}
		} else if utils.GetTypeString(v) == "bool" {
			res := GetBool(k)
			if res != v {
				t.Errorf("GetCfg(%s) = %v, want %v", k, res, v)
			}
		} else if utils.GetTypeString(v) == "int64" {
			res := GetInt64(k)
			if res != v {
				t.Errorf("GetCfg(%s) = %v, want %v", k, res, v)
			}
		} else {
			t.Errorf("cfg OSType Error")
		}
	}

	resCfg := GetString(cfgStrKey)

	if resCfg != cfgEnvValue {
		t.Errorf("Set config Failed, want: [%s], got: [%s]", cfgEnvValue, resCfg)
	}
}
