package models

import (
    "errors"
    "gorm.io/gorm"
)

type Todo struct {
    gorm.Model
    Content string
    Category string
}

const (
    MaxMustCount = 5
    MaxWantCount = 50
)

func CreateTodo(content string, category string) (err error) {
    var count int64
    if category == "must" {
        Db.Model(&Todo{}).Where("category = ?", "must").Count(&count)
        if count >= MaxMustCount {
            return errors.New("Over limit")
        }
    } else if category == "want" {
        Db.Model(&Todo{}).Where("category = ?", "want").Count(&count)
        if count >= MaxWantCount {
            return errors.New("Over limit")
        }
    } else {
        return errors.New("Only use must or want")
    } 
            
    todo := Todo{
        Content: content,
        Category: category,
    }
    return Db.Create(&todo).Error
}

func DeleteTodo(id int) error {
    err := Db.Delete(&Todo{}, id).Error
    return err
}

func GetTodo(id int) (Todo, error) {
    var todo Todo
    err := Db.First(&todo, id).Error

    return todo, err
}

func UpdateTodo(t Todo) (err error) {
    var oldTodo Todo
    if err := Db.First(&oldTodo, t.ID).Error; err != nil {
        return err
    }

    if oldTodo.Category != t.Category {
        var count int64
        if t.Category == "must" {
            Db.Model(&Todo{}).Where("category = ?", "must").Count(&count)
            if count >= MaxMustCount {
                return errors.New("Over limit")
            }
        } else if t.Category == "want" {
            Db.Model(&Todo{}).Where("category = ?", "want").Count(&count)
            if count >= MaxWantCount {
                return errors.New("Over limit")
            }
        } else {
            return errors.New("Only use must or want")
        }
    }

    return Db.Save(&t).Error
}
