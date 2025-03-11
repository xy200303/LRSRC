package template

import (
	"bytes"
	"fmt"
	"html/template"
)

// 获取渲染模板的HTML
func RenderHtml(templateFile string, data interface{}) (string, error) {
	tmpl := template.Must(template.ParseFiles(templateFile))
	var renderedTemplate bytes.Buffer
	err := tmpl.Execute(&renderedTemplate, data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return "", err
	}
	// 获取渲染后的 HTML 字符串
	htmlContent := renderedTemplate.String()
	return htmlContent, nil
}

// RenderRegisterHtml 渲染注册邮箱
func RenderRegisterHtml(code string, duration int, sender string) (string, error) {
	data := &map[string]interface{}{
		"Code":     code,
		"Duration": duration,
		"Sender":   sender,
	}
	return RenderHtml("backend/templates/email_template/register_template.html", data)
}

// RenderForgetHtml 渲染注册邮箱
func RenderForgetHtml(code string, duration int, sender string) (string, error) {
	data := &map[string]interface{}{
		"Code":     code,
		"Duration": duration,
		"Sender":   sender,
	}
	return RenderHtml("backend/templates/email_template/forget_template.html", data)
}
