package middleware

import (
	"sync"
	"time"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

type userRateLimit struct {
	requests []time.Time
	mutex    sync.Mutex
}

var (
	rateLimitMap = make(map[string]*userRateLimit)
	mapMutex     sync.Mutex
)

func init() {
	go cleanupOldEntries()
}

func RateLimitMiddleware(requestsPerMinute int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("user")
		if !ok {
			internal.HandleError(c, internal.ErrUnauthorized)
			c.Abort()
			return
		}

		userIDStr := userID.(string)

		mapMutex.Lock()
		if rateLimitMap[userIDStr] == nil {
			rateLimitMap[userIDStr] = &userRateLimit{
				requests: make([]time.Time, 0, requestsPerMinute),
			}
		}

		limitTracker := rateLimitMap[userIDStr]
		mapMutex.Unlock()

		limitTracker.mutex.Lock()
		defer limitTracker.mutex.Unlock()

		now := time.Now()
		cutoff := now.Add(-1 * time.Minute)

		writeIdx := 0
		for _, t := range limitTracker.requests {
			if t.After(cutoff) {
				limitTracker.requests[writeIdx] = t
				writeIdx++
			}
		}

		limitTracker.requests = limitTracker.requests[:writeIdx]

		if len(limitTracker.requests) >= requestsPerMinute {
			internal.HandleError(c, internal.ErrTooManyRequests)
			c.Abort()
			return
		}

		limitTracker.requests = append(limitTracker.requests, now)

		c.Next()
	}
}

func cleanupOldEntries() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		cutoff := now.Add(-5 * time.Minute)

		usersToDelete := make([]string, 0)

		mapMutex.Lock()
		for userID, tracker := range rateLimitMap {
			tracker.mutex.Lock()
			if len(tracker.requests) == 0 {
				usersToDelete = append(usersToDelete, userID)
			} else {
				allOld := true
				for _, t := range tracker.requests {
					if t.After(cutoff) {
						allOld = false
						break
					}
				}
				if allOld {
					usersToDelete = append(usersToDelete, userID)
				}
			}
			tracker.mutex.Unlock()
		}

		for _, userID := range usersToDelete {
			delete(rateLimitMap, userID)
		}
		mapMutex.Unlock()
	}
}
