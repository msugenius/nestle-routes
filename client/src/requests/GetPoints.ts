import { $fetch } from '../axios'
import { IPoint } from '../interfaces/IPoint'

export const GetPoints = async (route: number): Promise<IPoint[]> => {
	const url = `points?route=${route}` 
	try {
		const { data, status,  } = await $fetch.get<IPoint[]>(url)
		if (status != 200) {
			console.warn(`Points requst status != 200`)
			return [] as IPoint[]
		}
		return data
	} catch (err) {
		console.error(err)
		return [] as IPoint[]
	}
}