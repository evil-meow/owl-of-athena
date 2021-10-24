# Owl of athena

![build status](https://github.com/evil-meow/owl-of-athena/actions/workflows/publish-image.yml/badge.svg)

The bot that helps me to deploy pet projects

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
- [ ] Add this repo's infra to ArgoCD
- [ ] Auto create infra repository with base and production if not there already
- [ ] Add target repository to ArgoCD
- [ ] Auto add EVILMEOW_REGISTRY_TOKEN repository secret
- [ ] Notify deployments in Slack?
- [ ] A way to apply base to new environments
- [ ] Create Sentry entry
- [ ] Push dashboards to grafana

### v1.0

- [ ] API to create new services from owlgithub.com
- [ ] Links to relevant pages in external services in owlgithub.com