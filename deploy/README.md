

## Dependencies
- Assuming you have `aws-cli` installed and get credentials saved.
```sh
brew install awscli
aws configure
```
- Assume you have public key present at location `~/.ssh/id_rsa.pub`.
```sh
ssh-keygen -t rsa
```


## How to use
1. Deploy AWS stack.
```sh
sh deploy.sh
```
2. URL should show up in the results.
3. That URL can be used as PoC.

## Bring down the stack

```sh
aws cloudformation delete-stack --stack-name jp-search-combinations
```