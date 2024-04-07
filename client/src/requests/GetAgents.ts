import { $fetch } from '../axios'
import { IAgent } from '../interfaces/IAgent'

export const GetAgents = async (day: number, region: string): Promise<IAgent[]> => {
	const url = `agents?day=${day}&region=${region}` 
	try {
		const { data, status,  } = await $fetch.get<IAgent[]>(url)
		if (status != 200) {
			console.warn(`Agents requst status != 200`)
			return [] as IAgent[]
		}
		return data
	} catch (err) {
		console.error(err)
		return [] as IAgent[]
	}
}