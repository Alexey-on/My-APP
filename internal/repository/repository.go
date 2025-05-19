package repository

import "internal/model"

var (
	laptops    []*model.Laptop
	headphones []*model.Headphones
	tvs        []*model.TV
)

switch v (type) := model. {
	
case *model.Laptop:
	model.Laptop = append(model.Laptop,v)

case *model.Headphones:
	*model.Headphones = append(*model.Headphones,v)	

case *model.TV = append(*model.TV,v)	
	
}