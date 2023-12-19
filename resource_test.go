package main

import (
	"sync"
	"testing"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/require"
)

func TestMergeResources(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {

		// create an app registration
		name := "test"
		adAppReg, err := azuread.NewApplicationRegistration(ctx, name, &azuread.ApplicationRegistrationArgs{
			DisplayName: pulumi.String(name),
		},
		)
		if err != nil {
			return err
		}
		ctx.Export("adAppReg", adAppReg)

		// install dex helm chart
		dexHelm, err := helmv3.NewChart(ctx, "dex", helmv3.ChartArgs{
			Chart: pulumi.String("dex"),
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String("https://charts.dexidp.io"),
			},
			Values: pulumi.Map{
				"config": pulumi.Map{
					"connectors": pulumi.Array{
						pulumi.Map{
							"config": pulumi.Map{
								"clientID": adAppReg.ClientId, // this line triggers the error
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		var wg sync.WaitGroup
		wg.Add(1)

		pulumi.All(dexHelm.Ready).ApplyT(func(args []interface{}) pulumi.ArrayOutput {
			wg.Done()
			return pulumi.ArrayOutput{}
		})

		wg.Wait()

		return nil
	},
		WithMocks("demo-project", "demo-stack", Mocks(0)),
	)
	require.NoError(t, err)
}

type Mocks int

func (Mocks) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	return args.Name + "_id", args.Inputs, nil
}

func (Mocks) Call(args pulumi.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}

func WithMocks(project, stack string, mocks pulumi.MockResourceMonitor) pulumi.RunOption {
	return func(info *pulumi.RunInfo) {
		info.Project, info.Stack, info.Mocks = project, stack, mocks
	}
}
