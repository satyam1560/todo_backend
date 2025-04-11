package auth

import (
    "context"
    firebase "firebase.google.com/go/v4"
    "google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebase() error {
    opt := option.WithCredentialsFile("firebase_adminSDK.json")
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        return err
    }
    FirebaseApp = app
    return nil
}
