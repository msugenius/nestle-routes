import axios from 'axios'

export const $fetch = axios.create({
	baseURL: "http://10.0.15.134:8000/api/v1"
})
