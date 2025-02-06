<h1 align="center">aw-sync-suite-plugins</h1>
<p align="center">

   <a href="https://github.com/phrp720/aw-sync-suite-plugins/actions/workflows/tests.yaml?query=branch%3Amain">
    <img title="Tests" src="https://github.com/phrp720/aw-sync-suite-plugins/actions/workflows/tests.yaml/badge.svg?branch=main" alt="tests"/>
  </a>

  <a href="https://github.com/phrp720/aw-sync-suite-plugins/releases">
    <img title="Latest release" src="https://img.shields.io/github/v/release/phrp720/aw-sync-suite-plugins" alt="Latest release">
  </a>
</p>

<p align="center">
 <br>
  📖 For detailed documentation, visit our <a href="https://github.com/phrp720/aw-sync-suite-plugins/wiki">Plugin</a> and <a href="https://github.com/phrp720/aw-sync-suite/wiki">Agent</a> GitHub Wiki.
</p>

## 🔍 About
This is a Repository that hosts the plugins of aw-sync-suite [Agent](https://github.com/phrp720/aw-sync-suite/blob/master/aw-sync-agent/README.md).For versions of aw-sync-suite  >= 0.1.3

Plugins of aw-sync-suite Agent are used to extend the functionality of the agent. The plugins are written in Go and are executed by the agent during the aggregation stage.


## 🔌 Available Plugins


| Plugin    | Description                       | Has Config | Config File              | Documentation                                                       |
|-----------|-----------------------------------|------------|--------------------------|---------------------------------------------------------------------|
| `filters` | Filters the data of ActivityWatch | ✅          | `aw-plugin-filters.yaml` | [📄](https://github.com/phrp720/aw-sync-suite-plugins/wiki/Filters) |
| `scripts` | Runs third party Scripts          | ✅          | `aw-plugin-scripts.yaml` | [📄](https://github.com/phrp720/aw-sync-suite-plugins/wiki/Scripts) |

## ⚙️ How It Works

For documentation about the plugin workflow in `aw-sync-agent`, please refer [here](https://github.com/phrp720/aw-sync-suite-plugins/wiki/%E2%9A%99%EF%B8%8F-Plugin-Workflow).

## 🛠️ How to create a plugin

For instructions on creating a plugin, please refer [here](https://github.com/phrp720/aw-sync-suite-plugins/wiki/%F0%9F%93%9D-How-to-Create-a-Plugin).

## 🔗 How to integrate a plugin

For instructions on integrating a plugin, please refer [here](https://github.com/phrp720/aw-sync-suite-plugins/wiki/%F0%9F%9B%A0%EF%B8%8F--How-to-Integrate-a-Plugin).

## 📝 License

This project is licensed under the **MIT license**.

See [LICENSE](https://github.com/phrp720/aw-sync-suite/blob/master/LICENSE) for more information.
