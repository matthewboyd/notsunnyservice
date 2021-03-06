package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	pb "github.com/matthewboyd/notsunnyservice/pb"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

type Handler struct {
	Logger log.Logger
	Db     pgxpool.Pool
	Redis  redis.Client
}

type server struct {
	pb.UnimplementedNotSunnyActivitiesServer
	Handler
}

type Activities struct {
	Name     string
	Postcode string
	Sunny    bool
}

func (h *server) GetAllWeatherActivities(ctx context.Context, in *pb.NotSunnyActivitiesParams) (*pb.ActivityResponse, error) {
	log.Println("In the GetAllWeatherService")
	log.Printf("the server: %v and the handler of server: %v and db direct from server: %v", h, h.Handler, h.Db)
	//var a Activities
	//var newActivityList []Activities
	//notSunnyActivitiesQuery := "SELECT * FROM activities where sunny = $1"
	//rows, err := h.Handler.Db.Query(ctx, notSunnyActivitiesQuery, false)
	//log.Printf("rows are: %v", rows)
	//if err != nil {
	//	log.Fatalln("An error occurred", err)
	//}
	//defer rows.Close()
	//log.Println("just before db")
	//for rows.Next() {
	//	err = rows.Scan(&a.Name, &a.Postcode, &a.Sunny)
	//	if err != nil {
	//		log.Fatalln("Error in scanning db rows", err)
	//	}
	//	newActivityList = append(newActivityList, a)
	//}
	//choosenActivity, _ := h.retrieveActivity(newActivityList)
	//return &pb.ActivityResponse{Allweatheractivity: fmt.Sprintf("%s %s", choosenActivity.Name, choosenActivity.Postcode)}, nil
	return &pb.ActivityResponse{Allweatheractivity: "testing"}, nil
}

func (h *server) retrieveActivity(newActivityList []Activities) (Activities, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomNumber := r1.Intn(len(newActivityList))
	choosenActivity := newActivityList[randomNumber]
	return choosenActivity, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterNotSunnyActivitiesServer(s, &server{})

	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
