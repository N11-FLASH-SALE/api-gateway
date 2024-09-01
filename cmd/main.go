package main

import (
	api "api/api"
	"api/api/handler"
	"api/casbin"
	"api/config"
	"api/genproto/sale"
	"api/genproto/user"
	"api/logs"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf := config.Load()
	hand := NewHandler()
	router := api.Router(hand)
	log.Printf("server is running...")
	log.Fatal(router.Run(conf.API_ROUTER))
}

func NewHandler() handler.HandlerInterface {
	conf := config.Load()
	connUser, err := grpc.NewClient(conf.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	connSale, err := grpc.NewClient(conf.SALE_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	Product := sale.NewProductClient(connSale)
	Process := sale.NewProcessClient(connSale)
	Wishlis := sale.NewWishlistClient(connSale)
	Feedback := sale.NewFeedbackClient(connSale)
	Bought := sale.NewBoughtClient(connSale)
	User := user.NewUserClient(connUser)
	Notification := user.NewNotificationsClient(connUser)
	Card := user.NewCardsClient(connUser)

	logs := logs.NewLogger()
	en, err := casbin.CasbinEnforcer(logs)
	if err != nil {
		log.Fatal("error in creating casbin enforcer", err)
	}
	return &handler.Handler{
		User:         User,
		Product:      Product,
		Process:      Process,
		Wishlist:     Wishlis,
		Feedback:     Feedback,
		Bought:       Bought,
		Notification: Notification,
		Cards:        Card,
		Log:          logs,
		Enforcer:     en,
	}
}
