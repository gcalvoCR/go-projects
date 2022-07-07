package controllers

import (
	"encoding/json"
	"fmt"
	"go-3-bookstore/pkg/models"
	"go-3-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MessageResponse struct {
	Message string `json:"message"`
}

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	// newBooks:=models.GetAllBooks()
	// res,_ := json.Marshal(newBooks) //convert into json
	// w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusInternalServerError)
	// w.Write(res)

	resp := MessageResponse{Message: "This is problematic"}
	json.NewEncoder(w).Encode(resp)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookId:= vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}

	bookDetails,_:= models.GetBookById(Id)
	res,_:= json.Marshal(bookDetails) // convert to json

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) //json version 
}

func CreateBook(w http.ResponseWriter, r * http.Request){
	CreateBook := &models.Book{}

	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res,_ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookId:= vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(Id)
	res,_ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars:= mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails,db := models.GetBookById(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res,_ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)



}
