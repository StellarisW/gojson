# gojson

[![build](https://img.shields.io/badge/build-1.01-brightgreen)](https://github.com/StellarisW/StellarisW)[![go-version](https://img.shields.io/badge/go-%3E%3D1.8-30dff3?logo=go)](https://github.com/StellarisW/StellarisW)[![Go Report Card](https://goreportcard.com/badge/github.com/emirpasic/gods)](https://goreportcard.com/report/github.com/emirpasic/gods)[![PyPI](https://img.shields.io/badge/License-BSD_2--Clause-green.svg)](https://github.com/emirpasic/gods/blob/master/LICENSE)

> 一个强大的json框架

# 💡  简介

gojson是一个支持数据多种方式读取，智能解析，操作便捷的一个json框架

# 🚀 功能

- `json`序列与反序列化
- `json`指定字段查找
- 支持`json`数据操作 (修改，插入，删除)
- `json`数据可视化

# 🌟 亮点

- 性能出色
    - `json`对象相对于其他项目携带的数据(字段)更少
    - 底层使用大量指针
    - 优化的递归函数
- 功能强大
    - 支持将任意形式`(嵌套,指针,切片,数组,map,空接口)`的结构体，任意形式的字段(`匿名`,`导出`,`不可导出`)unmarshal成`json`格式
    - 支持将其他格式`(yaml,toml,xml...)`的数据转化成`json`格式
    - 支持从文件读取数据
    - 支持`tag`校验
    - `json`友好输出
    - 并发安全
- 操作便捷
    - 函数链式操作
    - 智能解析数据
    - 可使用函数少，但功能强大


# ⚙ 代码结构

<details>
<summary>展开查看</summary>
<pre>
<code>
    ├── internal  ----------------------(内部工具包)
    	├── conv  ----------------------(数据转换)
    		├── byte.go
    		├── consts.go
    		├── interfaces.go
    		├── map.go
    		├── string.go
    	├── encoding  ------------------(编码包)
    		├── ini
    			├── ini.go
    		├── toml
    			├── toml.go
    		├── xml
    			├── xml.go
    		├── yaml
    			├── yaml.go
    	├── mutex  ---------------------(读写锁)
    		├── mutes.go
    	├── regex  ---------------------(正则匹配)
    		├── regex.go 
    	├── type  ----------------------(类型相关的操作包)
    		├── stringx
    			├── string.go
    		├── reflection
    			├── reflection.go
    		├── structx
    			├── field.go
    			├── structx.go
    ├── const.go  ----------------------(常量定义)
    ├── dump.go  -----------------------(数据可视化相关的函数)
    ├── err.go  ------------------------(错误定义)
    ├── gojson.go  ---------------------(用户可操作函数)
    ├── load.go  -----------------------(数据加载相关的函数)
    ├── operator.go  -------------------(json数据相关的操作函数)
    ├── option.go  ---------------------(选项相关的函数)
</code>
</pre>
</details>

# 🎬 快速开始

## 创建json对象(初始化)

```go
j:=gojson.New()
```

## json序列化

### 从对象中加载

```go
// 字符串,[]byte

jsonStr := `{"name": "json", "age": 18}`
yamlStr := `
name: `json`
age: 18
`

j1:=gojson.New().LoadContent(jsonStr)
j2:=gojson.New().LoadContent(yamlStr)

// 结构体
person:= struct {
		name string		
		age int
	}{name: "json",age: 18}
	
j3:=gojson.New().LoadContent(person)

// 切片
sli:=[]interface{}{"json",18,"male"}

j4:=gojson.New().LoadContent(sli)

// map

m:=map[string]interface{}{"name":"json","age":18}

j5:=gojson.New().LoadContent(m)
```

### 从文件中加载

```go
path:="./example.txt"
j:=gojson.New().LoadFile(path)
```

### 加载时可选选项

```go
//  Options
//  把选项结构体从Json对象分离出来,Json携带的数据更少,性能更优
//  type Options struct {
// 	  Safe           bool   // 需要并发安全时开启,使用读写锁
//	  ContentType    string // 设定数据类型,没有设定需要后面程序来判断
//	  StrNumber      bool   // 是否将数字判断为字符串
//	  LoadUnexported bool   // 是否加载不可导出字段 // TODO: 待支持
//  }
jsonStr := `{"name": "json", "age": 18}`
j:=gojson.New().LoadContentWithOptions(jsonStr, Options{
		Safe:           true,
		ContentType:    "json",
		StrNumber:      true,
		LoadUnexported: false,
	})
```

## json反序列化

```go
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

jsonStr := `{"name": "json", "age": 18}`
j:=gojson.New().LoadContent(jsonStr)
j.Unmarshal(Person)
```

## json可视化

### 输出json对象

```go
jsonStr := `{"name": "json", "age": 18}`
j:=gojson.New().LoadContent(jsonStr)
j.Dump()
/*
output:
	{
	    mu:          {
	        w:           {
	            state: 0,
	            sema:  0,
	        },
	        writerSem:   0,
	        readerSem:   0,
	        readerCount: 0,
	        readerWait:  0,
	    },
	    JsonContent: {
	        "name": "json",
	        "age":  18,
	    },
	    IsValid:     true,
	}
*/
```

### 输出json字符串

```go
jsonStr := `{"name": "json", "age": 18}`
j:=gojson.New().LoadContent(jsonStr)
j.DumpContent()
/*
output:
	{
    	"age":  18,
   	    "name": "json",
	}
*/
```

### 可选选项

```go
jsonStr := `{"name": "json", "age": 18}`
j := New().LoadContent(jsonStr)
j.DumpWithOptions(j.JsonContent,DumpOption{
	WithType:     true,
	ExportedOnly: false,
})
/*
output:
	map[string]interface {}(2) {
    	string("name"): string(4) "json",
    	string("age"):  float64(18),
	}
*/
```

## json操作

### 查找

```go
--------------------------------------------------
jsonStr := `{"name": "json", "age": 18}`
j := New().LoadContent(jsonStr)
fmt.Println(j.Get("name"))
/*
output:
	json
*/
--------------------------------------------------
jsonStr := `
[{
		"name":"json",
		"age":18
	},
	{
		"name":"Bob",
		"age":15
	}
]
`
j := New().LoadContent(jsonStr)
fmt.Println(j.Get("[0].name"))
/*
output:
	json
*/
--------------------------------------------------
```

### 修改

```go
jsonStr := `
[{
		"name":"json",
		"age":18
	},
	{
		"name":"Bob",
		"age":15
	}
]
`
j := New().LoadContent(jsonStr)
j.Set("[0].age", 20)
fmt.Println(j.Get("[0].age"))
/*
output:
	20
*/
```

### 插入

```go
jsonStr := `
[{
		"name":"json",
		"age":18
	},
	{
		"name":"Bob",
		"age":15
	}
]
`
j := New().LoadContent(jsonStr)
j.Set("[2]", map[string]interface{}{"name": "kate", "age": 21})
j.DumpContent()
/*
output:
	[
    	{
        	"name": "json",
        	"age":  18,
    	},
    	{
        	"name": "Bob",
        	"age":  15,
    	},
    	{
        	"name": "kate",
        	"age":  21,
    	},
	]
*/
```

### 删除

```go
jsonStr := `
[{
		"name":"json",
		"age":18
	},
	{
		"name":"Bob",
		"age":15
	}
]
`
j := New().LoadContent(jsonStr)
j.Set("[1]", nil)
j.DumpContent()
/*
output:
	[
    	{
        	"name": "json",
       		"age":  18,
    	},
	]
*/
```

# 📊 性能测试

# 🛠 环境要求

- golang 版本 >= 1.18

# 📌 TODO

- json的序列化

    - [x] `string`，`[]byte`的序列化

    - [x] `json`格式

    - [x] 其他类型的格式(`toml`，`yaml`，`xml`，`ini`)
        - [x] toml

        - [x] yaml

        - [ ] xml

        - [x] ini

    - [x] 结构体序列化

    - [x] 切片,数组序列化

    - [x] `map`序列化

    - [x] 多`tag`校验

    - [x] 从文件读取

    - [ ] 发起http请求并读取
- json的反序列化
    - [x] 映射到结构体
- json可视化
    - [x] map
    - [x] 切片
    - [x] 结构体
- json的数据操作

    - [x] 查找
    - [x] 修改
    - [x] 插入
    - [x] 删除

# 🎈 结语

# 📔 参考文献

[CSDN Golang自定义结构体转map](https://blog.csdn.net/pyf09/article/details/111027686?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522165856096916782395381810%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=165856096916782395381810&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~pc_rank_34-7-111027686-null-null.142^v33^pc_rank_34,185^v2^control&utm_term=go%20%E7%BB%93%E6%9E%84%E4%BD%93%E8%BD%AC%E6%8D%A2%E6%88%90map%5Bstring%5Dinterface%7B%7D&spm=1018.2226.3001.4187)

[GitHub structs](https://github.com/fatih/structs/)

[GitHub mapstructure](https://github.com/mitchellh/mapstructure)

# 🔑 JetBrains 开源证书支持

感谢JetBrains提供的Goland支持

<a href="https://www.jetbrains.com/?from=gnet" target="_blank"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="250" align="middle"/></a>
