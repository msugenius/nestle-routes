import { IMarker } from '@/interfaces/IMarker'
import { IPoint } from '@/interfaces/IPoint'
import { v4 as uuidv4 } from 'uuid'

export const PointsToMarkers = (points: IPoint[] | undefined): IMarker[] => {
  if (!points) {
    console.warn('Points undefined!')
    return []
  }

  if (points.length === 0) {
    console.warn('Points is empty!')
    return []
  }

  return points.map(point => {
    return { location: [point.latitude, point.longitude], title: `${point.name}\n${point.address}`, key: uuidv4() }
  })
}
