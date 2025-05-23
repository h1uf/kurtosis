package v3

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                           DO NOT CHANGE THIS FILE!
  If you change this file, it will break config for users who have instantiated an
           overrides file with this version of config overrides!
    Instead, to make changes, you will need to add a new version of the config
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

type KurtosisClusterConfigV3 struct {
	Type *string `yaml:"type,omitempty"`
	// If we ever get another type of cluster that has configuration, this will need to be polymorphically deserialized
	Config         *KubernetesClusterConfigV3 `yaml:"config,omitempty"`
	LogsAggregator *LogsAggregatorConfigV3    `yaml:"logs-aggregator,omitempty"`
}
