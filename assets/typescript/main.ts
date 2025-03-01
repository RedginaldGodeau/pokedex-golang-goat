import Alpine from "alpinejs";
import { Alpine as AlpineType } from "alpinejs";
import { PokemonByID, PokemonSearch } from "./pokemonFunc";
import { PokemonApi } from "@api/apis";
import { Configuration, DefaultConfig } from "@api/runtime";

declare global {
  const Alpine: AlpineType;
}

document.addEventListener('DOMContentLoaded', () => {
  window.Alpine = Alpine;
  Alpine.start();
})

document.addEventListener("alpine:init", () => {
  const api = new PokemonApi(new Configuration({...DefaultConfig, basePath: "http://localhost:8080".replace(/\/+$/, "")}))

  Alpine.magic("searchPokemon", () => PokemonSearch(api))
  Alpine.magic("getPokemon", () => PokemonByID(api))
})