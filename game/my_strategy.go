package main

import (
	. "aicup2019/model"
	мФмт "fmt"
)

//MyStrategy -- экспортируемая структура для стратегии
type MyStrategy struct{}

//NewMyStrategy -- возвращает указатель на MyStrategy
func NewMyStrategy() MyStrategy {
	return MyStrategy{}
}

//Вычисляет дистанцию до квардрата
func _ДистанцияКвадрат(a Vec2Float64, b Vec2Float64) float64 {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.X)*(a.Y-b.Y)
}

//Получает действие для стратегии
func (strategy MyStrategy) getAction(пЮнит Unit, пИгра Game, пОтлад Debug) UnitAction {
	var враг *Unit //Создание врагов
	for _, другие := range пИгра.Units {
		if другие.PlayerId != пЮнит.PlayerId {
			if враг == nil || _ДистанцияКвадрат(пЮнит.Position, другие.Position) < _ДистанцияКвадрат(пЮнит.Position, враг.Position) {
				враг = &другие
			}
		}
	}
	var оружие *LootBox
	for _, лутБокс := range пИгра.LootBoxes {
		switch лутБокс.Item.(type) {
		case *ItemWeapon:
			if оружие == nil || _ДистанцияКвадрат(пЮнит.Position, лутБокс.Position) < _ДистанцияКвадрат(пЮнит.Position, оружие.Position) {
				оружие = &лутБокс
			}
		}
	}
	позЦель := пЮнит.Position
	if пЮнит.Weapon == nil && оружие != nil {
		позЦель = оружие.Position
	} else if враг != nil {
		позЦель = враг.Position
	}
	пОтлад.Draw(CustomDataLog{
		Text: мФмт.Sprintf("позЦель: %v", позЦель),
	})
	прицел := Vec2Float64{
		X: 0,
		Y: 0,
	}
	if враг != nil {
		прицел = Vec2Float64{
			X: враг.Position.X - пЮнит.Position.X,
			Y: враг.Position.Y - пЮнит.Position.Y,
		}
	}
	прыжок := позЦель.Y > пЮнит.Position.Y
	if позЦель.X > пЮнит.Position.X && пИгра.Level.Tiles[int(пЮнит.Position.X+1)][int(пЮнит.Position.Y)] == TileWall {
		прыжок = true
	}
	if позЦель.X < пЮнит.Position.X && пИгра.Level.Tiles[int(пЮнит.Position.X-1)][int(пЮнит.Position.Y)] == TileWall {
		прыжок = true
	}
	return UnitAction{
		Velocity:   позЦель.X - пЮнит.Position.X,
		Jump:       прыжок,
		JumpDown:   !прыжок,
		Aim:        прицел,
		SwapWeapon: false,
		PlantMine:  false,
	}
}
