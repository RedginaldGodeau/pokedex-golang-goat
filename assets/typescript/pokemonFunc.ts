import { PokemonApi } from "@api/apis";
import { PokemonPokemonResponse } from "@api/models";
import { getColorByType, getPokemonSprite } from "./utils/pokemon";

interface PokemonSearchProps {
  pokemonList: PokemonPokemonResponse[]
  searchInput: string
  initData: () => Promise<void>
  searchSubmit: () => Promise<void>
  getPokemonType: (type: string) => string
}

export function PokemonSearch(api: PokemonApi): PokemonSearchProps {

  return {
    pokemonList: [],
    searchInput: "",
    getPokemonType(type: string) { return getColorByType(type)},
    async initData() {
      try {
        this.pokemonList = await api.apiPokemonGet()
      } catch (e) {
        console.error(e)
      }
    },
    async searchSubmit(){
      try {
        if (this.searchInput === "")
          this.pokemonList = await api.apiPokemonGet();
        else
          this.pokemonList = await api.apiPokemonSearchSearchGet({search:this.searchInput});
      } catch (e) {
        console.error(e)
      }
    }
  }
}

interface PokemonByIDProps {
  pokemon: PokemonPokemonResponse | undefined
  sprite: string
  initData: (id: number) => void
  getPokemonType: (type: string) => string
}

export function PokemonByID(api: PokemonApi): PokemonByIDProps {

  return {
    sprite: "",
    pokemon: undefined,
    getPokemonType(type: string) { return getColorByType(type)},
    async initData(id: number) {
      getPokemonSprite(id).then(img => this.sprite = img)
      api.apiPokemonPokedexPokedexGet({pokedex: id.toString()}).then( data => this.pokemon = data[0])
    },
  }
}