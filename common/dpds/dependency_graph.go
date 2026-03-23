package dpds

import (
	"fmt"
	"sync"
)

type DependencyGraph struct {
	mu           sync.RWMutex
	dependents   map[string][]string
	dependencies map[string][]string
}

func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		dependents:   make(map[string][]string),
		dependencies: make(map[string][]string),
	}
}

func (dg *DependencyGraph) AddDependency(taskID, dependsOnID string) error {
	dg.mu.Lock()
	defer dg.mu.Unlock()

	if taskID == dependsOnID {
		return fmt.Errorf("task cannot depend on itself")
	}

	if dg.wouldCreateCycle(taskID, dependsOnID) {
		return fmt.Errorf("adding dependency would create a cycle")
	}

	dg.dependents[dependsOnID] = append(dg.dependents[dependsOnID], taskID)
	dg.dependencies[taskID] = append(dg.dependencies[taskID], dependsOnID)

	return nil
}

func (dg *DependencyGraph) RemoveDependency(taskID, dependsOnID string) error {
	dg.mu.Lock()
	defer dg.mu.Unlock()

	dg.dependents[dependsOnID] = removeString(dg.dependents[dependsOnID], taskID)
	dg.dependencies[taskID] = removeString(dg.dependencies[taskID], dependsOnID)

	return nil
}

func (dg *DependencyGraph) wouldCreateCycle(taskID, dependsOnID string) bool {
	visited := make(map[string]bool)
	return dg.hasPath(dependsOnID, taskID, visited)
}

func (dg *DependencyGraph) hasPath(from, to string, visited map[string]bool) bool {
	if from == to {
		return true
	}
	if visited[from] {
		return false
	}
	visited[from] = true

	for _, neighbor := range dg.dependents[from] {
		if dg.hasPath(neighbor, to, visited) {
			return true
		}
	}
	return false
}

func (dg *DependencyGraph) GetDependents(taskID string) []string {
	dg.mu.RLock()
	defer dg.mu.RUnlock()
	return dg.dependents[taskID]
}

func (dg *DependencyGraph) GetDependencies(taskID string) []string {
	dg.mu.RLock()
	defer dg.mu.RUnlock()
	return dg.dependencies[taskID]
}

func (dg *DependencyGraph) HasDependencies(taskID string) bool {
	dg.mu.RLock()
	defer dg.mu.RUnlock()
	return len(dg.dependencies[taskID]) > 0
}

func (dg *DependencyGraph) RemoveTask(taskID string) {
	dg.mu.Lock()
	defer dg.mu.Unlock()

	for _, dependent := range dg.dependents[taskID] {
		dg.dependencies[dependent] = removeString(dg.dependencies[dependent], taskID)
	}
	delete(dg.dependents, taskID)
	delete(dg.dependencies, taskID)
}

func (dg *DependencyGraph) UpdateDependencies(taskID string, newDeps []string) {
	dg.mu.Lock()
	defer dg.mu.Unlock()

	for _, dep := range dg.dependencies[taskID] {
		dg.dependents[dep] = removeString(dg.dependents[dep], taskID)
	}
	dg.dependencies[taskID] = newDeps
	for _, dep := range newDeps {
		dg.dependents[dep] = append(dg.dependents[dep], taskID)
	}
}

func removeString(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}