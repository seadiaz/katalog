package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/walmartdigital/katalog/src/domain"
)

func (s *Server) getResourcesByType(resource domain.Resource) []interface{} {
	resources := s.resourcesRepository.GetAllResources()
	list := arraylist.New()
	for _, r := range resources {
		res := r.(domain.Resource)
		if res.GetType() == resource.GetType() {
			list.Add(r)
		}
	}
	return list.Values()
}

func (s *Server) createService(w http.ResponseWriter, r *http.Request) {
	var service domain.Service
	json.NewDecoder(r.Body).Decode(&service)
	resource := domain.Resource{
		K8sResource: &service,
	}
	s.resourcesRepository.CreateResource(resource)
	json.NewEncoder(w).Encode(service)
}

func (s *Server) updateService(w http.ResponseWriter, r *http.Request) {
	var service domain.Service
	json.NewDecoder(r.Body).Decode(&service)
	resource := domain.Resource{
		K8sResource: &service,
	}
	s.resourcesRepository.UpdateResource(resource)
	json.NewEncoder(w).Encode(service)
}

func (s *Server) deleteService(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	s.resourcesRepository.DeleteResource(id)
	fmt.Fprintf(w, "deleted service id: %s", id)
}

func (s *Server) getAllServices(w http.ResponseWriter, r *http.Request) {
	services := s.getResourcesByType(domain.Resource{K8sResource: &domain.Service{}})
	json.NewEncoder(w).Encode(services)
}

func (s *Server) countServices(w http.ResponseWriter, r *http.Request) {
	services := s.getResourcesByType(domain.Resource{K8sResource: &domain.Service{}})
	json.NewEncoder(w).Encode(struct{ Count int }{len(services)})
}

func (s *Server) createDeployment(w http.ResponseWriter, r *http.Request) {
	var deployment domain.Deployment
	json.NewDecoder(r.Body).Decode(&deployment)
	resource := domain.Resource{K8sResource: &deployment}
	s.resourcesRepository.CreateResource(resource)
	(*s.metrics)["createDeployment"].(*prometheus.CounterVec).WithLabelValues(resource.GetID()).Inc()
	json.NewEncoder(w).Encode(deployment)
}

func (s *Server) updateDeployment(w http.ResponseWriter, r *http.Request) {
	var deployment domain.Deployment
	json.NewDecoder(r.Body).Decode(&deployment)
	resource := domain.Resource{K8sResource: &deployment}
	result, err := s.resourcesRepository.UpdateResource(resource)

	if err != nil {
		glog.Errorf("Error occurred trying to update resource (id: %s)", resource.GetID())
	}

	if result != nil {
		(*s.metrics)["updateDeployment"].(*prometheus.CounterVec).WithLabelValues(resource.GetID()).Inc()
	}
	json.NewEncoder(w).Encode(deployment)
}

func (s *Server) deleteDeployment(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := s.resourcesRepository.DeleteResource(id)

	if err != nil {
		fmt.Fprintf(w, "deleted deployment id: %s", id)
	}
	(*s.metrics)["deleteDeployment"].(*prometheus.CounterVec).WithLabelValues(id).Inc()
}

func (s *Server) getAllDeployments(w http.ResponseWriter, r *http.Request) {
	deployments := s.getResourcesByType(domain.Resource{K8sResource: &domain.Deployment{}})
	json.NewEncoder(w).Encode(deployments)
}

func (s *Server) countDeployments(w http.ResponseWriter, r *http.Request) {
	deployments := s.getResourcesByType(domain.Resource{K8sResource: &domain.Deployment{}})
	json.NewEncoder(w).Encode(struct{ Count int }{len(deployments)})
}
