package response

import "github.com/thisisthemurph/gob/model"

type ClusterListResponse struct {
	Clusters []model.Cluster `json:"clusters"`
}
