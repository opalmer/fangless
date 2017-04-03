package fangless

import (
	"sync"
	"time"
	"github.com/spf13/viper"
)

// Snake is used to wrap a Viper struct and 'defang' it
// to make it a bit safer to work with by hiding certain
// functionality.
type Snake struct {
	viper *viper.Viper
	mtx   *sync.RWMutex
}

//
// Fangless provided functions
//

// Lock locks the underlying mutex.
func (s *Snake) Lock() {
	s.mtx.Lock()
}

// Unlock unlocks the underlying mutex.
func (s *Snake) Unlock() {
	s.mtx.Unlock()
}

// Viper returns the underlying viper struct. Before returning
// Lock() is called so be sure you eventually call Unlock().
func (s *Snake) Viper() *viper.Viper {
	s.mtx.Lock()
	return s.viper
}

//
// Read functions
//

// Get returns the provided key from the underlying Viper configuration.
func (s *Snake) Get(key string) interface{} {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.Get(key)
}

// GetString returns the provided key from the underlying Viper configuration.
func (s *Snake) GetString(key string) string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetString(key)
}

// GetBool returns the provided key from the underlying Viper configuration.
func (s *Snake) GetBool(key string) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetBool(key)
}

// GetInt returns the provided key from the underlying Viper configuration.
func (s *Snake) GetInt(key string) int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetInt(key)
}

// GetInt64 returns the provided key from the underlying Viper configuration.
func (s *Snake) GetInt64(key string) int64 {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetInt64(key)
}

// GetFloat64 returns the provided key from the underlying Viper configuration.
func (s *Snake) GetFloat64(key string) float64 {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetFloat64(key)
}

// GetTime returns the provided key from the underlying Viper configuration.
func (s *Snake) GetTime(key string) time.Time {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetTime(key)
}

// GetDuration returns the provided key from the underlying Viper configuration.
func (s *Snake) GetDuration(key string) time.Duration {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetDuration(key)
}

// GetStringSlice returns the provided key from the underlying Viper configuration.
func (s *Snake) GetStringSlice(key string) []string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetStringSlice(key)
}

// GetStringMap returns the provided key from the underlying Viper configuration.
func (s *Snake) GetStringMap(key string) map[string]interface{} {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetStringMap(key)
}

// GetStringMapString returns the provided key from the underlying Viper configuration.
func (s *Snake) GetStringMapString(key string) map[string]string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetStringMapString(key)
}

// GetStringMapStringSlice returns the provided key from the underlying Viper configuration.
func (s *Snake) GetStringMapStringSlice(key string) map[string][]string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetStringMapStringSlice(key)
}

// GetSizeInBytes returns the provided key from the underlying Viper configuration.
func (s *Snake) GetSizeInBytes(key string) uint {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.GetSizeInBytes(key)
}

// IsSet returns true if the provided key is set.
func (s *Snake) IsSet(key string) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.viper.IsSet(key)
}

//
// Write functions
//

// SetDefault calls SetDefault on the underlying Viper configuration
func (s *Snake) SetDefault(key string, value interface{}) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.viper.SetDefault(key, value)
}

// Set calls SetDefault on the underlying Viper configuration
func (s *Snake) Set(key string, value interface{}) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.viper.Set(key, value)
}

// New takes a Viper struct and returns a *Snake struct.
func New(v *viper.Viper) *Snake {
	return &Snake{viper: v, mtx: &sync.RWMutex{}}
}
