package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import fr.smartpark.navigator.data.ParkingZone
import fr.smartpark.navigator.data.ParkingZoneRepository
import javax.inject.Inject

class ParkingZoneListViewModel @Inject constructor(zoneRepository: ParkingZoneRepository) :
    ViewModel() {
    val zones: LiveData<List<ParkingZone>> = zoneRepository.getZones()
}