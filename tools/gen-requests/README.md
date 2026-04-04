# gen-requests - Request 结构体生成器

从 Model 结构体自动生成 Request 结构体的工具。

## 功能特性

- 自动从 Model 生成 CreateRequest 和 UpdateRequest
- 自动添加 JSON tag（蛇形命名）
- 自动添加 validate tag
- 自动生成 ToModel() 和 UpdateModel() 方法
- 支持 include/exclude 过滤
- 使用 copier 进行结构体复制

## 使用方法

### 1. 编译工具

```bash
cd tools/gen-requests
go build -o ../../bin/gen-requests .
```

### 2. 直接使用

```bash
# 生成单个 Model 的 Request
./bin/gen-requests internal/app/server/model/video.go

# 生成并输出到指定文件
./bin/gen-requests -output internal/app/server/handler/video_requests_gen.go internal/app/server/model/video.go

# 只包含特定 Model
./bin/gen-requests -include Video,Actress internal/app/server/model/model.go

# 排除特定 Model
./bin/gen-requests -exclude BWH internal/app/server/model/actress.go
```

### 3. 使用 go:generate（推荐）

在 Model 文件顶部添加：

```go
//go:generate ../../../../bin/gen-requests -output ../handler/video_requests_gen.go video.go
```

然后运行：

```bash
cd internal/app/server/model
go generate
```

## 生成的内容

工具会为每个 Model 生成：

1. `{Model}CreateRequest` - 创建请求结构体，所有字段都有 `validate:"required"`
2. `{Model}UpdateRequest` - 更新请求结构体，包含 ID 字段
3. `ToModel() *model.{Model}` - 转换为 Model 方法
4. `UpdateModel(m *model.{Model})` - 更新 Model 方法

## 示例

输入（Model）：

```go
type Video struct {
    BaseModel
    SerialNumber string
    CoverPath    string
    Title        *string
}
```

输出（Request）：

```go
type VideoCreateRequest struct {
    SerialNumber string `json:"serial_number" validate:"required"`
    CoverPath    string `json:"cover_path" validate:"required"`
    Title        *string `json:"title" validate:"required"`
}

type VideoUpdateRequest struct {
    ID           uint   `json:"id" validate:"required"`
    SerialNumber string `json:"serial_number"`
    CoverPath    string `json:"cover_path"`
    Title        *string `json:"title"`
}
```
