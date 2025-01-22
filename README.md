# aw-sync-agent-plugins
Repository that hosts the plugins of aw-sync-suite Agent



## Plugins

Plugins of aw-sync-suite Agent are used to extend the functionality of the agent. The plugins are written in Go and are executed by the agent in the aggregation stage.

### Core Plugin Structure

| Method            | Signature                                                                     |
|-------------------|-------------------------------------------------------------------------------|
| `Initialize`      | `Initialize()`                                                                |
| `Execute`         | `Execute(watcher string, events Events, userID string, includeHostName bool)` |
| `ReplicateConfig` | `ReplicateConfig(path string)`                                                |
| `RawName`         | `RawName() string`                                                            |
| `Name`            | `Name() string`                                                               |
