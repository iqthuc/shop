# 🛡️ Authentication - Specification

Thiết kế hệ thống đăng nhập/đăng xuất hiện đại, an toàn và hỗ trợ mở rộng.

---

## ✅ MVP Checklist

- [ ] Đăng nhập bằng email/password
- [ ] Đăng xuất
- [ ] Access Token & Refresh Token
- [ ] Duy trì phiên đăng nhập (session)
- [ ] Hỗ trợ đăng nhập đa thiết bị
- [ ] Đăng nhập bằng tài khoản Google

---

## 📊 Database Schema

### users

| Field | Type | Note |
|-------|------|------|
| id | UUID (PK) | |
| email | TEXT | Unique |
| password_hash | TEXT | Null nếu dùng Google |
| full_name | TEXT | |
| avatar_url | TEXT | |
| created_at | TIMESTAMP | |
| updated_at | TIMESTAMP | |

### sessions

| Field | Type | Note |
|-------|------|------|
| id | UUID (PK) | |
| user_id | UUID (FK) | |
| refresh_token | TEXT | Unique |
| user_agent | TEXT | |
| ip_address | TEXT | |
| expires_at | TIMESTAMP | |
| created_at | TIMESTAMP | |

---

## 🔐 Token Design

- **Access Token**
  - JWT
  - Expire: 15 phút
  - Payload: `user_id`, `session_id`, `exp`, `iat`

- **Refresh Token**
  - Random string hoặc JWT đơn giản
  - Expire: 7–30 ngày
  - Lưu DB (trong `sessions`)

---

## 🔁 Flow

### Đăng nhập (email/password)
1. Client gửi email + password
2. Server xác thực → tạo session
3. Trả về access_token + refresh_token

### Refresh Token
1. Client gửi refresh_token
2. Server xác thực → tạo access_token mới

### Đăng xuất
1. Client gửi refresh_token
2. Server xoá session trong DB

---

## 📱 Đăng nhập đa thiết bị

- Mỗi thiết bị ↔ 1 session
- Quản lý riêng biệt
- Cho phép đăng xuất từng thiết bị

---

## 🔐 Bảo mật

- Brute force → rate limiting
- JWT ngắn hạn
- Refresh Token: lưu ở httpOnly cookie hoặc client nhưng mã hoá
- Rotation token nếu muốn chống reuse

---

## 🌐 Đăng nhập bằng Google (OAuth)

1. Client lấy `id_token` từ Google
2. Gửi `id_token` đến server
3. Server verify, tạo user nếu cần
4. Tạo session → trả access_token + refresh_token

---

## 🚀 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/login` | Đăng nhập |
| POST | `/refresh-token` | Làm mới token |
| POST | `/logout` | Xoá session |
| GET | `/sessions` | Danh sách thiết bị |
| DELETE | `/sessions/{id}` | Đăng xuất 1 thiết bị |
| POST | `/auth/google` | Đăng nhập bằng Google |

---

> Ghi chú:
