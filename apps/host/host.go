package host

import (
	"context"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/imdario/mergo"
	"net/http"
	"strconv"
	"time"
)

type Service interface {
	 CreateHost(ctx context.Context ,host *Host)(*Host , error)
	 DescibeHost(ctx context.Context, request *DescribeHostRequest) (*Host , error)
	 UpdateHost(ctx context.Context ,request *UpdateHostRequest)(*Host , error)
	 QueryHost(ctx context.Context , request *QueryHostRequest)(*Host  , error)
	 DeleteHost(ctx context.Context , request *DeleteHostRequest)(*Host,error)
}

type HostSet struct {
	 Items []*Host
	 Total  int
}

var (
	 Validate = validator.New()
)


// 对象全量更新
func (h *Host) Put(obj *Host) error {
	if obj.Id != h.Id {
		return fmt.Errorf("id not equal")
	}

	*h.Resource = *obj.Resource
	*h.Describe = *obj.Describe
	return nil
}

// 对象的局部更新
func (h *Host) Patch(obj *Host) error {
	// if obj.Name != "" {
	// 	h.Name = obj.Name
	// }
	// if obj.CPU != 0 {
	// 	h.CPU = obj.CPU
	// }
	// 比如 obj.A  obj.B  只想修改obj.B该属性
	return mergo.MergeWithOverwrite(h, obj)
}

func(h *Host) Validate()error{
	 return Validate.Struct(h)
}


func(h *Host) InjectDeault(){
	if h.CreateAt == 0{
		h.CreateAt = time.Now().UnixMilli()
	}
}

type Host struct {
	*Resource // 公共
	*Describe // 独有属性
}

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

//vender 的值
type Vendor int
const (
	PrivateIDC Vendor = iota
	Tencent
	AliYun
	HuaWei
)

type Resource struct {
	Id          string            `json:"id"  validate:"required"`     // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`                      // 厂商
	Region      string            `json:"region"  validate:"required"` // 地域
	CreateAt    int64             `json:"create_at"`                   // 创建时间
	ExpireAt    int64             `json:"expire_at"`                   // 过期时间
	Type        string            `json:"type"  validate:"required"`   // 规格
	Name        string            `json:"name"  validate:"required"`   // 名称
	Description string            `json:"description"`                 // 描述
	Status      string            `json:"status"`                      // 服务商中的状态
	Tags        map[string]string `json:"tags"`                        // 标签
	UpdateAt    int64             `json:"update_at"`                   // 更新时间
	SyncAt      int64             `json:"sync_at"`                     // 同步时间
	Account     string            `json:"accout"`                      // 资源的所属账号
	PublicIP    string            `json:"public_ip"`                   // 公网IP
	PrivateIP   string            `json:"private_ip"`                  // 内网IP
}

type Describe struct {
	CPU          int    `json:"cpu" validate:"required"`    // 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU类型
	OSType       string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
}

func NewQueryHostFromHTTP(r *http.Request) *QueryHostRequest {
	req := NewQueryHostRequest()
	// query string
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		req.PageSize, _ = strconv.Atoi(pss)
	}

	pns := qs.Get("page_number")
	if pns != "" {
		req.PageNumber, _ = strconv.Atoi(pns)
	}

	req.Keywords = qs.Get("kws")
	return req
}

func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

type QueryHostRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
}

func (req *QueryHostRequest) GetPageSize() uint {
	return uint(req.PageSize)
}

func (req *QueryHostRequest) OffSet() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

func NewDescribeHostRequestWithId(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}

type DescribeHostRequest struct {
	Id string
}

type UPDATE_MODE string

const (
	// 全量更新
	UPDATE_MODE_PUT UPDATE_MODE = "put"
	// 局部更新
	UPDATE_MODE_PATCH UPDATE_MODE = "patch"
)

func NewPutUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_MODE_PUT,
		Host:       h,
	}
}

func NewPatchUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_MODE_PATCH,
		Host:       h,
	}
}

type UpdateHostRequest struct {
	UpdateMode UPDATE_MODE `json:"update_mode"`
	*Host
}

type DeleteHostRequest struct {
	Id string
}


