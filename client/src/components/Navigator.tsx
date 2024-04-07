import { IPoint } from '@/interfaces/IPoint'
import { GetAgents } from '@/requests/GetAgents'
import { GetPoints } from '@/requests/GetPoints'
import { useMainStore } from '@/stores/main.store'
import { PointsToMarkers } from '@/utils/point_to_marker'
import { useQuery } from '@tanstack/react-query'
import { useEffect, useState } from 'react'
import Map from './Map'

const Navigator = () => {
  const { setCenter } = useMainStore()
  const [day, setDay] = useState(1)
  const [region, setRegion] = useState('000000021')
  const [route, setRoute] = useState(-1)
  const [points, setPoints] = useState<IPoint[]>([])

  const { data: agents } = useQuery({
    queryKey: ['agents', day, region],
    queryFn: () => GetAgents(day, region),
    staleTime: Infinity,
  })

  const { data: recievedPoints } = useQuery({
    queryKey: ['agents', route],
    queryFn: () => GetPoints(route),
    staleTime: Infinity,
    retry: false,
  })

  useEffect(() => {
    if (recievedPoints) {
      setPoints(recievedPoints)
      if (recievedPoints.length > 0) {
        setCenter([recievedPoints[0].latitude, recievedPoints[0].longitude])
      }
    }
  }, [recievedPoints])

  return (
    <>
      <div className='grid grid-cols-12 gap-2 p-3'>
        <label className='form-control w-full max-w-xs col-span-3'>
          <select
            className='select select-bordered select-sm w-full max-w-xs'
            value={day}
            onChange={e => setDay(Number(e.target.value))}
          >
            <option value={8}>Всі</option>
            <option value={1}>Понеділок</option>
            <option value={2}>Вівторок</option>
            <option value={3}>Середа</option>
            <option value={4}>Четвер</option>
            <option value={5}>П'ятниця</option>
            <option value={7}>Неділя</option>
          </select>
        </label>
        <label className='form-control w-full max-w-xs col-span-3'>
          <select
            className='select select-bordered select-sm w-full max-w-xs'
            value={region}
            onChange={e => setRegion(e.target.value)}
          >
            <option value='000000021'>Хмельницький</option>
            <option value='000000022'>Чернівці</option>
          </select>
        </label>
        <label className='form-control w-full max-w-xs col-span-3'>
          <select
            className='select select-bordered select-sm w-full max-w-xs'
            value={route}
            onChange={e => {
              setRoute(Number(e.target.value))
            }}
          >
            <option value={-1} key={-1}>
              Оберіть агента
            </option>
            <option value={0} key={0}>
              Всі
            </option>
            {agents?.map(agent => (
              <option value={agent.id} key={agent.id}>
                {agent.name}
              </option>
            ))}
          </select>
        </label>
        <button className='btn btn-error' onClick={() => setPoints([])}>
          Очистити
        </button>
      </div>
      <Map markers={PointsToMarkers(points)} />
    </>
  )
}

export default Navigator
