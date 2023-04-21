package server

import (
	"log"
	"time"
)

type Secret struct {
	Key          string `json:"key"`
	Content      string `json:"content"`
	CreatedAt    time.Time
	Lifespan     time.Duration
	DeleteOnRead bool
}

type SecretController struct {
	Secrets        map[string]Secret
	SecretLifespan time.Duration
}

func newSecretController(lifespan time.Duration) *SecretController {
	return &SecretController{make(map[string]Secret), lifespan}
}

func (sm *SecretController) isSecretExpired(key string) bool {
	secret, ok := sm.Secrets[key]
	if ok {
		isExpired := time.Since(secret.CreatedAt) > secret.Lifespan
		if isExpired {
			log.Printf("EXPIRE secret: %s", key)
			delete(sm.Secrets, key)
		}
		return isExpired
	}
	return false
}

func (sm *SecretController) GetSecret(key string) string {
	secret, ok := sm.Secrets[key]
	log.Printf("GET secret: %s", key)
	resp := ""
	if ok {
		if !sm.isSecretExpired(key) {
			resp = secret.Content
		} else {
			sm.DeleteSecret(key)
		}
		if secret.DeleteOnRead {
			sm.DeleteSecret(key)
		}
	}
	return resp
}

func (sm *SecretController) UpdateSecret(key string, content string) string {
	log.Printf("UPDATE secret: %s", key)
	sm.Secrets[key] = Secret{key, content, time.Now(), sm.SecretLifespan, true}
	return content
}

func (sm *SecretController) DeleteSecret(key string) string {
	log.Printf("DELETE secret: %s", key)
	_, ok := sm.Secrets[key]
	if ok {
		delete(sm.Secrets, key)
		return ""
	}
	return ""
}
