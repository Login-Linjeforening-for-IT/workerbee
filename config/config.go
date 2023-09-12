package config

import (
	"fmt"
	"io/fs"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var ErrConfigTypeNotStruct = fmt.Errorf("config type is not struct")

type configField struct {
	tag          string
	defaultValue string
}

// using reflect to bind env variables to config struct
//
//	and set default values based on tags
func bind[T any](v *viper.Viper, tagName string, withDefaults bool) error {
	var dummy T
	t := reflect.TypeOf(dummy)

	for kind := t.Kind(); kind == reflect.Pointer; kind = t.Kind() { // Ensuring no pointer
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return ErrConfigTypeNotStruct
	}

	numFields := t.NumField()
	fields := make([]configField, 0, numFields)
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup(tagName)
		if !ok || tag == "" { // no mapstructure tag found, using field name instead
			tag = field.Name
		}

		defaultValue, ok := field.Tag.Lookup("default")
		if !ok || defaultValue == "" {
			defaultValue = ""
		}

		fields = append(fields, configField{
			tag:          tag,
			defaultValue: defaultValue,
		})
	}

	for _, field := range fields {
		// BindEnv has the capability to bind any number of env variables at once using call
		// BindEnv will fail if no arguments are passed i.e. viper.BindEnv()
		// we can ignore this error because we explicitly pass an argument
		_ = v.BindEnv(field.tag)
		if withDefaults && field.defaultValue != "" {
			v.SetDefault(field.tag, field.defaultValue)
		}
	}

	return nil
}

func Must[T any](obj T, err error) T {
	Require(err)
	return obj
}

func Require(err error) {
	if err != nil {
		panic(err)
	}
}

type loadOpts struct {
	filepath    string
	requireFile bool // require config file to exist

	withDefaults bool
	tagName      string
	envPrefix    string
}

func defaultLoadOpts() *loadOpts {
	return &loadOpts{
		filepath:     "",
		requireFile:  false,
		withDefaults: true,
		tagName:      "config",
		envPrefix:    "",
	}
}

type loadOpt func(*loadOpts)

func WithFile(path string) loadOpt {
	return func(opts *loadOpts) {
		opts.filepath = path
		opts.requireFile = false
	}
}

func WithRequireFile(path string) loadOpt {
	return func(opts *loadOpts) {
		opts.filepath = path
		opts.requireFile = true
	}
}

func WithTagName(name string) loadOpt {
	return func(opts *loadOpts) {
		opts.tagName = name
	}
}

func WithEnvPrefix(prefix string) loadOpt {
	return func(opts *loadOpts) {
		opts.envPrefix = prefix
	}
}

var WithoutDefaults loadOpt = func(opts *loadOpts) {
	opts.withDefaults = false
}

func LoadInto[T any](obj *T, opts ...loadOpt) error {
	v := viper.New()
	conf := defaultLoadOpts()

	for _, opt := range opts {
		opt(conf)
	}

	err := bind[T](v, conf.tagName, conf.withDefaults)
	if err != nil {
		return err
	}

	v.AutomaticEnv()

	if conf.filepath != "" {
		v.SetConfigFile(conf.filepath)
	}

	if conf.envPrefix != "" {
		v.SetEnvPrefix(conf.envPrefix)
	}

	err = v.ReadInConfig()
	if err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError, *fs.PathError:
			if conf.requireFile {
				return err
			}
		case viper.UnsupportedConfigError:
			return err
		default:
			return err
		}
	}

	return v.Unmarshal(obj, func(c *mapstructure.DecoderConfig) {
		c.TagName = conf.tagName
	})
}

func Load[T any](opts ...loadOpt) (*T, error) {
	var t T
	err := LoadInto(&t, opts...)
	return &t, err
}

func MustLoadInto[T any](obj *T, opts ...loadOpt) {
	Require(LoadInto(obj, opts...))
}

func MustLoad[T any](opts ...loadOpt) *T {
	return Must(Load[T](opts...))
}
