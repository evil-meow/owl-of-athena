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

### v0.2

- [x] Validate service name since it cannot contain dashes or dots
- [ ] Create a single file per folder in kustomization so we can easily remove things
- [ ] Auto add EVILMEOW_REGISTRY_TOKEN repository secret to Github
- [ ] Wait for namespace to be created after apply to copy secret instead of a fixed number of seconds

### v0.3

- [ ] Creating argo apps that allow to deploy to different environments
- [ ] Clean docker images periodically
- [ ] Clean docker images on repo deletion
- [ ] Notify failed deployments in Slack
- [ ] Notify failed builds in Slack

### v0.4

- [ ] Watch owl.yaml for changes and update the descriptors
- [ ] Watch for merges in main and recommit the version number in github. This should remove the dependency on image-updater

### v1.0

- [ ] Catalog of services in database
- [ ] API to create new services from owlgithub.com
- [ ] Create Sentry entry
- [ ] Push dashboards to grafana
- [ ] Links to relevant pages in external services in owlgithub.com