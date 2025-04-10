# System overview - Mô tả toàn bộ hệ thống
🛒 Trang web bán vợt cầu lông – Danh sách chức năng backend

## ✅ MVP – Chức năng bắt buộc

- [ ] **Quản lý sản phẩm**
  - [ ] Admin thêm/sửa/xoá sản phẩm
  - [ ] Người dùng xem danh sách & chi tiết sản phẩm

- [ ] **Đăng nhập/Đăng xuất người dùng**
  - [ ] Đăng nhập bằng email/password (JWT)
  - [ ] Refresh Token để duy trì phiên
  - [ ] Đăng xuất & xoá token
  - [ ] Lưu thông tin người dùng

- [ ] **Giỏ hàng (Cart)**
  - [ ] Thêm sản phẩm vào giỏ
  - [ ] Xoá sản phẩm khỏi giỏ
  - [ ] Cập nhật số lượng sản phẩm
  - [ ] Lưu giỏ hàng theo người dùng (session/token)

- [ ] **Đặt hàng (Order)**
  - [ ] Tạo đơn hàng từ giỏ
  - [ ] Lưu đơn hàng vào database
  - [ ] Gửi mail xác nhận đơn hàng

- [ ] **Quản lý đơn hàng (Admin)**
  - [ ] Xem danh sách đơn hàng
  - [ ] Cập nhật trạng thái đơn hàng (đã duyệt, đang giao,...)

---

## 🔄 Essential Features – Nên có sớm

- [ ] **Tìm kiếm sản phẩm**
  - [ ] Tìm theo tên sản phẩm

- [ ] **Lọc/sắp xếp sản phẩm**
  - [ ] Lọc theo loại, thương hiệu, giá
  - [ ] Sắp xếp theo giá, đánh giá

- [ ] **Lịch sử đơn hàng (người dùng)**
  - [ ] Xem danh sách đơn hàng đã đặt
  - [ ] Xem chi tiết đơn hàng

- [ ] **Đánh giá sản phẩm**
  - [ ] Viết đánh giá sau khi mua
  - [ ] Hiển thị đánh giá công khai

- [ ] **Quản lý kho (Inventory)**
  - [ ] Hiển thị số lượng tồn kho
  - [ ] Giảm số lượng khi đặt hàng

- [ ] **Thông báo (Notification)**
  - [ ] Gửi mail khi đặt hàng thành công
  - [ ] Thông báo khuyến mãi (email, giao diện)

- [ ] **Giao diện quản trị (CMS)**
  - [ ] Giao diện quản lý sản phẩm
  - [ ] Giao diện quản lý đơn hàng
  - [ ] Phân quyền người dùng

---

## 🚀 Advanced Features – Nâng cấp khi cần

- [ ] **Thanh toán online**
  - [ ] Tích hợp Momo/VNPay
  - [ ] Tạo hóa đơn & xác nhận thanh toán

- [ ] **Khuyến mãi/mã giảm giá**
  - [ ] Tạo mã giảm giá
  - [ ] Áp dụng mã khi thanh toán

- [ ] **Yêu thích sản phẩm (Wishlist)**
  - [ ] Lưu danh sách yêu thích
  - [ ] Hiển thị ở trang cá nhân

- [ ] **Tracking đơn hàng**
  - [ ] Cập nhật trạng thái giao hàng
  - [ ] Hiển thị tiến độ giao hàng

- [ ] **Blog / Tin tức**
  - [ ] Viết bài tư vấn, review sản phẩm
  - [ ] Danh sách bài viết

- [ ] **Chatbot hỗ trợ**
  - [ ] Tích hợp chatbot cơ bản
  - [ ] Hỏi đáp nhanh qua giao diện

---

## 📌 Ghi chú bổ sung


---
