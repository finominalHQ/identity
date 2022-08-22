# Finominal
## Identity

### Introduction
Identity handles agent(users and services) authentication and authorization.

### Features
- RBAC
- Host all information necessary to authenticate and authorize an agent
- An agent identity is sniffed from requests (JWT token in header)
- Model after K8s default RBAC

## Reference
It's based on [Go Buffalo Boilerplate](https://github.com/chsqur/boilerplate-go)