package brokerapi

import (
	"net/http"
	"strings"
	"testing"

	"github.com/arschles/assert"
	"github.com/arschles/testsrv"
	"github.com/deis/steward/web"
)

func provisionReqBody() string {
	return `{
		"organization_guid":"abc",
		"plan_id":"def",
		"service_id":"ghi",
		"space_guid":"jkl"
	}`
}

func TestProvisionUnauthorized(t *testing.T) {
	feAuth := &web.BasicAuth{Username: "testFEUser", Password: "testFEPass"}
	hdl := Handler(nil, nil, feAuth, nil)
	srv := testsrv.StartServer(hdl)
	defer srv.Close()

	// no basic auth
	req, err := http.NewRequest("PUT", srv.URLStr()+"/v2/service_instances/abcd", strings.NewReader(provisionReqBody()))
	assert.NoErr(t, err)
	res, err := http.DefaultClient.Do(req)
	assert.NoErr(t, err)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "response status code")

	// incorrect basic auth
	req, err = http.NewRequest("PUT", srv.URLStr()+"/v2/service_instances/abcd", strings.NewReader(provisionReqBody()))
	assert.NoErr(t, err)
	req.SetBasicAuth("nothing", "nada")
	res, err = http.DefaultClient.Do(req)
	assert.NoErr(t, err)
	assert.Equal(t, res.StatusCode, http.StatusUnauthorized, "response status code")
}