package producthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	kafebar "github.com/kafebar/kafebar/api/kafebar"
)

type handler struct {
	ps kafebar.ProductService
	es kafebar.EventsService
}

func New(ps kafebar.ProductService, es kafebar.EventsService) http.Handler {
	mux := http.NewServeMux()

	h := &handler{ps, es}

	mux.HandleFunc("GET /api/products", h.getProducts)
	mux.HandleFunc("POST /api/products", h.createProduct)
	mux.HandleFunc("PUT /api/products", h.updateProduct)
	mux.HandleFunc("DELETE /api/products/{productId}", h.deleteProduct)

	return mux
}

func (h *handler) getProducts(w http.ResponseWriter, req *http.Request) {
	products, err := h.ps.GetProducts(req.Context())
	if err != nil {

		err := fmt.Errorf("cannot get products: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (h *handler) createProduct(w http.ResponseWriter, req *http.Request) {
	var product kafebar.Product

	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {

		err := fmt.Errorf("invalid product body: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProduct, err := h.ps.CreateProduct(req.Context(), product)
	if err != nil {
		err := fmt.Errorf("cannot create product: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.es.Broadcast(req.Context(), kafebar.Event{
		Type: kafebar.EventTypeProductCreated,
		Data: createdProduct,
	})

	json.NewEncoder(w).Encode(createdProduct)
}

func (h *handler) updateProduct(w http.ResponseWriter, req *http.Request) {
	var product kafebar.Product

	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		err := fmt.Errorf("invalid product body: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProduct, err := h.ps.UpdateProduct(req.Context(), product)
	if err != nil {
		err := fmt.Errorf("cannot update product: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.es.Broadcast(req.Context(), kafebar.Event{
		Type: kafebar.EventTypeProductUpdated,
		Data: createdProduct,
	})

	json.NewEncoder(w).Encode(createdProduct)
}

func (h *handler) deleteProduct(w http.ResponseWriter, req *http.Request) {
	productId, err := strconv.Atoi(req.PathValue("productId"))
	if err != nil {
		err := fmt.Errorf("invalid product id: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ps.DeleteProduct(req.Context(), productId)
	if err != nil {
		err := fmt.Errorf("cannot delete product: %w", err)
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.es.Broadcast(req.Context(), kafebar.Event{
		Type: kafebar.EventTypeOrderDeleted,
		Data: productId,
	})
}
