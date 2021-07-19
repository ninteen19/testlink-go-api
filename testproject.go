package testlink

type TestProjectRequest struct {
	DevKey          string `json:"devKey" xmlrpc:"devKey"`
	TestProjectName string `json:"testProjectName" xmlrpc:"testprojectname"`
}

func (r *TestProjectRequest) SetDevKey(devKey string) {
	r.DevKey = devKey
}
