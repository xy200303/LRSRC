package utils

import "regexp"

// 编译正则表达式以检查电子邮件格式
var emailRegexp = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

// 编译正则表达式以检查手机号码格式（假设中国大陆手机号码）
var phoneRegexp = regexp.MustCompile(`^1[3-9]\d{9}$`)

// 编译正则表达式以检查是否只包含数字和英文字母
var alphanumericRegexp = regexp.MustCompile(`^[A-Za-z0-9]+$`)

// IsEmailValid validateEmail 检查给定的字符串是否为有效的电子邮件地址
func IsEmailValid(email string) bool {
	return emailRegexp.MatchString(email)
}

// IsPhoneValid validatePhone 检查给定的字符串是否为有效的手机号码
func IsPhoneValid(phone string) bool {
	return phoneRegexp.MatchString(phone)
}

// IsAlphanumericValid 检查给定的字符串是否只包含数字和英文字母
func IsAlphanumericValid(input string) bool {
	return alphanumericRegexp.MatchString(input)
}
