package infrastructure

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/tidwall/gjson"
	"gopkg.in/go-playground/assert.v1"
)

func TestLandingZone(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := Factory(app, "dev", LandingZone)

	// THEN
	bytess, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytess)
	displayName := template.Get("Resources.#").String()
	t.Log("------------")
	t.Log(template.Get("Resources").String())
	b := template.Get("Resources").String()
	var out bytes.Buffer
	json.Indent(&out, []byte(b), "=", "\t")
	out.WriteTo(os.Stdout)
	t.Log()
	assert.Equal(t, "MyCoolTopic", displayName)
}
