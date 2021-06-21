# Terraform Provider Velocloud

個人的にTerraformのProviderを学習するために作成したリポジトリです。
リポジトリを公開としていますが、特に意図はなく参照・コピーなどは自由に行ってください。

ライセンスは、MITライセンスとします。

# Terraform Provider Hashicups

Run the following command to build the provider

```shell
go build -o terraform-provider-hashicups
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```
