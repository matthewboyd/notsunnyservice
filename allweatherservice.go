package notsunnyservice

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sony/gobreaker"
	"log"
	"math/rand"
	"time"
)

type Handler struct {
	Logger         log.Logger
	Db             pgxpool.Pool
	Redis          redis.Client
	CircuitBreaker *gobreaker.CircuitBreaker
}

type Activities struct {
	Name     string
	Postcode string
	Sunny    bool
}

func (h *Handler) GetNotSunnyActivities() string {
	var a Activities
	var newActivityList []Activities
	var ctx context.Context
	notSunnyActivitiesQuery := "SELECT * FROM activities where sunny = $1"
	rows, err := h.Db.Query(ctx, notSunnyActivitiesQuery, false)
	if err != nil {
		log.Fatalln("An error occurred", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&a.Name, &a.Postcode, &a.Sunny)
		if err != nil {
			log.Fatalln("Error in scanning db rows", err)
		}
		newActivityList = append(newActivityList, a)
	}
	choosenActivity, _ := h.retrieveActivity(newActivityList)
	return fmt.Sprintf("%s %s", choosenActivity.Name, choosenActivity.Postcode)
}

func (h *Handler) retrieveActivity(newActivityList []Activities) (Activities, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomNumber := r1.Intn(len(newActivityList))
	choosenActivity := newActivityList[randomNumber]
	return choosenActivity, nil
}
