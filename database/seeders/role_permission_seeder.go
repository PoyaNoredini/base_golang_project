package seeders

import (
    "BaseProject/models"
    "log"
    "gorm.io/gorm"
)

type RolePermissionSeeder struct {
    DB *gorm.DB
}

func (s *RolePermissionSeeder) Run() {
    roles := []struct {
        Title       string
        Description string
        Permissions []string
    }{
        {
            Title:       "Administrator",
            Description: "System Administrator",
            Permissions: []string{
             
                "create-role",
                "view-role",
                "update-role",
                "delete-role",

                "create-user",
                "view-user",
                "update-user",
                "delete-user",
                "deactive-user",

                "create-permission",
                "view-permission",
                "update-permission",
                "delete-permission",
            },
        },
        {
            Title:       "User",
            Description: "Normal User",
            Permissions: []string{
                "view-role",
                "view-permission",
            },
        },
        {
            Title:       "Supporter",
            Description: "Supporter",
            Permissions: []string{},
        },
    }

    for _, roleData := range roles {
        // updateOrCreate role 
        var role models.Role
        s.DB.Where(models.Role{Title: roleData.Title}).FirstOrCreate(&role, models.Role{
            Title:       roleData.Title,
            Discription: roleData.Description,
        })

        for _, permTitle := range roleData.Permissions {
            // updateOrCreate permission
            var permission models.Permission
            s.DB.Where(models.Permission{Title: permTitle}).FirstOrCreate(&permission, models.Permission{
                Title:       permTitle,
                Discription: "",
            })

            // updateOrCreate role_permission pivot
            var rolePermission models.RolePermission
            s.DB.Where(models.RolePermission{
                RoleID:       role.ID,
                PermissionID: permission.ID,
            }).FirstOrCreate(&rolePermission)
        }

        log.Printf("✅ Seeded role: %s with %d permissions\n", roleData.Title, len(roleData.Permissions))
    }
}