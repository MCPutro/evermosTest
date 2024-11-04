package main

import (
	"context"
	"evermosTest/config"
	"evermosTest/internal/database"
	"evermosTest/internal/entity"
	"evermosTest/internal/handler/http/order"
	"evermosTest/internal/handler/http/product"
	"evermosTest/internal/handler/request"
	"evermosTest/internal/repository"
	"evermosTest/internal/route"
	"evermosTest/internal/service"
	"log"
)

func main() {

	cfg := config.NewConfig()

	db, err := database.NewMySQL(cfg)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.Order{})

	repoManager := repository.NewRepositoryManager(db)
	serviceManager := service.NewServiceManager(repoManager)

	productHandler := product.NewProductHandler(serviceManager.ProductService())
	orderHandler := order.NewOrderHandler(serviceManager.OrderService())

	//sample product
	serviceManager.ProductService().Create(context.Background(), &request.Product{
		Name:  "POP Ice",
		Price: 100,
		Stock: 15,
	})

	app := route.NewRouter(productHandler, orderHandler)

	err = app.Listen(":" + cfg.GetAppPort())
	if err != nil {
		log.Fatalln(err)
	}

	////add product
	//productRequest := request.Product{
	//	Name:  "Es Kelapa",
	//	Price: 200,
	//	Stock: 12,
	//}
	//
	//productService.Create(context.Background(), &productRequest)

	//for i := 0; i < 1000; i++ {
	//	go func(x int) {
	//		checkout := request.Checkout{
	//			ProductId: 2,
	//			Quantity:  3,
	//		}
	//
	//		result, err2 := productService.Checkout(context.Background(), &checkout)
	//		if err2 != nil {
	//			log.Println("ERROR Checkout id :", x, " error :", err)
	//		} else {
	//			log.Println("Berhasil Checkout id :", x, " -> ", result)
	//		}
	//	}(i)
	//}
	//
	//time.Sleep(10 * time.Second)
	//fmt.Println("selesai")
}
