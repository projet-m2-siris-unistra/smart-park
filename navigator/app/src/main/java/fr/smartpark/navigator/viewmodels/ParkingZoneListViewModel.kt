package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import fr.smartpark.navigator.data.ParkingZone
import fr.smartpark.navigator.data.ParkingZoneRepository

class ParkingZoneListViewModel internal constructor(zoneRepository: ParkingZoneRepository): ViewModel() {
    val zones: LiveData<List<ParkingZone>> = zoneRepository.getZones()
}