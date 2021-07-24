# [containerd](https://github.com/containerd/containerd) SquashOverlay snapshotter plugin

SquashOverlay snapshotter plugin for containerd. It is based on the overlayfs plugin and change the underlying RO layers storage to squashfs.

This plugin is tested on Linux with Fedora.

## Usage
1. Build the plugin binary: `go build cmd/main.go`

2. Run plugin: `./main /var/run/squashoverlay.sock /tmp/squashoverlay`, where /tmp/squashoverlay is the directory that snapshots will be stored.

3. Condig containerd: add following section to your /etc/containerd/config.toml file

```
[proxy_plugins]
  [proxy_plugins.squashoverlay]
    type = "snapshot"
    address = "/var/run/squashoverlay.sock"
```

4. Start containerd.

5. e.g. `ctr pull --snapshotter=squashoverlay ...` or `CONTAINERD_SNAPSHOTTER=squashoverlay ctr ...`
