package solution

import (
	"sync"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

var singleInstance *solution
var once sync.Once

type solution struct {
	stack      awscdk.Stack
	properties solutionProperties
}

type solutionProperties struct {
	awscdk.StackProps
	appProperties interface{}
}

func New(scope constructs.Construct, id *string) *solution {
	if singleInstance == nil {
		once.Do(
			func() {
				props := solutionProperties{}
				solution := solution{
					stack:      awscdk.NewStack(scope, id, &props.StackProps),
					properties: props,
				}
				deployInfra(solution)

			})
	}
	return singleInstance
}

func deployInfra(s solution) {

}
