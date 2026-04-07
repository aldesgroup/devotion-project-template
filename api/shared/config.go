// ------------------------------------------------------------------------------------------------
// configuring this app
// ------------------------------------------------------------------------------------------------
package shared

import g "github.com/aldesgroup/goald"

// registering the config object
func init() {
	g.RegisterConfig(&renameThisAppConfig{
		IBaseConfig: g.NewBaseConfig(),
		CustomConf:  RenameThisCustomPart{},
	})
}

// defining the global config object, that should capture all the YAML file content
type renameThisAppConfig struct {
	g.IBaseConfig `json:"base"`        // do not change the JSON name!
	CustomConf    RenameThisCustomPart `json:"custom"` // do not change the JSON name!
}

// making this global config object legit by implementing this required method
func (thisConf *renameThisAppConfig) CustomConfig() g.ICustomConfig {
	return thisConf.CustomConf
}

// defining the custom part of the config, that is specific to this app
type RenameThisCustomPart struct {
	CustomConfigItem []string // this can be anything here, depending on the app's needs for config
}

// utility to quickly access the custom part of the config
func Config(appContext g.AppContext) *RenameThisCustomPart {
	return appContext.CustomConfig().(*RenameThisCustomPart)
}
