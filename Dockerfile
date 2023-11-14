# Build stage
FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Bởi vì defaul linux alpine không có curl
# Sử dụng câu lệnh để thêm curl phục vụ câu lệnh tiếp theo
RUN apk add curl

# Dòng lệnh bạn đang thấy trong Dockerfile là để tải về và giải nén [Golang Migrate](https://github.com/golang-migrate/migrate) từ [GitHub Releases](https://github.com/golang-migrate/migrate/releases) vào hình ảnh Docker.
# Dưới đây là giải thích từng phần của lệnh:
#
# - `curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz`: Sử dụng `curl` để tải về tệp `migrate.linux-amd64.tar.gz` từ URL được cung cấp. Tùy thuộc vào hệ điều hành và kiến trúc của bạn, bạn có thể cần thay đổi URL này.
#
# - `| tar xvz`: Dùng pipeline (`|`) để chuyển dữ liệu đầu ra của `curl` thành đầu vào của lệnh `tar`. Lệnh `tar` sau đó giải nén và giải mã tệp `.tar.gz`. Cụ thể:
#  - `x`: Giải nén tệp.
#  - `v`: Hiển thị tiến trình chi tiết.
#  - `z`: Sử dụng `gzip` để giải nén.
#
#Sau khi lệnh này được thực thi, bạn sẽ có thư mục chứa các tệp thực thi của Golang Migrate trong hình ảnh Docker của bạn, và bạn có thể sử dụng chúng trong các bước tiếp theo của Dockerfile.
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]