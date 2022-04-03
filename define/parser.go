package define

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
)

type Parser interface {
	Parse(definePath string) (Define, error)
}

type ParserFunc func(string) (Define, error)

func (f ParserFunc) Parse(definePath string) (Define, error) {
	return f(definePath)
}

func NewParser() Parser {
	return NewParserWithFileReader(ioutil.ReadFile)
}

func NewParserWithFileReader(f FileReaderFunc) Parser {
	return parser{f}
}

type FileReaderFunc func(string) ([]byte, error)

type parser struct {
	fileReader FileReaderFunc
}

func (p parser) Parse(configPath string) (Define, error) {
	var res Define
	data, err := p.fileReader(configPath)
	if err != nil {
		logrus.Errorf("read define file content fail! path:%s, err:%s", configPath, err)
		return res, err
	}

	viper.SetConfigType(path.Ext(configPath)[1:])
	if err = viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		logrus.Errorf("read define fail! path:%s, data:%s, err:%s", configPath, data, err)
		return res, err
	}
	if err = viper.Unmarshal(&res); err != nil {
		logrus.Errorf("unmarshal define fail! path:%s, data:%s, err:%s", configPath, data, err)
		return res, err
	}

	return res, nil
}
