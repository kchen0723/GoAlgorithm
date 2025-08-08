package main

//https://www.cnblogs.com/davis12/p/18111154

const PORT = ":50052"

// func main() {
// 	lis, err := net.Listen("tcp", PORT)
// 	if err != nil {
// 		panic(err)
// 	}

// 	srv := grpc.NewServer()
// 	log.Println("now registing")
// 	RegisterMathServiceServer(srv, &MathService{})
// 	// RegisterAuthServiceServer(srv, &AuthService{})

// 	log.Println("now starting")
// 	err = srv.Serve(lis)
// 	if err != nil {
// 		panic(err)
// 	}

// }
