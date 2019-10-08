package gen

//go:generate go run msgraph-download.go -pretty -baseURL https://graph.microsoft.com/v1.0 -out metadata-v1.0.xml
//go:generate go run msgraph-download.go -pretty -baseURL https://graph.microsoft.com/beta -out metadata-beta.xml
//go:generate go run msgraph-generate.go -baseURL https://graph.microsoft.com/v1.0 -in metadata-v1.0.xml -out ../v1.0/msgraph.go
//go:generate go run msgraph-generate.go -baseURL https://graph.microsoft.com/beta -in metadata-beta.xml -out ../beta/msgraph.go
