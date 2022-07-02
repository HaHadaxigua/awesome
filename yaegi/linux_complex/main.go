package main

import (
	"awesome/yaegi/linux_complex/types"
	_ "embed"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"reflect"
)

//go:generate yaegi extract github.com/spf13/cast
//go:generate yaegi extract awesome/yaegi/linux_complex/types

//go:embed script/fib.go
var fibScript string

var Symbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["github.com/spf13/cast/cast"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"StringToDate":                  reflect.ValueOf(cast.StringToDate),
		"StringToDateInDefaultLocation": reflect.ValueOf(cast.StringToDateInDefaultLocation),
		"ToBool":                        reflect.ValueOf(cast.ToBool),
		"ToBoolE":                       reflect.ValueOf(cast.ToBoolE),
		"ToBoolSlice":                   reflect.ValueOf(cast.ToBoolSlice),
		"ToBoolSliceE":                  reflect.ValueOf(cast.ToBoolSliceE),
		"ToDuration":                    reflect.ValueOf(cast.ToDuration),
		"ToDurationE":                   reflect.ValueOf(cast.ToDurationE),
		"ToDurationSlice":               reflect.ValueOf(cast.ToDurationSlice),
		"ToDurationSliceE":              reflect.ValueOf(cast.ToDurationSliceE),
		"ToFloat32":                     reflect.ValueOf(cast.ToFloat32),
		"ToFloat32E":                    reflect.ValueOf(cast.ToFloat32E),
		"ToFloat64":                     reflect.ValueOf(cast.ToFloat64),
		"ToFloat64E":                    reflect.ValueOf(cast.ToFloat64E),
		"ToInt":                         reflect.ValueOf(cast.ToInt),
		"ToInt16":                       reflect.ValueOf(cast.ToInt16),
		"ToInt16E":                      reflect.ValueOf(cast.ToInt16E),
		"ToInt32":                       reflect.ValueOf(cast.ToInt32),
		"ToInt32E":                      reflect.ValueOf(cast.ToInt32E),
		"ToInt64":                       reflect.ValueOf(cast.ToInt64),
		"ToInt64E":                      reflect.ValueOf(cast.ToInt64E),
		"ToInt8":                        reflect.ValueOf(cast.ToInt8),
		"ToInt8E":                       reflect.ValueOf(cast.ToInt8E),
		"ToIntE":                        reflect.ValueOf(cast.ToIntE),
		"ToIntSlice":                    reflect.ValueOf(cast.ToIntSlice),
		"ToIntSliceE":                   reflect.ValueOf(cast.ToIntSliceE),
		"ToSlice":                       reflect.ValueOf(cast.ToSlice),
		"ToSliceE":                      reflect.ValueOf(cast.ToSliceE),
		"ToString":                      reflect.ValueOf(cast.ToString),
		"ToStringE":                     reflect.ValueOf(cast.ToStringE),
		"ToStringMap":                   reflect.ValueOf(cast.ToStringMap),
		"ToStringMapBool":               reflect.ValueOf(cast.ToStringMapBool),
		"ToStringMapBoolE":              reflect.ValueOf(cast.ToStringMapBoolE),
		"ToStringMapE":                  reflect.ValueOf(cast.ToStringMapE),
		"ToStringMapInt":                reflect.ValueOf(cast.ToStringMapInt),
		"ToStringMapInt64":              reflect.ValueOf(cast.ToStringMapInt64),
		"ToStringMapInt64E":             reflect.ValueOf(cast.ToStringMapInt64E),
		"ToStringMapIntE":               reflect.ValueOf(cast.ToStringMapIntE),
		"ToStringMapString":             reflect.ValueOf(cast.ToStringMapString),
		"ToStringMapStringE":            reflect.ValueOf(cast.ToStringMapStringE),
		"ToStringMapStringSlice":        reflect.ValueOf(cast.ToStringMapStringSlice),
		"ToStringMapStringSliceE":       reflect.ValueOf(cast.ToStringMapStringSliceE),
		"ToStringSlice":                 reflect.ValueOf(cast.ToStringSlice),
		"ToStringSliceE":                reflect.ValueOf(cast.ToStringSliceE),
		"ToTime":                        reflect.ValueOf(cast.ToTime),
		"ToTimeE":                       reflect.ValueOf(cast.ToTimeE),
		"ToTimeInDefaultLocation":       reflect.ValueOf(cast.ToTimeInDefaultLocation),
		"ToTimeInDefaultLocationE":      reflect.ValueOf(cast.ToTimeInDefaultLocationE),
		"ToUint":                        reflect.ValueOf(cast.ToUint),
		"ToUint16":                      reflect.ValueOf(cast.ToUint16),
		"ToUint16E":                     reflect.ValueOf(cast.ToUint16E),
		"ToUint32":                      reflect.ValueOf(cast.ToUint32),
		"ToUint32E":                     reflect.ValueOf(cast.ToUint32E),
		"ToUint64":                      reflect.ValueOf(cast.ToUint64),
		"ToUint64E":                     reflect.ValueOf(cast.ToUint64E),
		"ToUint8":                       reflect.ValueOf(cast.ToUint8),
		"ToUint8E":                      reflect.ValueOf(cast.ToUint8E),
		"ToUintE":                       reflect.ValueOf(cast.ToUintE),
	}
	Symbols["awesome/yaegi/linux_complex/types/types"] = map[string]reflect.Value{
		// type definitions
		"Os": reflect.ValueOf((*types.Os)(nil)),
		// function, constant and variable definitions
		"AwsOs": reflect.ValueOf(&types.AwsOs).Elem(),
	}
}

func main() {
	var err error
	i := interp.New(interp.Options{
		Unrestricted: true,
	})

	if err = i.Use(stdlib.Symbols); err != nil {
		logrus.Error("failed to use stdlib")
		panic(err)
	}

	if err = i.Use(Symbols); err != nil {
		logrus.Error("failed to import third package")
		panic(err)
	}

	if _, err := i.Eval(fibScript); err != nil {
		panic(err)
	}

	v, err := i.Eval("script.Main")
	if err != nil {
		logrus.Error("failed to load main function")
		panic(err)
	}
	fn := v.Interface().(func())
	fn()

	getR, err := i.Eval("script.GiveInvoker")
	if err != nil {
		logrus.Error("failed to load func GiveInvoker")
	}
	getRFn := getR.Interface().(func() *struct {
		Code    int
		Message string
	})
	rFn := getRFn()
	fmt.Println(rFn)
}
