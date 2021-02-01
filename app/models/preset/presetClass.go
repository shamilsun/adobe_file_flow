package preset

import (
	"../database"
	"../helpers"
	"../imageFormat"
	//	"github.com/andlabs/ui"
	"log"
)

var db = database.GetDB()

func PresetExists(pl *IPresetList, v *IPreset) bool {
	if pl != nil && pl.List != nil && len(pl.List) > 0 {
		for i := 0; i < len(pl.List); i++ {
			if pl.List[i].Id == v.Id {
				log.Println("exist")
				return true
			}
		}
	}
	log.Println("not exist")
	return false
}

func (pl *IPresetList) removePreset(v *IPreset) {
	var list []*IPreset

	if len(pl.List) > 0 {
		for i := 0; i < len(pl.List); i++ {
			if pl.List[i].Id != v.Id {
				list = append(list, pl.List[i])
			}
		}
	}
	log.Println("removed")
	log.Println(list)
	pl.List = list
	//return list
}

func (pl *IPresetList) AddPreset(p *IPreset) {
	if !PresetExists(pl, p) {
		pl.List = append(pl.List, p)
	}
	log.Println("pl.list appended:")
	log.Println(pl.List)
}

func (pl *IPresetList) RemovePreset(p *IPreset) {
	if PresetExists(pl, p) {
		pl.removePreset(p)
	}
	log.Println(pl.List)
}

func (p *IPreset) AddFormat(f imageFormat.AImageFormat) {
	if !helpers.ItemExists(p.Formats, f) {
		p.Formats = append(p.Formats, f)
	}
	log.Println(p.Formats)
}

func (p *IPreset) RemoveFormat(f imageFormat.AImageFormat) {
	if helpers.ItemExists(p.Formats, f) {
		p.Formats = helpers.RemoveItem(p.Formats, f)
	}
}

func (p *IPreset) AddSize(s APresetSize) {
	if !helpers.ItemExists(p.PresetSizes, s) {
		p.PresetSizes = append(p.PresetSizes, s)
	}
}

func (p *IPreset) RemoveSize(s APresetSize) {
	if helpers.ItemExists(p.PresetSizes, s) {
		p.PresetSizes = helpers.RemoveItem(p.PresetSizes, s)
	}
}

func (p *IPreset) Save() {
	log.Println(p)
	if p.Id == 0 {
		_, err := db.Model(p).Insert()
		if err != nil {
			panic(err)
		}

	} else {
		_, err := db.Model(p).WherePK().Update() //.Where("uid = ?0", s.ID).Update()
		if err != nil {
			panic(err)
		}
	}
}

func GetProjectPresets() []*IPreset {
	var r []*IPreset
	err := db.Model(&r).Select()
	if err != nil {
		panic(err)
	}

	log.Println("project presets:")
	log.Println(r)

	return r
}
