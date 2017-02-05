package beater

type ZkillPackage struct {
	Payload ZkillKill `json:"package"`
}

type Killmail struct {
	SolarSystem   SolarSystem `json:"solarSystem"`
	KillTime      string     `json:"killTime"`
	AttackerCount int        `json:"attackerCount"`
	Attackers     []Attacker `json:"attackers"`
	Victim        Victim     `json:"victim"`
	War           EveAPIItem `json:"war"`
}

type SolarSystem struct {
	EveAPIItem
	Region         string     `json:"region"`
	Constellation  string     `json:"constellation"`
	SecurityStatus float64    `json:"SecurityStatus"`
}

type Attacker struct {
	EvePlayer
	FinalBlow  bool       `json:"finalBlow"`
	DamageDone int        `json:"damageDone"`
	Faction    EveAPIItem `json:"faction"`
}

type Victim struct {
	EvePlayer
	DamageTake int      `json:"damageTaken"`
	Items      []Item   `json:"items"`
	Position   Position `json:"position"`
}

type Item struct {
	Type              EveAPIItem `json:"itemType"`
	Flag              int        `json:"flag"`
	QuantityDestroyed int        `json:"quantityDestroyed"`
	QuantityDropped   int        `json:"quantityDropped"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type EvePlayer struct {
	Alliance       EveAPIItem `json:"alliance"`
	Corporation    EveAPIItem `json:"corporation"`
	Character      EveAPIItem `json:"character"`
	Ship           EveAPIItem `json:"shipType"`
	WeaponType     EveAPIItem `json:"weaponType"`
	SecurityStatus float64    `json:"securityStatus"`
}

type ZkillMetadata struct {
	LocationID int     `json:"locationID"`
	Hash       string  `json:"hash"`
	TotalValue float64 `json:"totalValue"`
	Points     int     `json:"points"`
	Source     string  `json:"href"`
}

type EveAPIItem struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Source string `json:"href"`
}

type ZkillKill struct {
	KillID   int           `json:"killID"`
	Kill     Killmail      `json:"killmail"`
	Metadata ZkillMetadata `json:"zkb"`
}
