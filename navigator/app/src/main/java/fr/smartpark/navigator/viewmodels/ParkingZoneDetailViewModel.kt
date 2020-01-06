package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.*
import fr.smartpark.navigator.data.ParkingZone
import fr.smartpark.navigator.data.ParkingZoneRepository
import kotlinx.coroutines.launch
import javax.inject.Inject

class ParkingZoneDetailViewModel @Inject constructor(
    zoneRepository: ParkingZoneRepository
) : ViewModel() {
    private val _zoneId = MutableLiveData<String>()
    val zone: LiveData<ParkingZone> = _zoneId.switchMap { id -> zoneRepository.getZone(id) }

    fun start(zoneId: String) {
        _zoneId.postValue(zoneId)
    }
}