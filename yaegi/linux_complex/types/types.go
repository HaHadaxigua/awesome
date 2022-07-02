package types

type Os struct {
	Cloud    string
	Platform string
}

var AwsOs = Os{
	Cloud:    "Aws",
	Platform: "Windows",
}
