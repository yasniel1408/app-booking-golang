# router go chi

https://github.com/go-chi/chi

# Install

go get -u github.com/go-chi/chi/v5

# Limpiar las dependencias

go mod tidy

# Correr los test
go test -v

go test -cover

go test -coverprofile=coverage.out && go tool cover -html=coverage.out