package types

import "time"

type ResourceDTO struct {
	ID        string     `json:"id"`                                                             // m3u8 + name 的 md5 字符串
	URL       string     `json:"url"`                                                            // m3u8 url
	Name      string     `json:"name"`                                                           // 文件名, 不含后缀
	Status    string     `json:"status" validate:"oneof=waiting downloading unfinished success"` // 下载状态
	CreatedAt *time.Time `json:"created_at"`
}
