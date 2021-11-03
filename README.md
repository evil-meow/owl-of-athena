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
- [ ] Auto create infra repository with base and production if not there already
- [ ] Add this repo's infra to ArgoCD using declarative syntax

### v0.2

- [ ] Auto add EVILMEOW_REGISTRY_TOKEN repository secret to Github

### v0.3

- [ ] Notify failed deployments in Slack
- [ ] Notify failed builds in Slack

### v1.0

- [ ] API to create new services from owlgithub.com
- [ ] Create Sentry entry
- [ ] Push dashboards to grafana
- [ ] Links to relevant pages in external services in owlgithub.com