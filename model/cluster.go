package model

import "fmt"

type Cluster struct {
	ID                        string          `json:"cluster_id"`
	Name                      string          `json:"cluster_name"`
	SparkVersion              string          `json:"spark_version"`
	CreatorUserName           string          `json:"creator_user_name"`
	SparkContextID            int64           `json:"spark_context_id,omitempty"`
	DriverHealthy             bool            `json:"driver_healthy"`
	AzureAttributes           AzureAttributes `json:"azure_attributes"`
	NodeTypeID                string          `json:"node_type_id"`
	DriverNodeTypeID          string          `json:"driver_node_type_id"`
	AutoterminationMinutes    int             `json:"autotermination_minutes"`
	EnableElasticDisk         bool            `json:"enable_elastic_disk"`
	Source                    string          `json:"cluster_source"`
	EnableLocalDiskEncryption bool            `json:"enable_local_disk_encryption"`
}

func (c Cluster) String() string {
	return fmt.Sprintf("Cluster: %s", c.ID)
}

type AzureAttributes struct {
	FirstOnDemand   int     `json:"first_on_demand"`
	Availability    string  `json:"availability"`
	SpotBidMaxPrice float64 `json:"spot_bid_max_price"`
}
