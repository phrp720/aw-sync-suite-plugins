<h1 align="center">aw-sync-agent-plugins</h1>

## 🔍 About
This is a Repository that hosts the plugins of aw-sync-suite [Agent](https://github.com/phrp720/aw-sync-suite/blob/master/aw-sync-agent/README.md).

Plugins of aw-sync-suite Agent are used to extend the functionality of the agent. The plugins are written in Go and are executed by the agent during the aggregation stage.

## ⚙️ Plugins


| Plugin    | Description                       | Has Config | Config File            |
|-----------|-----------------------------------|------------|------------------------|
| `filters` | Filters the data of ActivityWatch | ✅          | `aw-plugin-filtes.yml` |



## 🛠️ How to create a plugin

### Core Plugin Structure

To write a plugin, you need to create a Go folder in the `plugins` directory.
Inside this  folder you should contain the plugin implementation idea which will implements the `Plugin` interface as a core of the plugin.



| Method            | Signature                                                                     |
|-------------------|-------------------------------------------------------------------------------|
| `Initialize`      | `Initialize()`                                                                |
| `Execute`         | `Execute(watcher string, events Events, userID string, includeHostName bool)` |
| `ReplicateConfig` | `ReplicateConfig(path string)`                                                |
| `RawName`         | `RawName() string`                                                            |
| `Name`            | `Name() string`                                                               |


## 📝 License

This project is licensed under the **MIT license**.

See [LICENSE](https://github.com/phrp720/aw-sync-suite/blob/master/LICENSE) for more information.