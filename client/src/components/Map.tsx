import { useEffect } from 'react'
import {
  MapContainer,
  Marker,
  Popup,
  TileLayer,
  useMapEvents,
} from 'react-leaflet'
import { IMarker } from '../interfaces/IMarker'
import { useMainStore } from '../stores/main.store'

interface Props {
  markers: IMarker[]
}

const Map = (props: Props) => {
  return (
    <MapContainer
      center={[49.842957, 24.031111]}
      zoom={12}
      scrollWheelZoom={true}
      style={{ height: '92vh', width: '100vw' }}
    >
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url='https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png'
      />
      {props.markers.map(marker => (
        <Marker position={marker.location} key={marker.key}>
          <Popup>{marker.title}</Popup>
        </Marker>
      ))}
      <Centrer></Centrer>
    </MapContainer>
  )
}

function Centrer() {
  const { center } = useMainStore()
  const map = useMapEvents({})

  useEffect(() => {
    map.flyTo(center)
  }, [center])

  return null
}

export default Map
