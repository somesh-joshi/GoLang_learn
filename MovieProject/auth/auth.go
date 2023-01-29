package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/somesh-joshi/MovieProject/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type card struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty"`
	Password string             `josn:"password,omitempty"`
}

var collection = db.Db.Collection("auth")

func insertuser(Card card) (*mongo.InsertOneResult, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(Card.Password), bcrypt.DefaultCost)
	Card.Password = string(hash)
	inserted, err := collection.InsertOne(context.Background(), Card)
	return inserted, err
}

func findByname(Card card) (card, error) {
	var data card
	filter := bson.M{"username": Card.Username}
	err := collection.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		return card{}, err
	} else {
		return data, nil
	}
}

func Logup(w http.ResponseWriter, r *http.Request) {
	var Card card
	_ = json.NewDecoder(r.Body).Decode(&Card)
	if Card.Username == "" || Card.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Missing username and password`))
	} else {
		inserted, err := insertuser(Card)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`Error in add user`))
		} else {
			id := inserted.InsertedID
			token, err := getToken(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`Error in create token, Please login`))
			} else {
				json.NewEncoder(w).Encode(token)
			}
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var Card card
	_ = json.NewDecoder(r.Body).Decode(&Card)
	if Card.Username == "" || Card.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Missing username and password`))
	} else {
		data, err := findByname(Card)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`Invalid username`))
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(Card.Password))
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`Wrong passwords`))
			} else {
				id := data.ID
				token, _ := getToken(id)
				json.NewEncoder(w).Encode(token)
			}
		}
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Missing Authorization Header")
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["id"].(string)
		r.Header.Set("id", name)
		next.ServeHTTP(w, r)
	})
}
