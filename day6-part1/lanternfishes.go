package main

type LanternFish struct {
	SpawnTimer int
}

func NewLanternFish(spawnTimer int) *LanternFish {
	return &LanternFish{SpawnTimer: spawnTimer}
}

func (lf *LanternFish) Cycle() (spawnNewLanternFish bool) {
	if lf.SpawnTimer == 0 {
		spawnNewLanternFish = true
		lf.SpawnTimer = 6
		return
	}

	lf.SpawnTimer--
	return
}

type Swarm []*LanternFish

func (s *Swarm) Cycle() {
	var newLanternFishes []*LanternFish

	for _, lanternFish := range *s {
		if spawnNewLanternFish := lanternFish.Cycle(); spawnNewLanternFish {
			newLanternFishes = append(newLanternFishes, NewLanternFish(8))
		}
	}

	*s = append(*s, newLanternFishes...)
}
