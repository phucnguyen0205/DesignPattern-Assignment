# Bài thi: Adapter Pattern

```text
Họ và tên: Nguyễn Hữu Phúc
Mã sinh viên: 1440501245922
Lớp: SE1-K24
Bài thi: Adapter Pattern
```

## Việc cần làm

Đọc `assignment.md` trong bộ đề. Hoàn thành các TODO trong `main.go`: tạo `SMSAdapter` chứa `*LegacySMS`, triển khai `Notifier.Send()` và dùng adapter qua biến kiểu `Notifier` trong `main()`. Không sửa thư viện `LegacySMS`.

Starter hiện biên dịch được nhưng chưa gửi thông báo. Không yêu cầu unit test, database hay HTTP API.

## Kiểm tra và chạy

Từ thư mục chứa `go.mod` (Go 1.22 trở lên):

```bash
go build ./...
go run .
```

Sau khi hoàn thành, kết quả cần là:

```text
SMS: Xin chao, don hang cua ban da duoc xac nhan!
```

## Câu trả lời của sinh viên

Mỗi câu trả lời khoảng 2–3 dòng.

1. Adapter Pattern giải quyết vấn đề gì?

   Adapter Pattern giải quyết vấn đề không tương thích giữa các interface khác nhau. Nó cho phép kết nối class có interface không tương thích (như LegacySMS với method SendSMS) vào hệ thống yêu cầu một interface khác (Notifier với method Send). Nhờ đó, code mới có thể sử dụng thư viện cũ mà không cần sửa đổi code gốc.

2. Tại sao không cần chỉnh sửa `LegacySMS`?

   Không sửa LegacySMS vì đây là thư viện đã được kiểm thử và triển khai sẵn, việc sửa có thể gây hỏng các chức năng khác đang dùng nó. Ngoài ra, trong thực tế thư viện có thể do bên thứ ba cung cấp hoặc nằm trong codebase cũ không nên thay đổi. Adapter Pattern cho phép bọc ngoài và chuyển đổi interface mà giữ nguyên logic bên trong của thư viện gốc.

## Nộp bài

Nộp `main.go`, `go.mod`, `README.md` và link GitHub repository. Điền đủ thông tin cá nhân phía trên và bảo đảm người chấm truy cập được repository.

Link repository: https://github.com/phucnguyen0205/DesignPattern-Assignment.git
# DesignPattern-Assignment
