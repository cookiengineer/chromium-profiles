
# Chromium Profiles

Chromium Profile CLI management tool to make Linux users be able to use Corporate Enterprise
environments where Microslop's O365 decides that blocking Linux users will increase security.

### Features and Opinions

- Chromium profile management CLI tool
- Allows multiple profiles in `~/Work/Sandboxes/<name>` for separated environments
- Bundles a Farble User Agent extension
- Bundles a Farble Clipboard extension


### Usage

First setup is a little annoying, because Chromium's `--load-extension` parameter doesn't
work reliably for various reasons. The `chromium-profiles create <name>` command will do
the following:

- Create a new profile in `~/Work/Sandboxes/<name>`
- Copy the browser extensions to `~/Work/Sandboxes/<name>/chromium-extensions`
- Launch `chromium` with the profile folder parameters
- Show instructions how to install the Browser Extensions via `chrome://extensions`

```bash
# List available Profiles
chromium-profiles list;

# Create a Profile in ~/Work/Sandboxes/enterprise-company-name;
chromium-profiles create enterprise-company-name;

#
# Follow CLI instructions to setup Farble Browser Extensions
#
```

After the Chromium Profile has been setup, it can be launched again in an easier manner:

```bash
chromium-profiles launch enterprise-company-name;
```


### License

WTFPL

