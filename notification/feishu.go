package notification

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type FeishuCard struct {
	HeaderColor string             `json:"header_color"`
	HookUrl     string             `json:"hook_url" binding:"required"`
	Title       string             `json:"title" binding:"required"`
	Content     string             `json:"content"`
	Images      []string           `json:"images"`
	Fields      []FeishuCardField  `json:"fields"`
	Remark      string             `json:"remark"`
	Buttons     []FeishuCardButton `json:"buttons"`
}

type FeishuCardField struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type FeishuCardButton struct {
	Text string `json:"text" binding:"required"`
	Url  string `json:"url" binding:"required"`
}

// 发送群机器人卡片消息
func SendGroupCardMsg(params FeishuCard) (err error) {
	if params.HookUrl == "" {
		err = errors.New("hook url is empty")
		return
	}
	if params.HeaderColor == "" {
		params.HeaderColor = "blue"
	}

	// 请求体
	data := map[string]any{
		"msg_type": "interactive",
	}

	elements := []any{}

	// 处理 fields
	if params.Fields != nil {
		fields := []map[string]any{}
		for _, field := range params.Fields {
			fields = append(fields, map[string]any{
				"is_short": true,
				"text": map[string]any{
					"tag":     "lark_md",
					"content": fmt.Sprintf("**%s: **\n%s", field.Name, field.Value),
				},
			})
		}

		elements = append(elements, map[string]any{
			"tag":    "div",
			"fields": fields,
		})
	}

	// 处理富文本内容
	if params.Content != "" {
		elements = append(elements, map[string]any{
			"tag":     "markdown",
			"content": params.Content,
		})
	}

	if params.Images != nil {
		for _, image := range params.Images {
			elements = append(elements, map[string]any{
				"tag":     "img",
				"img_key": image,
			})
		}
	}

	// 分割线
	if len(params.Buttons) > 0 || params.Remark != "" {
		elements = append(elements, map[string]any{
			"tag": "hr",
		})
	}

	// 按钮
	if len(params.Buttons) > 0 {
		buttons := []map[string]any{}
		for _, button := range params.Buttons {
			buttons = append(buttons, map[string]any{
				"tag": "button",
				"text": map[string]any{
					"tag":     "plain_text",
					"content": button.Text,
				},
				"type": "primary",
				"multi_url": map[string]string{
					"url":         button.Url,
					"pc_url":      "",
					"android_url": "",
					"ios_url":     "",
				},
			})
		}

		elements = append(elements, map[string]any{
			"tag":     "action",
			"actions": buttons,
		})
	}

	// 备注
	if params.Remark != "" {
		elements = append(elements, map[string]any{
			"tag": "note",
			"elements": []any{
				map[string]string{
					"tag":     "plain_text",
					"content": params.Remark,
				},
			},
		})
	}

	card := map[string]any{
		"header": map[string]any{
			"template": params.HeaderColor,
			"title": map[string]any{
				"content": params.Title,
				"tag":     "plain_text",
			},
		},
		"config": map[string]any{
			"wide_screen_mode": true,
		},
	}

	card["elements"] = elements
	data["card"] = card
	str, _ := json.Marshal(data)

	_, err = resty.New().SetDebug(viper.GetBool("dev")).R().
		SetBody(string(str)).
		Post(params.HookUrl)
	if err != nil {
		log.Printf("发送飞书群机器人消息失败：%v", err)
	}

	return
}
