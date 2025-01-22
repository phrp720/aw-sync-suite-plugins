# aw-sync-agent-plugins

This is a Repository that hosts the plugins of aw-sync-suite Agent.


## Plugins

Plugins of aw-sync-suite Agent are used to extend the functionality of the agent. The plugins are written in Go and are executed by the agent in the aggregation stage.


| Plugin    | Description                         | Has Config | Config File            |
|-----------|-------------------------------------|------------|------------------------|
| `filters` | `Filters the data of ActivityWatch` | ✅          | `aw-plugin-filtes.yml` |



## How to write a plugin

### Core Plugin Structure

To write a plugin, you need to create a Go folder in the `plugins` directory.
Inside this  folder you should contain the plugin implementation idea which will implements the `Plugin` interface as a core of the plugin.

```go

| Method            | Signature                                                                     |
|-------------------|-------------------------------------------------------------------------------|
| `Initialize`      | `Initialize()`                                                                |
| `Execute`         | `Execute(watcher string, events Events, userID string, includeHostName bool)` |
| `ReplicateConfig` | `ReplicateConfig(path string)`                                                |
| `RawName`         | `RawName() string`                                                            |
| `Name`            | `Name() string`                                                               |
