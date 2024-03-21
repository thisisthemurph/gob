package api

import (
	"encoding/json"
	"fmt"

	"github.com/thisisthemurph/gob/model"
	"github.com/thisisthemurph/gob/response"
)

// ClusterAPI is an API for interacting with Databricks clusters.
type ClusterAPI struct {
	Host       string
	BaseURL    string
	Databricks DatabricksAPI
}

func NewClusterAPI(databricksHostURL, token string) ClusterAPI {
	clustersURL := fmt.Sprintf("%s/api/2.0/clusters", databricksHostURL)
	return ClusterAPI{
		Host:       databricksHostURL,
		BaseURL:    clustersURL,
		Databricks: NewDatabricksAPI(clustersURL, token),
	}
}

// List returns information about all clusters.
func (api ClusterAPI) List() ([]model.Cluster, error) {
	var rsp response.ClusterListResponse

	b, err := api.Databricks.Get("/list")
	if err != nil {
		return rsp.Clusters, err
	}

	err = json.Unmarshal(b, &rsp)
	return rsp.Clusters, err
}

// Get retrieves the information for a cluster given its identifier.
func (api ClusterAPI) Get(clusterID string) (model.Cluster, error) {
	var cluster model.Cluster

	endpoint := fmt.Sprintf("/get?cluster_id=%s", clusterID)
	b, err := api.Databricks.Get(endpoint)
	if err != nil {
		return cluster, err
	}

	err = json.Unmarshal(b, &cluster)
	return cluster, err
}
