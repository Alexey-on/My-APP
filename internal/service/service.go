package service

import (
    "internal/model"
    "internal/repository"
)

func ProcessProducts() 
   
    laptop := &model.Laptop{
        Product: model.Product{
            ID: "",
            Name: "Ноутбук",
            Price:  
        },
    
        
    }

	headphones := &model.Headphones{
        Product: model.Product{
            ID: "",
            Name: "Беспроводные наушники",
            Price: 
        },
       
    }

	tv := &model.TV{
        Product: model.Product{
            ID: "",
            Name: "Телевизор Smart",
            Price: ,
        },
        

         repository.Store(Laptop)
         repository.Store(Headphones)
         repository.Store(TV)
}


func StartProcessing() {
 ticker := time.NewTicker(time.Second * 5) // каждые 5 секунд
 go func() {
 for range ticker.C {
 ProcessProducts()
 }
 }()
}