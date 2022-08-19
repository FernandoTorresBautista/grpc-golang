package main

import (
	"context"
	"fmt"
	"grpc-course/products/productpb"

	"io"
	"log"

	"google.golang.org/grpc"
)

// Remember we call all functions from here but it could be different
// clients, each one with its service implemented
func main() {
	fmt.Println("Go client is running")
	// create the grpc client
	opts := grpc.WithInsecure() // no SSL certificates
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}
	// @doc: The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding
	// function returns.
	defer cc.Close() //

	c := productpb.NewProductServiceClient(cc)

	// Creating Product
	fmt.Println("------Creating Product------")
	product := &productpb.Product{
		Name:  "Smartphone AA",
		Price: 20300.50,
	}
	createdProduct, err := c.CreateProduct(context.Background(), &productpb.CreateProductRequest{
		Product: product,
	})
	if err != nil {
		log.Fatalf("Failed to create product %v", err)
	}
	fmt.Printf("Product created %v \n", createdProduct)
	// We going to use these product Id through all calls
	productID := createdProduct.GetProduct().GetId()
	// productID := "62fc231b60ef784f17008d17" // first product created

	//Reading the product
	fmt.Println("-----Getting the Product-----")
	_, err2 := c.GetProduct(context.Background(), &productpb.GetProductRequest{ProductId: productID})
	if err2 != nil {
		fmt.Printf("Error happened while getting: %v \n", err2)
	}
	getProductReq := &productpb.GetProductRequest{ProductId: productID}
	getProductRes, getProductErr := c.GetProduct(context.Background(), getProductReq)
	if getProductErr != nil {
		fmt.Printf("Error happened while getting: %v \n", getProductErr)
	}
	fmt.Printf("Product gotten: %v \n", getProductRes)

	// update Product
	fmt.Println("-----Updating  Product-----")
	// It is a new struct productpb with the id of
	// the already created product and the new name and price
	newProduct := &productpb.Product{
		Id:    productID,
		Name:  "Name changed Iphone XX",
		Price: 40500,
	}
	updateRes, updateErr := c.UpdateProduct(context.Background(), &productpb.UpdateProductRequest{Product: newProduct})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v \n", updateErr)
	}
	fmt.Printf("Product was updated: %v\n", updateRes)

	// Delete Product
	fmt.Println("-----Deleting Product-----")
	deleteRes, deleteErr := c.DeleteProduct(context.Background(), &productpb.DeleteProductRequest{ProductId: productID})
	if deleteErr != nil {
		fmt.Printf("Error while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Product was deleted: %v \n", deleteRes.GetProductId())

	// List Products
	fmt.Println("-----Listing Product-----")
	stream, err := c.ListProduct(context.Background(), &productpb.ListProductRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving products: %v", err)
		}
		fmt.Println(res.GetProduct())
	}
}
