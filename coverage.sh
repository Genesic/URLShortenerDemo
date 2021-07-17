go test ./... -cover=true --coverprofile ./coverage/outfile
go tool cover -html=./coverage/outfile -o ./coverage/cover.html