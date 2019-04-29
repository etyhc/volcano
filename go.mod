module volcano

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.4
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190426145343-a29dc8fdc734
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190428183149-804c0c7841b5
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190428024724-550556f78a90
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.4.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190425155659-357c62f0e4bb
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
	lemna => ../lemna
)

require (
	github.com/golang/protobuf v1.3.1
	google.golang.org/grpc v1.19.0
	lemna v0.0.0-00010101000000-000000000000
)
