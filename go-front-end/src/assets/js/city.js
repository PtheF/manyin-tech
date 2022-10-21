import cityList from "@/../public/cityList.json"
import { isEmpty } from "./utils"

export const getAllProvinces = function() {
    let provs = []

    for(let prov of cityList.provinceList) {
        provs.push({name: prov.name, hasNext: true})
    }

    return provs
}

export const getCitiesByProv = function(prov) {
    let cities = []

    if(isEmpty(prov)){
        return cities
    }

    for(let p of cityList.provinceList){
        if(p.name === prov){
            for(let city of p.cityList){
                cities.push({name: city.name, hasNext: true})
            }
        }
    }

    if(cities.length === 1 && cities[0].name === '市辖区'){
        return getCountyByProvCity(prov, cities[0].name)
    }

    return cities
}

export const getCountyByProvCity = function(prov, city) {
    let county = []

    if(isEmpty(prov) || isEmpty(city)){
        return county
    }

    for(let p of cityList.provinceList){
        if(p.name === prov){
            for(let c of p.cityList){
                if(c.name === city){
                    for(let ec of c.countyList){
                        county.push({name: ec.name, id: ec.id, hasNext: false})
                    }
                }
            }
        }
    }

    return county
}
