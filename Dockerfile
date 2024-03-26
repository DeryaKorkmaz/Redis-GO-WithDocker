# Golang için resmi Docker imajını kullan
FROM golang:latest 

# Çalışma dizinini /app olarak belirle
WORKDIR /app

# Golang uygulamanızı kopyala
COPY . .

# Uygulamayı derle
RUN go build -o main .

# Çalıştırılabilir dosyayı çalıştır
CMD ["./main"]
