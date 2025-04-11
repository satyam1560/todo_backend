package auth

// import (
// 	"context"

// 	"github.com/google/uuid"
// 	db "github.com/satyam1560/todo_backend/internal/database/generated"
// )

// type UserAuth struct {
//     Q *db.Queries
// }

// // You can call this from the handler
// func (u *UserAuth) GetOrCreateUserByPhone(ctx context.Context, phone string) (db.User, error) {
//     user, err := u.Q.GetUserByPhone(ctx, phone)
//     if err == nil {
//         return user, nil
//     }

//     // If user not found, create new
//     input := db.CreateUserParams{
//         ID:    uuid.New(),
//         Phone: phone,
//     }

//     return u.Q.CreateUser(ctx, input)
// }
