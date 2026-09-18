package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/manojselukar/terraform-provider-vaultsecure/vaultsecure"
)

func main() {
	err := tfsdk.Serve(context.Background(), vaultsecure.New, tfsdk.ServeOpts{Name: "vaultsecure"})
	if err != nil {
		return
	}
}
