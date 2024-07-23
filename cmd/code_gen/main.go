package main

import (
	"fmt"
	"games/app/global"
	"games/app/provider"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

func main() {
	provider.Register()

	genSystemDb()
	genGameDb()
}

func genSystemDb() {
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath:      "app/dal/system/repository",
		OutFile:      "gen.go",
		ModelPkgPath: "model",
		WithUnitTest: false,
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true,
		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: true,
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: false,
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		// Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		Mode: 0,
	})

	// 自定义字段的数据类型
	// 如果需要兼容protobuf，请统一数字类型为int64,
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `xxxxx_field, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{jsonField}

	g.UseDB(global.DB.Jenny)
	tables := [...]string{
		"users",
		"category",
		"article",
		"image",
	}
	for _, value := range tables {
		a := g.GenerateModel(value, fieldOpts...)
		g.ApplyBasic(a)
	}
	g.Execute()
}

func genGameDb() {
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath:      "app/dal/game/repository",
		OutFile:      "gen.go",
		ModelPkgPath: "model",
		WithUnitTest: false,
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true,
		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: true,
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: false,
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		// Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		Mode: 0,
	})

	// 自定义字段的数据类型
	// 如果需要兼容protobuf，请统一数字类型为int64,
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `xxxxx_field, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{jsonField}
	g.UseDB(global.DB.Jenny)
	tables := [...]string{
		"system",
		"steam_game_image",
		"steam_game_video",
	}
	for _, value := range tables {
		a := g.GenerateModel(value, fieldOpts...)
		g.ApplyBasic(a)
	}
	g.Execute()
}

func example() {
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath:      "app/dal/repository",
		OutFile:      "gen.go",
		ModelPkgPath: "model",
		WithUnitTest: false,
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true,
		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: true,
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: false,
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		// Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		Mode: 0,
	})
	fmt.Println("-------------------")
	fmt.Println(global.DB.Jenny)

	// 自定义字段的数据类型
	// 如果需要兼容protobuf，请统一数字类型为int64,
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `xxxxx_field, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{jsonField}

	g.UseDB(global.DB.Jenny)
	tables := [...]string{
		"users",
		"category",
		"article",
		"system",
		"steam_game_image",
		"steam_game_video",
	}
	for _, value := range tables {
		a := g.GenerateModel(value, fieldOpts...)
		g.ApplyBasic(a)
	}
	g.Execute()

	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`created_time`、`updated_time`, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_time`, 表字段数据类型为: DATETIME
	// autoCreateTimeField := gen.FieldGORMTag("created_time", "column:created_time;type:datetime;autoCreateTime")
	// autoUpdateTimeField := gen.FieldGORMTag("updated_time", "column:updated_time;type:datetime;autoUpdateTime")
	// softDeleteField := gen.FieldType("deleted_time", "deleted_time.DeletedAt")

	// 创建全部模型文件, 并覆盖前面创建的同名模型
	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
	//allModel := g.GenerateAllTable(fieldOpts...)
	//fmt.Println(allModel)

	/**
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	Address := g.GenerateModel("address")
	// 创建有关联关系的模型文件
	User := g.GenerateModel("game",
		append(
			fieldOpts,
			// game 一对多 address 关联, 外键`uid`在 address 表中
			gen.FieldRelate(field.HasMany, "Address", Address, &field.RelateConfig{GORMTag: "foreignKey:UID"}),
		)...,
	)
	Address = g.GenerateModel("address",
		append(
			fieldOpts,
			gen.FieldRelate(field.BelongsTo, "User", User, &field.RelateConfig{GORMTag: "foreignKey:UID"}),
		)...)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(User, Address)
	*/

	// Generate struct `User` based on table `users`
	//g.GenerateModel("users", fieldOpts...)

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(genModel.User{})
	//g.ApplyBasic(allModel...)

	g.UseDB(global.DB.Xh)
	allModel := g.GenerateAllTable(fieldOpts...)
	g.ApplyBasic(allModel...)
	g.Execute()
}
