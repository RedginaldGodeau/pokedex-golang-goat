export function getColorByType(type: string) {
  switch (type) {
    case "Normal": return "bg-gray-300"
    case "Grass": return "bg-green-600"
    case "Fire": return "bg-amber-600"
    case "Poison": return "bg-purple-600"
    case "Flying": return "bg-sky-600"
    case "Water": return "bg-cyan-600"
    case "Dragon": return "bg-blue-600"
    case "Bug": return "bg-lime-600"
    case "Ice": return "bg-teal-600"
    case "Electric": return "bg-yellow-600"
    case "Ground": return "bg-stone-600"
    case "Fairy": return "bg-pink-600"
    case "Fighting": return "bg-orange-700"
    case "Psychic": return "bg-rose-300"
    case "Rock": return "bg-stone-400"
    case "Steel": return "bg-gray-600"
    case "Ghost": return "bg-indigo-300"
    default: return "bg-gray-300"
  }
}

export async function getPokemonSprite(id: number) : Promise<string> {
  try {
    const resp = await fetch(`https://pokeapi.co/api/v2/pokemon-form/${id}/`)
    if (!resp.ok) throw resp.status
    const json = await resp.json()
    return json["sprites"]["front_default"]
  } catch (err) {
    console.error(err)
  }
  return ""
}