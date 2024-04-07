import { create } from 'zustand'
import { IMarker } from '../interfaces/IMarker'

interface MainState {
  center: [number, number]
  setCenter: (newCenter: [number, number]) => void
  markers: IMarker[]
}

export const useMainStore = create<MainState>()(set => ({
  center: [49.842957, 24.031111],
  setCenter: newCenter => set({ center: newCenter }),
  markers: [],
}))
