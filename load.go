package gojson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gojson/internal/encoding/ini"
	"gojson/internal/encoding/toml"
	"gojson/internal/encoding/xml"
	"gojson/internal/encoding/yaml"
	"gojson/internal/regex"
)

func (j *Json) parseContent(content []byte, options Options) (interface{}, error) {
	var err error
	if options.ContentType == "" {
		options.ContentType = getContentType(content)
		if options.ContentType == "" {
			return nil, errors.New(invalidContentType)
		}
	}
	switch options.ContentType {
	case ContentTypeJson:

	case ContentTypeXml:
		if content, err = xml.ToJson(content); err != nil {
			return nil, errors.New(invalidContentType)
		}
	case ContentTypeYaml:
		if content, err = yaml.ToJson(content); err != nil {
			return nil, errors.New(invalidContentType)
		}
	case ContentTypeToml:
		if content, err = toml.ToJson(content); err != nil {
			return nil, errors.New(invalidContentType)
		}
	case ContentTypeIni:
		if content, err = ini.ToJson(content); err != nil {
			return nil, errors.New(invalidContentType)
		}
	}

	// 使用json decoder将数据解码成map[string]interface{}形式

	var jsonContent interface{}
	decoder := json.NewDecoder(bytes.NewReader(content))
	// 解码时是否将数字看作字符
	if options.StrNumber {
		decoder.UseNumber()
	}
	if err := decoder.Decode(&jsonContent); err != nil {
		fmt.Printf("%v, err: %v", decodeErr, err)
		return nil, errors.New(decodeErr + "->" + err.Error())
	}
	switch jsonContent.(type) {
	// 解码器没有把数据解析成map[string]interface{}的情况
	case string, []byte:
		return nil, errors.New(decodeErr)
	}
	// 携带解析完后的jsonContent递归下去
	return jsonContent, nil
}

// getContentType 通过正则表达式判断数据的格式
func getContentType(content []byte) string {
	if json.Valid(content) {
		return ContentTypeJson
	} else if regex.CheckXml(content) {
		return ContentTypeXml
	} else if regex.CheckYaml(content) {
		return ContentTypeYaml
	} else if regex.CheckToml(content) {
		return ContentTypeToml
	} else if regex.CheckIni(content) {
		return ContentTypeIni
	} else {
		return ""
	}
}
