# pulumi-azure-app-reg-helm-go-error

Demonstrates a bug in pulumi. See https://github.com/pulumi/pulumi/pull/14949

When running the test in this repo, pulumi crashes with the following error:

```
Running tool: /usr/bin/go test -timeout 30s -run ^TestMergeResources$ github.com/mheers/pulumi-azure-app-reg-helm-go-error

panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0xb9a089]

goroutine 87 [running]:
github.com/pulumi/pulumi/sdk/v3/go/internal.(*OutputState).fulfillValue(0xc000529d50, {0x3eb3c20?, 0x0?, 0x0?}, 0x1, 0x0, {0xc000682440, 0x1, 0x1}, {0x0, ...})
	/home/marcel/go/pkg/mod/github.com/pulumi/pulumi/sdk/v3@v3.96.2/go/internal/types.go:200 +0x109
github.com/pulumi/pulumi/sdk/v3/go/internal.(*OutputState).applyTWithApplier.func1()
	/home/marcel/go/pkg/mod/github.com/pulumi/pulumi/sdk/v3@v3.96.2/go/internal/types.go:634 +0x37b
created by github.com/pulumi/pulumi/sdk/v3/go/internal.(*OutputState).applyTWithApplier in goroutine 10
	/home/marcel/go/pkg/mod/github.com/pulumi/pulumi/sdk/v3@v3.96.2/go/internal/types.go:610 +0x33b
FAIL	github.com/mheers/pulumi-azure-app-reg-helm-go-error	0.023s
FAIL
```
