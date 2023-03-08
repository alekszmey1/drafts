package main

import (
	"strconv"
	"strings"
)

type basket struct {
	ti string //Название ингредиента
	pi int    //стоимость ингредиента
	ai int    //количество ингредиента
	ui string // единица измерения
}

func newBasket(l string) *basket {
	s := strings.Split(l, " ")
	b := basket{}
	b.ti = s[0]
	b.pi, _ = strconv.Atoi(s[1])
	b.ai, _ = strconv.Atoi(s[2])
	b.ui = s[3]
	return &b
}

// хранилище блюд
type basketStorage struct {
	basketName map[string]*basket
}

// создание хранилища блюд
func newBasketStorage() *basketStorage {
	return &basketStorage{basketName: make(map[string]*basket)}
}

// добавление блюда в хранилище
func (bs *basketStorage) put(b *basket) {
	bs.basketName[b.ti] = b
}
func (bs *basketStorage) getAll() []*basket {
	var baskets []*basket
	for _, v := range bs.basketName {
		baskets = append(baskets, v)
	}
	return baskets
}
