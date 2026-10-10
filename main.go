package main

import "fmt"

// Notifier là interface mà ứng dụng sử dụng để gửi thông báo.
type Notifier interface {
	Send(message string) error
}

// LegacySMS là thư viện cũ. Không chỉnh sửa struct và phương thức này.
type LegacySMS struct{}

func (s *LegacySMS) SendSMS(content string) {
	fmt.Println("SMS:", content)
}

// SMSAdapter wraps LegacySMS to implement the Notifier interface.
type SMSAdapter struct {
	legacy *LegacySMS
}

// Send implements the Notifier interface by calling LegacySMS.SendSMS.
func (a *SMSAdapter) Send(message string) error {
	a.legacy.SendSMS(message)
	return nil
}

func main() {
	// Khởi tạo LegacySMS và SMSAdapter.
	legacy := &LegacySMS{}
	adapter := &SMSAdapter{legacy: legacy}

	// Gán adapter cho biến kiểu Notifier.
	var notifier Notifier = adapter

	// Gọi Send với thông báo yêu cầu:
	// "Xin chao, don hang cua ban da duoc xac nhan!"
	notifier.Send("Xin chao, don hang cua ban da duoc xac nhan!")
}