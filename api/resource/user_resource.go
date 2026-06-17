package resource

import "BaseProject/models"

type RoleResource struct {
    ID    uint   `json:"id"`
    Title string `json:"title"`
}

type UserResource struct {
    ID          uint           `json:"id"`
    FirstName   string         `json:"first_name"`
    LastName    string         `json:"last_name"`
    Phone       string         `json:"phone_number"`
    Email       string         `json:"email"`
    Roles       []RoleResource `json:"roles"`
}

func NewUserResource(user models.User) UserResource {
    roles := []RoleResource{}
    for _, userRole := range user.UserRoles {
        roles = append(roles, RoleResource{
            ID:    userRole.Role.ID,
            Title: userRole.Role.Title,
        })
    }

    return UserResource{
        ID:        user.ID,
        FirstName: user.First_name,
        LastName:  user.Last_name,
        Phone:     user.Phone_number,
        Email:     user.Email,
        Roles:     roles,
    }
}

func NewUserResourceCollection(users []models.User) []UserResource {
    result := []UserResource{}
    for _, user := range users {
        result = append(result, NewUserResource(user))
    }
    return result
}