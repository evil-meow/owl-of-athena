# Owl of athena

![build status](https://github.com/evil-meow/owl-of-athena/actions/workflows/publish-image.yml/badge.svg)

The bot that helps me to deploy pet projects

## Usage

Place a yml file at the root of the project and ping the bot with

```
/add_service repo-name
```

The yml needs to have the following format

```
version: 1
images:
  - name: image-name
    url: <url-to-expose> (optional)
```

## Deployment

The deployment needs to provide the following variables

```
SLACK_AUTH_TOKEN
SLACK_CHANNEL_ID
SLACK_APP_TOKEN
GITHUB_TOKEN
```

## Roadmap

### v0.1

- [x] Command to create a new service
- [x] Add CI
- [x] Auto create infra repository with base and production if not there already
- [x] Add this repo's infra to ArgoCD using declarative syntax
- [ ] Create `evilmeow-registry-secret` to download images in new namespaces
- [ ] Auto add EVILMEOW_REGISTRY_TOKEN repository secret to Github

### v0.2

- [ ] Validate service name since it cannot contain dashes or dots
- [ ] Creating argo apps that allow to deploy to different environments
- [ ] Create a single file per folder in kustomization so we can easily remove things

### v0.3

- [ ] Notify failed deployments in Slack
- [ ] Notify failed builds in Slack

### v0.4

- [ ] Watch owl.yaml for changes and update the descriptors

### v1.0

- [ ] Catalog of services in database
- [ ] API to create new services from owlgithub.com
- [ ] Create Sentry entry
- [ ] Push dashboards to grafana
- [ ] Links to relevant pages in external services in owlgithub.com