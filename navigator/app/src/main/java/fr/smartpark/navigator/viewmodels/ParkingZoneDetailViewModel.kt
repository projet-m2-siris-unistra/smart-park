package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import fr.smartpark.navigator.data.ParkingZone
import fr.smartpark.navigator.data.ParkingZoneRepository

class ParkingZoneDetailViewModel internal constructor(
    zoneRepository: ParkingZoneRepository,
    zoneId: String
) : ViewModel() {
    val zone: LiveData<ParkingZone> = zoneRepository.getZone(zoneId)
}