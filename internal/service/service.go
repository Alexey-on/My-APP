package service

import (
    "internal/model"
    "internal/repository"
)

// Функция бизнес-логики
func ProcessProducts() {
    // Создаем разные типы товаров
    laptop := &model.Laptop{
        Product: model.Product{
            ID: "",
            Name: "Ноутбук",
            Price:  
        },
        CPU: ",",
        RAM: ',',
    }

	headphones := &model.Headphones{
        Product: model.Product{
            ID: "",
            Name: "Беспроводные наушники",
            Price: 
        },
        Type: "беспроводные",
        Connection: "Bluetooth",
    }

	tv := &model.TV{
        Product: model.Product{
            ID: "",
            Name: "Телевизор Smart",
            Price: ,
        },
        ScreenSize: ,
        Resolution: "",
    }
}
